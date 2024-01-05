package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func route1Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return events.APIGatewayProxyResponse{
        Body:       "Hello from /route1 (route1Handler)",
        StatusCode: 200,
    }, nil
}

func route2Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return events.APIGatewayProxyResponse{
        Body:       "Hello from /route2 (route2Handler)",
        StatusCode: 200,
    }, nil
}

func main() {
    log.Printf("Start lambda")
    routeHandlers := map[string]func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
        "/route1": route1Handler,
        "/route2": route2Handler,
    }

    lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
        routeHandler, ok := routeHandlers[request.Path]
        if !ok {
            return events.APIGatewayProxyResponse{
                Body:       "Invalid route",
                StatusCode: 404,
            }, nil
        }
        return routeHandler(request)
    })
}