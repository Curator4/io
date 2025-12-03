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

# docker
db-up: 
	docker compose up -d db
db-down:
	docker compose down
db-logs:
	docker compose logs -f db
