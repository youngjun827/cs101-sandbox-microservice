package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    fmt.Println("event", event)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		IsBase64Encoded: false,
		MultiValueHeaders: map[string][]string{
			"X-Custom-Header": {"My value", "My other value"},
		},
		Body: "PLEASE WORK",
	}

    return response, nil
}

func main() {
    lambda.Start(HandleRequest)
}