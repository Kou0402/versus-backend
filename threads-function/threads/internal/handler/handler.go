package handler

import (
	"context"
	"encoding/json"

	"threads/internal/actions"
	"threads/internal/models"

	"github.com/aws/aws-lambda-go/events"
)

var Routes = map[string]actions.Action{
	"GET": &actions.FetchThreads{},
}

// Handler is Lambda handler
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var threads []models.Thread
	var err error
	if val, ok := request.PathParameters["threadId"]; ok {
		threads, err = actions.FetchThread(models.ThreadID(val))
	} else {
		action := Routes["GET"]
		threads, err = action.Run(request)
	}
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	threadJSON, _ := json.Marshal(threads)

	return events.APIGatewayProxyResponse{
		Body:       string(threadJSON),
		StatusCode: 200,
	}, nil
}
