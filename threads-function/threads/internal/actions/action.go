package actions

import (
	"threads/internal/models"

	"github.com/aws/aws-lambda-go/events"
)

type ActionFactory func() Action

type Action interface {
	Run(request events.APIGatewayProxyRequest) ([]models.Thread, error)
}
