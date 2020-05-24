package actions

import (
	"github.com/aws/aws-lambda-go/events"
)

type JSONData string

type ActionFactory func() Action

type Action interface {
	Run(request events.APIGatewayProxyRequest) (JSONData, error)
}
