# build
FROM golang:1.25.5 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY . .
RUN go build -o server ./cmd/server

# runtime
FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/server .
COPY --from=builder /app/sql/schema ./sql/schema
COPY --from=builder /app/entrypoint.sh .
RUN chmod +x entrypoint.sh
EXPOSE 50051
CMD ["./entrypoint.sh"]
