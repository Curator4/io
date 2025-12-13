import { Client, GatewayIntentBits, Events, Partials } from 'discord.js';
import { GrpcClient } from './grpc/client.js';
import { handleMessage } from './handlers/message.js';

// grpc client, see ./grpc/client
const grpcHost = process.env.GRPC_HOST || 'localhost';
const grpcPort = parseInt(process.env.GRPC_PORT || '50051');
const grpcClient = new GrpcClient(grpcHost, grpcPort);

// discord client
const token = process.env.DISCORD_TOKEN;
const discordClient = new Client({
  intents: [
    GatewayIntentBits.Guilds,
    GatewayIntentBits.GuildMessages,
    GatewayIntentBits.MessageContent,
    GatewayIntentBits.DirectMessages,
  ],
  partials: [Partials.Channel, Partials.Message],
});

// rdy message on client init
discordClient.once(Events.ClientReady, (readyClient) => {
  console.log(`ready, logged in as ${readyClient.user.tag}`);
});

// message events
discordClient.on(Events.MessageCreate, async (message) => {
  await handleMessage(message, grpcClient);
});

// login with token
discordClient.login(token);
