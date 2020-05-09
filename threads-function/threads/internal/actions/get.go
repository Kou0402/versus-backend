package actions

import (
	"threads/internal/models"
	"threads/internal/repositories"

	"github.com/aws/aws-lambda-go/events"
)

type ThreadsGetter struct {
	ThreadsRepository repositories.ThreadsRepository
}

func NewThreadsGetter() Action {
	return &ThreadsGetter{
		ThreadsRepository: &repositories.ThreadsRepositoryImpl{},
	}
}

func (g *ThreadsGetter) Run(request events.APIGatewayProxyRequest) ([]models.Thread, error) {
	if val, ok := request.PathParameters["threadId"]; ok {
		return g.ThreadsRepository.FetchThread(models.ThreadID(val))
	}
	return g.ThreadsRepository.FetchThreads()
}
