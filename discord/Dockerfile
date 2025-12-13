# build
FROM node:20 AS builder
WORKDIR /app
RUN apt-get update && \
  apt-get install -y --no-install-recommends \
    protobuf-compiler \
    libprotobuf-dev && \
  rm -rf /var/lib/apt/lists/*

# Copy package files from discord directory
COPY discord/package*.json ./
RUN npm ci

# Copy discord source and proto files
COPY discord/ .
COPY proto/ ../proto/

RUN npm run proto:gen && \
  npm run build

# runtime
FROM node:20-slim
WORKDIR /app
COPY discord/package*.json ./
RUN npm ci --omit=dev
COPY --from=builder /app/dist ./dist
CMD ["node", "dist/index.js"]
