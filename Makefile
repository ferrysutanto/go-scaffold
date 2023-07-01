include .env
export $(shell sed 's/=.*//' .env)

create_migration:
	migrate create -ext sql -dir .dev/db-migrations/ $(migration_name)

run_migration:
	migrate -path .dev/db-migrations/ -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" up

drop_migration:
	migrate -path .dev/db-migrations/ -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" drop -f

build_api:
	go build -o bin/api build/api/main.go

run_api:
	build_api && go run bin/api

dev_api:
	air