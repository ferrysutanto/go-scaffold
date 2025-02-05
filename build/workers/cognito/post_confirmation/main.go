package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

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
			Sub               string `json:"sub"`
			EmailVerified     string `json:"email_verified"`
			CognitoUserStatus string `json:"cognito:user_status"`
			GivenName         string `json:"given_name"`
			FamilyName        string `json:"family_name"`
			Email             string `json:"email"`
		} `json:"userAttributes"`
	} `json:"request"`
	Response map[string]interface{} `json:"response"`
}

func handleRequest(ctx context.Context, event json.RawMessage) (json.RawMessage, error) {
	log.Println("Processing post-confirmation event")

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
