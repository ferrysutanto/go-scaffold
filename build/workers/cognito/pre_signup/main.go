package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// payload is the event payload for the pre-signup trigger
type payload struct {
	Version       string `json:"version"`
	Region        string `json:"region"`
	UserPoolID    string `json:"userPoolId"`
	UserName      string `json:"userName"`
	CallerContext struct {
		AwsSdkVersion string `json:"awsSdkVersion"`
		ClientID      string `json:"clientId"`
	} `json:"callerContext"`
	TriggerSource string `json:"triggerSource"`
	Request       struct {
		UserAttributes struct {
			Email string `json:"email"`
		} `json:"userAttributes"`
		ValidationData map[string]interface{} `json:"validationData"`
	} `json:"request"`
	Response struct {
		AutoConfirmUser bool `json:"autoConfirmUser"`
		AutoVerifyEmail bool `json:"autoVerifyEmail"`
		AutoVerifyPhone bool `json:"autoVerifyPhone"`
	}
}

func handleRequest(ctx context.Context, event json.RawMessage) (json.RawMessage, error) {
	log.Println("Processing pre-signup event")

	log.Println("Event: ", string(event))

	// Unmarshal the event
	var payload payload
	if err := json.Unmarshal(event, &payload); err != nil {
		return nil, err
	}

	// Set the response

	// Return a valid JSON response
	return event, nil
}

func main() {
	lambda.Start(handleRequest)
}
