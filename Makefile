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

# protobuf
proto:
	mkdir -p backend/internal/proto
	protoc -I=proto \
	       --go_out=backend/internal/proto --go_opt=module=github.com/curator4/io/backend/internal/proto \
	       --go-grpc_out=backend/internal/proto --go-grpc_opt=module=github.com/curator4/io/backend/internal/proto \
	       proto/io.proto

# docker
db-up:
	docker compose up -d db
db-down:
	docker compose down
db-logs:
	docker compose logs -f db
db-shell:
	psql "$(DATABASE_URL)"
