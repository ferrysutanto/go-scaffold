package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, event json.RawMessage) error {
	log.Println("Processing event")

	log.Println("Event: ", string(event))

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
