include .env
export $(shell sed 's/=.*//' .env)

create_migration:
	migrate create -ext sql -dir .dev/db-migrations/ $(name)

run_migration:
	migrate -path .dev/db-migrations/ -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" up

drop_migration:
	migrate -path .dev/db-migrations/ -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" drop -f

build_api:
	go build -o bin/api build/api/main.go

run_api:
	make build_api && ./bin/api

dev_api:
	air

build_api_image:
	docker build -t ${DOCKER_API_IMAGE_NAME} -f build/api/Dockerfile .

run_api_container:
	docker run --env-file=.env -p ${DOCKER_API_HOST_PORT}:${APP_PORT} ${DOCKER_API_IMAGE_NAME}

build_cli:
	go build -o bin/cli build/cli/main.go

run_cli:
	make build_cli && ./bin/cli