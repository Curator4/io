import { Message } from 'discord.js';
import { GrpcClient } from '../grpc/client.js';
import {
  SendMessageRequest,
  StoreMessageRequest,
  Action,
  ConversationLifecycle
} from '../grpc/generated/io.js';

// helpers

// warrantsResponse is a helper that determines whether a message event warrants a response or not
// TODO: needs more features etc later on, local inferral etc, extra triggers
const warrantsResponse = (message: Message): boolean => {
  const isDM = message.channel.isDMBased();
  const isMentioned = message.mentions.has(message.client.user!);
  const startsWithPrefix = message.content.toLowerCase().startsWith('io');

  return isDM || isMentioned || startsWithPrefix;
};

// executeActions processes AI-requested actions (reactions, etc.)
const executeActions = async (
  message: Message,
  actions: Action[] | undefined,
): Promise<void> => {
  if (!actions || actions.length === 0) return;

  for (const action of actions) {
    try {
      // Handle reaction action
      if (action.reaction) {
        await message.react(action.reaction.emoji);
      }
      // Future action types can be handled here
    } catch (error) {
      console.error('Failed to execute action:', action, error);
      // Continue processing other actions even if one fails
    }
  }
};

// formatToolIndicators creates indicators for which tools were used
const formatToolIndicators = (actions: Action[] | undefined): string => {
  if (!actions || actions.length === 0) return '';

  const indicators: string[] = [];

  for (const action of actions) {
    if (action.webSearch) {
      if (action.webSearch.queries.length > 0) {
        indicators.push(`🔍 Web Search: "${action.webSearch.queries[0]}"`);
      }
    } else if (action.imageGeneration) {
      indicators.push('🎨 Generated Image');
    } else if (action.codeInterpreter) {
      indicators.push('💻 Ran Code');
    }
  }

  return indicators.length > 0 ? `\n\n*Tools used:*\n${indicators.join('\n')}` : '';
};

// formatLifecycleMessage creates a user-friendly message about conversation lifecycle
const formatLifecycleMessage = (lifecycle: ConversationLifecycle | undefined): string | null => {
  if (!lifecycle || !lifecycle.isNewConversation) {
    return null;
  }

  return `_✨ Started new conversation: ${lifecycle.conversationName}_`;
};

// sendMessage calls the sendMessage remote procedure, replies to the message with the response
const sendMessage = async (
  message: Message,
  grpcClient: GrpcClient,
): Promise<void> => {
  // typing
  if ('sendTyping' in message.channel) {
    await message.channel.sendTyping();
  }

  const request: SendMessageRequest = {
    content: { text: message.content, media: [] },
    username: message.author.username,
  };

  const response = await grpcClient.sendMessage(request);

  // Build response text
  let text = response.content?.text || 'No response';

  // Prepend lifecycle message if new conversation
  const lifecycleMsg = formatLifecycleMessage(response.lifecycle);
  if (lifecycleMsg) {
    text = lifecycleMsg + '\n\n' + text;
  }

  // Append tool indicators
  const toolIndicators = formatToolIndicators(response.actions);
  text += toolIndicators;

  // Discord has a 2000 character limit
  if (text.length > 2000) {
    text = text.substring(0, 1997) + '...';
  }

  await message.reply(text);

  // Execute AI-requested actions (like reactions)
  await executeActions(message, response.actions);
};

// storeMessage calls the storeMessage remote procedure, simply storing the message in database
const storeMessage = async (
  message: Message,
  grpcClient: GrpcClient,
): Promise<void> => {
  const request: StoreMessageRequest = {
    content: { text: message.content, media: [] },
    username: message.author.username,
  };

  await grpcClient.storeMessage(request);
};


// exports

// handleMessage is the main event handler, handles all message events, ie discord messages to server
// delegates to either sendMessage or storeMessage depending on results of warrantsResponse
export const handleMessage = async (
  message: Message,
  grpcClient: GrpcClient,
): Promise<void> => {
  // Ignore all bot messages (including our own)
  if (message.author.bot) return;

  try {
    if (warrantsResponse(message)) {
      await sendMessage(message, grpcClient);
    } else {
      await storeMessage(message, grpcClient);
    }
  } catch (error) {
    console.error('error handling message:', error);

    const errorMessage = error instanceof Error ? `Error: ${error.message}` : `Unknown error occurred`;
    try {
      await message.reply(`failed ❌, ${errorMessage}`);
    } catch (replyError) {
      console.error('failed to send error msg to discord:', replyError);
    }
  }
};
