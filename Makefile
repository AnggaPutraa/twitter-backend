include .env

db-up:
	docker-compose up -d

db-down:
	docker-compose down

migrate-up:
	migrate -path db/migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down

.PHONY: db-up db-down migrate-up migrate-down