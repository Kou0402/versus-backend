package actions

import (
	"encoding/json"
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

func (g *ThreadsGetter) Run(request events.APIGatewayProxyRequest) (JSONData, error) {
	var threads []models.Thread
	var err error

	if val, ok := request.PathParameters["threadId"]; ok {
		threads, err = g.ThreadsRepository.FetchThread(models.ThreadID(val))
	}
	threads, err = g.ThreadsRepository.FetchThreads()

	resultJSON, _ := json.Marshal(threads)
	return JSONData(resultJSON), err
}
