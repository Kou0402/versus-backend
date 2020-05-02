package actions

import (
	"threads/internal/models"
	"threads/internal/repositories"

	"github.com/aws/aws-lambda-go/events"
)

type ThreadsGetter struct{}

func (g *ThreadsGetter) Run(request events.APIGatewayProxyRequest) ([]models.Thread, error) {
	if val, ok := request.PathParameters["threadId"]; ok {
		return repositories.FetchThread(models.ThreadID(val))
	}
	return repositories.FetchThreads()
}
