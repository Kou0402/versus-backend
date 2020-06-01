package handler

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// Handler is Lambda handler
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintln("Hello"),
		StatusCode: 200,
	}, nil
}
