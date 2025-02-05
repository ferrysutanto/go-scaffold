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

# ECR section
login-ecr:
	aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com

# Lambda Cognito Section
pre-signup/build:
	go build -o bin/pre-signup build/workers/cognito/pre_signup/main.go

pre-signup/build-image:
	docker build -t ${DOCKER_COGNITO_PRE_SIGNUP_IMAGE_NAME} -f build/workers/cognito/pre_signup/Dockerfile .

pre-signup/tag:
	docker tag ${DOCKER_COGNITO_PRE_SIGNUP_IMAGE_NAME} ${DOCKER_COGNITO_PRE_SIGNUP_IMAGE_NAME}:latest

pre-signup/tag-ecr:
	docker tag ${DOCKER_COGNITO_PRE_SIGNUP_IMAGE_NAME} ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${DOCKER_COGNITO_PRE_SIGNUP_IMAGE_NAME}:latest

pre-signup/push-image:
	docker push ${DOCKER_COGNITO_PRE_SIGNUP_IMAGE_NAME}:latest

pre-signup/push-image-ecr:
	docker push ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${DOCKER_COGNITO_PRE_SIGNUP_IMAGE_NAME}:latest

pre-signup/build-push:
	make pre-signup/build-image && make pre-signup/tag && make pre-signup/push-image

pre-signup/build-push-ecr:
	make pre-signup/build-image && make pre-signup/tag-ecr && make pre-signup/push-image-ecr

pre-signup/update-lambda:
	@echo "Fetching latest image digest..."
	@echo "Updating Lambda function $(LAMBDA_COGNITO_PRE_SIGNUP_FUNCTION_NAME)..."
	aws lambda update-function-code \
	    --function-name $(LAMBDA_COGNITO_PRE_SIGNUP_FUNCTION_NAME) \
	    --image-uri $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/$(DOCKER_COGNITO_PRE_SIGNUP_IMAGE_NAME):latest
	@echo "Lambda function updated successfully!"

pre-signup/redeploy:
	@echo "Redeploying Lambda function..."
	aws lambda wait function-updated --function-name $(LAMBDA_COGNITO_PRE_SIGNUP_FUNCTION_NAME)
	@echo "Lambda function redeployed successfully!"

post-confirmation/build:
	go build -o bin/post-confirmation build/workers/cognito/post_confirmation/main.go

post-confirmation/build-image:
	docker build -t ${DOCKER_COGNITO_POST_CONFIRMATION_IMAGE_NAME} -f build/workers/cognito/post_confirmation/Dockerfile .

post-confirmation/tag:
	docker tag ${DOCKER_COGNITO_POST_CONFIRMATION_IMAGE_NAME} ${DOCKER_COGNITO_POST_CONFIRMATION_IMAGE_NAME}:latest

post-confirmation/tag-ecr:
	docker tag ${DOCKER_COGNITO_POST_CONFIRMATION_IMAGE_NAME} ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${DOCKER_COGNITO_POST_CONFIRMATION_IMAGE_NAME}:latest

post-confirmation/push-image:
	docker push ${DOCKER_COGNITO_POST_CONFIRMATION_IMAGE_NAME}:latest

post-confirmation/push-image-ecr:
	docker push ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${DOCKER_COGNITO_POST_CONFIRMATION_IMAGE_NAME}:latest

post-confirmation/build-push:
	make post-confirmation/build-image && make post-confirmation/tag && make post-confirmation/push-image

post-confirmation/build-push-ecr:
	make post-confirmation/build-image && make post-confirmation/tag-ecr && make post-confirmation/push-image-ecr

post-confirmation/update-lambda:
	@echo "Fetching latest image digest..."
	@echo "Updating Lambda function $(LAMBDA_COGNITO_POST_CONFIRMATION_FUNCTION_NAME)..."
	aws lambda update-function-code \
	    --function-name $(LAMBDA_COGNITO_POST_CONFIRMATION_FUNCTION_NAME) \
	    --image-uri $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/$(DOCKER_COGNITO_POST_CONFIRMATION_IMAGE_NAME):latest
	@echo "Lambda function updated successfully!"

post-confirmation/redeploy:
	@echo "Redeploying Lambda function..."
	aws lambda wait function-updated --function-name $(LAMBDA_COGNITO_POST_CONFIRMATION_FUNCTION_NAME)
	@echo "Lambda function redeployed successfully!"

post-authentication/build:
	go build -o bin/post-authentication build/workers/cognito/post_authentication/main.go

post-authentication/build-image:
	docker build -t ${DOCKER_COGNITO_POST_AUTHENTICATION_IMAGE_NAME} -f build/workers/cognito/post_authentication/Dockerfile .

post-authentication/tag:
	docker tag ${DOCKER_COGNITO_POST_AUTHENTICATION_IMAGE_NAME} ${DOCKER_COGNITO_POST_AUTHENTICATION_IMAGE_NAME}:latest

post-authentication/tag-ecr:
	docker tag ${DOCKER_COGNITO_POST_AUTHENTICATION_IMAGE_NAME} ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${DOCKER_COGNITO_POST_AUTHENTICATION_IMAGE_NAME}:latest

post-authentication/push-image:
	docker push ${DOCKER_COGNITO_POST_AUTHENTICATION_IMAGE_NAME}:latest

post-authentication/push-image-ecr:
	docker push ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${DOCKER_COGNITO_POST_AUTHENTICATION_IMAGE_NAME}:latest

post-authentication/build-push:
	make post-authentication/build-image && make post-authentication/tag && make post-authentication/push-image

post-authentication/build-push-ecr:
	make post-authentication/build-image && make post-authentication/tag-ecr && make post-authentication/push-image-ecr

post-authentication/update-lambda:
	@echo "Fetching latest image digest..."
	@echo "Updating Lambda function $(LAMBDA_COGNITO_POST_AUTHENTICATION_FUNCTION_NAME)..."
	aws lambda update-function-code \
	    --function-name $(LAMBDA_COGNITO_POST_AUTHENTICATION_FUNCTION_NAME) \
	    --image-uri $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/$(DOCKER_COGNITO_POST_AUTHENTICATION_IMAGE_NAME):latest
	@echo "Lambda function updated successfully!"

post-authentication/redeploy:
	@echo "Redeploying Lambda function..."
	aws lambda wait function-updated --function-name $(LAMBDA_COGNITO_POST_AUTHENTICATION_FUNCTION_NAME)
	@echo "Lambda function redeployed successfully!"

cognito-triggers/build-push:
	make pre-signup/build-push && make post-confirmation/build-push && make post-authentication/build-push

cognito-triggers/build-push-ecr:
	make pre-signup/build-push-ecr && make post-confirmation/build-push-ecr && make post-authentication/build-push-ecr

cognito-triggers/update-lambda:
	make pre-signup/update-lambda && make post-confirmation/update-lambda && make post-authentication/update-lambda

cognito-triggers/redeploy:
	make pre-signup/redeploy && make post-confirmation/redeploy && make post-authentication/redeploy

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