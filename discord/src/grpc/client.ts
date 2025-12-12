import { credentials } from '@grpc/grpc-js';
import {
  IOServiceClient,
  SendMessageRequest,
  SendMessageResponse,
} from './generated/io';

export class GrpcClient {
  private client: IOServiceClient;

  constructor(host: string, port: number) {
    const address = `${host}:${port}`;
    this.client = new IOServiceClient(address, credentials.createInsecure());
  }

  async sendMessage(request: SendMessageRequest): Promise<SendMessageResponse> {
    return new Promise((resolve, reject) => {
      this.client.sendMessage(request, (error, response) => {
        if (error) reject(error);
        else resolve(response);
      });
    });
  }
}
