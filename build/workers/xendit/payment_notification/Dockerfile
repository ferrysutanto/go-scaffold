# Stage 1: Build the Go binary
FROM golang:alpine AS builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY . .
RUN go mod download

# Build the Go binary for arm64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./xendit-payment-notification ./build/workers/xendit/payment_notification/main.go

# Stage 2: Create the final Docker image
FROM alpine:latest

WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/xendit-payment-notification /app/

# Set any other necessary configurations or dependencies

# Set the binary as the entrypoint
ENTRYPOINT [ "./xendit-payment-notification" ]
