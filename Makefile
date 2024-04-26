include .env
export $(shell sed 's/=.*//' .env)

migration_create:
	migrate create -ext sql -dir .dev/db-migrations/ $(name)

migration_run:
	migrate -path .dev/db-migrations/ -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" up

migration_drop:
	migrate -path .dev/db-migrations/ -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" drop -f

api_build:
	go build -o bin/api build/api/main.go

api_run:
	make build_api && ./bin/api

api_dev:
	air

api_build_image:
	docker build -t ${DOCKER_API_IMAGE_NAME} -f build/api/Dockerfile .

api_run_container:
	docker run --env-file=.env -p ${DOCKER_API_HOST_PORT}:${APP_PORT} ${DOCKER_API_IMAGE_NAME}

compose_run:
	docker compose --env-file=.env -f .dev/docker-compose.yml up

cli_build:
	go build -o bin/cli build/cli/main.go

cli_run:
	make build_cli && ./bin/cli

test:
	go test -v $(shell go list ./... | grep -v /vendor/) -cover -coverprofile=coverage.out

scan_docker:
	docker run --rm -e SONAR_HOST_URL=${SONAR_HOST_URL} -e SONAR_LOGIN=${SONAR_LOGIN} -e SONAR_PASSWORD=${SONAR_PASSWORD} -e SONAR_PROJECT_KEY=${SONAR_PROJECT_KEY} -e SONAR_SCANNER_OPTS="-Dsonar.projectKey=${SONAR_PROJECT_KEY} -Dsonar.sources=. -Dsonar.exclusions='**/*_test.go,bin/*,tmp/*' -Dsonar.go.coverage.reportPaths=coverage.out" -v $(shell pwd):/usr/src --net=host sonarsource/sonar-scanner-cli

scan:
	sonar-scanner -Dsonar.projectKey=${SONAR_PROJECT_KEY} -Dsonar.sources=. -Dsonar.host.url=${SONAR_HOST_URL} -Dsonar.login=${SONAR_LOGIN} -Dsonar.password=${SONAR_PASSWORD} -Dsonar.go.coverage.reportPaths=coverage.out

test_and_scan:
	make test && make docker_sonar_scan