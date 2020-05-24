package handler

import (
	"context"

	"threads/internal/actions"

	"github.com/aws/aws-lambda-go/events"
)

var routes = map[string]actions.ActionFactory{
	"GET":  actions.NewThreadsGetter,
	"POST": actions.NewThreadsPoster,
}

// CORS compatible
var headers = map[string]string{
	"Content-Type":                "application/json",
	"Access-Control-Allow-Origin": "*",
}

// Handler is Lambda handler
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	factory := routes[request.HTTPMethod]
	action := factory()
	jsonData, err := action.Run(request)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonData),
		StatusCode: 200,
		Headers:    headers,
	}, nil
}
