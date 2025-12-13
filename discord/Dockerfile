# build
FROM node:20 AS builder
WORKDIR /app
RUN apt-get update && \
  apt-get install -y --no-install-recommends protobuf-compiler && \
  rm -rf /var/lib/apt/lists/*
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run proto:gen && \
  npm run build

# runtime
FROM node:20-slim
WORKDIR /app
COPY package*.json ./
RUN npm ci --omit=dev
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/src/grpc/generated ./src/grpcgenerated
CMD ["node", "dist/index.js"]
