ENV_FILE ?= .env

# If user supplies -env=path/to/envfile, it will override the default
ifeq (,$(wildcard $(ENV_FILE)))
$(error Environment file '$(ENV_FILE)' not found)
endif

include $(ENV_FILE)
export

# DB Migration Section
migration/create:
	migrate create -ext sql -dir .dev/db-migrations/ $(name)

migration/run:
	migrate -path .dev/db-migrations/ -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" up

migration/drop:
	migrate -path .dev/db-migrations/ -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" drop -f

# API Section
api/build:
	go build -o bin/api build/api/main.go

api/run:
	make build/api && ./bin/api

api/dev:
	air

api/build_image:
	docker build -t ${DOCKER_API_IMAGE_NAME} -f build/api/Dockerfile .

api/run_container:
	docker run --env-file=.env -p ${DOCKER_API_HOST_PORT}:${APP_PORT} ${DOCKER_API_IMAGE_NAME}

# Lambda Section
pre-signup/build:
	go build -o bin/pre-signup build/workers/cognito/pre_signup/main.go

pre-signup/build-image:
	docker build -t ${DOCKER_COGNITO_PRE_SIGNUP_IMAGE_NAME} -f build/workers/cognito/pre_signup/Dockerfile .

# CLI Section
cli/build:
	go build -o bin/cli build/cli/main.go

cli/run:
	make build_cli && ./bin/cli

# Local Env Support Section
compose/up:
	docker compose --env-file=.dev/.env -f .dev/docker-compose.yml up -d

compose/down:
	docker compose --env-file=.dev/.env -f .dev/docker-compose.yml down

compose/swagger/run:
	docker compose --env-file=.dev/.env -f .dev/docker-compose.yml up --build api_swagger -d

compose/swagger/stop:
	docker compose --env-file=.env -f .dev/docker-compose.yml down api_swagger

# Test Section
test:
	go test -v $(shell go list ./... | grep -v /vendor/) -cover -coverprofile=coverage.out

scan/docker:
	docker run --rm -e SONAR_HOST_URL=${SONAR_HOST_URL} -e SONAR_LOGIN=${SONAR_LOGIN} -e SONAR_PASSWORD=${SONAR_PASSWORD} -e SONAR_PROJECT_KEY=${SONAR_PROJECT_KEY} -e SONAR_SCANNER_OPTS="-Dsonar.projectKey=${SONAR_PROJECT_KEY} -Dsonar.sources=. -Dsonar.exclusions='**/*_test.go,bin/*,tmp/*' -Dsonar.go.coverage.reportPaths=coverage.out" -v $(shell pwd):/usr/src --net=host sonarsource/sonar-scanner-cli

scan:
	sonar-scanner -Dsonar.projectKey=${SONAR_PROJECT_KEY} -Dsonar.sources=. -Dsonar.host.url=${SONAR_HOST_URL} -Dsonar.login=${SONAR_LOGIN} -Dsonar.password=${SONAR_PASSWORD} -Dsonar.go.coverage.reportPaths=coverage.out

test/scan:
	make test && make scan/docker