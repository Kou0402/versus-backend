package actions

import (
	"threads/internal/models"

	"github.com/aws/aws-lambda-go/events"
)

var ThreadsTable = "threads"

type Action interface {
	Run(request events.APIGatewayProxyRequest) ([]models.Thread, error)
}
