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

run_compose:
	docker compose --env-file=.env -f .dev/docker-compose.yml up

build_cli:
	go build -o bin/cli build/cli/main.go

run_cli:
	make build_cli && ./bin/cli

test:
	go test -v $(shell go list ./... | grep -v /vendor/) -cover -coverprofile=coverage.out

docker_sonar_scan:
	docker run --rm -e SONAR_HOST_URL=${SONAR_HOST_URL} -e SONAR_LOGIN=${SONAR_LOGIN} -e SONAR_PASSWORD=${SONAR_PASSWORD} -e SONAR_PROJECT_KEY=${SONAR_PROJECT_KEY} -v $(shell pwd):/usr/src sonarsource/sonar-scanner-cli

sonar_scan:
	sonar-scanner -Dsonar.projectKey=${SONAR_PROJECT_KEY} -Dsonar.sources=. -Dsonar.host.url=${SONAR_HOST_URL} -Dsonar.login=${SONAR_LOGIN} -Dsonar.password=${SONAR_PASSWORD} -Dsonar.go.coverage.reportPaths=coverage.out

test_and_scan:
	make test && make docker_sonar_scan