# Stage 1: Build the Go binary
FROM golang:alpine AS builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY . .
RUN go mod download

# Build the Go binary for arm64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./cognito-post-confirmation ./build/workers/cognito/post_authentication/main.go

# Stage 2: Create the final Docker image
FROM alpine:latest

WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/cognito-post-confirmation /app/

# Set any other necessary configurations or dependencies

# Set the binary as the entrypoint
ENTRYPOINT [ "./cognito-post-confirmation" ]
