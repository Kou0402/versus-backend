package handler

import (
	"context"
	"encoding/json"

	"threads/internal/actions"

	"github.com/aws/aws-lambda-go/events"
)

var routes = map[string]actions.ActionFactory{
	"GET": actions.NewThreadsGetter,
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
	threads, err := action.Run(request)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	threadJSON, _ := json.Marshal(threads)

	return events.APIGatewayProxyResponse{
		Body:       string(threadJSON),
		StatusCode: 200,
		Headers:    headers,
	}, nil
}
