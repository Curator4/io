import { Message } from 'discord.js';
import { GrpcClient } from '../grpc/client';
import { SendMessageRequest } from '../grpc/generated/io';

// helpers

// warrantsResponse is a helper that determines whether a message event warrants a response or not
// TODO: needs more features etc later on, local inferral etc, extra triggers
const warrantsResponse = (message: Message): boolean => {
  if (message.author.bot) return false;
  const isDM = message.channel.isDMBased();
  const isMentioned = message.mentions.has(message.client.user!);
  const startsWithPrefix = message.content.toLowerCase().startsWith('io');

  return isDM || isMentioned || startsWithPrefix;
};

// processMessage handles the remote procedure call to the backend grpc server, and replies to discord
const processMessage = async (
  message: Message,
  grpcClient: GrpcClient,
): Promise<void> => {
  // typing
  if ('sendTyping' in message.channel) {
    await message.channel.sendTyping();
  }

  // grpc request object
  const request: SendMessageRequest = {
    content: {
      text: message.content,
      media: [],
    },
    userId: message.author.id,
    role: 'user',
    conversationId: '',
  };

  // grpc call
  const response = await grpcClient.sendMessage(request);

  // discord reply
  await message.reply(
    response.assistantMessage?.content?.text || 'No response',
  );
};

// exports

// handleMessage is the event handler for all message events
// TODO: messages should maybe have differenlt handlers depending on context,
// like DMS should have a prompt that says it a dm not a group chat
export const handleMessage = async (
  message: Message,
  grpcClient: GrpcClient,
): Promise<void> => {
  // event handling
  try {
    // check if the message warrants a response
    if (!warrantsResponse(message)) return;
    // process the remote procedure call
    await processMessage(message, grpcClient);
  } catch (error) {
    console.error('error handling message: ', error);

    // respond to discord with error
    const errorMessage =
      error instanceof Error
        ? `Error: ${error.message}`
        : 'Unknown error occurred';
    try {
      await message.reply(`failed ❌, ${errorMessage}`);
    } catch (replyError) {
      console.error('failed to even send error msg to discord, F:', replyError);
    }
  }
};
