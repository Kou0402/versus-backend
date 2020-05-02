package actions

import (
	"threads/internal/models"

	"github.com/aws/aws-lambda-go/events"
)

type Action interface {
	Run(request events.APIGatewayProxyRequest) ([]models.Thread, error)
}
