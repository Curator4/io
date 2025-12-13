import { Message } from 'discord.js';
import { GrpcClient } from '../grpc/client.js';
import { SendMessageRequest, StoreMessageRequest } from '../grpc/generated/io.js';

// helpers

// warrantsResponse is a helper that determines whether a message event warrants a response or not
// TODO: needs more features etc later on, local inferral etc, extra triggers
const warrantsResponse = (message: Message): boolean => {
  const isDM = message.channel.isDMBased();
  const isMentioned = message.mentions.has(message.client.user!);
  const startsWithPrefix = message.content.toLowerCase().startsWith('io');

  return isDM || isMentioned || startsWithPrefix;
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
  let text = response.content?.text || 'No response';

  // Discord has a 2000 character limit
  if (text.length > 2000) {
    text = text.substring(0, 1997) + '...';
  }

  await message.reply(text);
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
