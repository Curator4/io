include .env
export

# migrations
migrate-up:
	goose -dir ./backend/sql/schema postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir ./backend/sql/schema postgres "$(DATABASE_URL)" down

migrate-reset:
	goose -dir ./backend/sql/schema postgres "$(DATABASE_URL)" down-to 0

migrate-status:
	goose -dir ./backend/sql/schema postgres "$(DATABASE_URL)" status

# sqlc
.PHONY: sqlc
sqlc:
	sqlc generate

# protobuf
.PHONY: proto
proto:
	# Create output directory for Go generated code
	mkdir -p backend/internal/proto
	# Generate Go protobuf and gRPC code
	# -I: proto source directory
	# --go_out: protobuf message code output
	# --go-grpc_out: gRPC service code output
	protoc -I=proto \
	       --go_out=backend/internal/proto --go_opt=module=github.com/curator4/io/backend/internal/proto \
	       --go-grpc_out=backend/internal/proto --go-grpc_opt=module=github.com/curator4/io/backend/internal/proto \
	       proto/io.proto
	# Generate TypeScript client code for Discord bot
	cd discord && npm run proto:gen

# docker
db-up:
	docker compose up -d db
db-down:
	docker compose down
db-logs:
	docker compose logs -f db
db-shell:
	psql "$(DATABASE_URL)"
