package handler

import (
	"context"
	"encoding/json"

	"threads/internal/actions"

	"github.com/aws/aws-lambda-go/events"
)

var routes = map[string]actions.Action{
	"GET": &actions.GetThreads{},
}

// Handler is Lambda handler
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	action := routes[request.HTTPMethod]
	threads, err := action.Run(request)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	threadJSON, _ := json.Marshal(threads)

	return events.APIGatewayProxyResponse{
		Body:       string(threadJSON),
		StatusCode: 200,
	}, nil
}
