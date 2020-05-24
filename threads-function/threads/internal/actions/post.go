package actions

import (
	"encoding/json"
	"threads/internal/models"
	"threads/internal/repositories"

	"github.com/aws/aws-lambda-go/events"
)

type ThreadsPoster struct {
	ThreadsRepository repositories.ThreadsRepository
}

func NewThreadsPoster() Action {
	return &ThreadsPoster{
		ThreadsRepository: &repositories.ThreadsRepositoryImpl{},
	}
}

func (g *ThreadsPoster) Run(request events.APIGatewayProxyRequest) ([]models.Thread, error) {
	var thread models.Thread
	err := json.Unmarshal([]byte(request.Body), &thread)
	if err != nil {
		return nil, err
	}

	err = g.ThreadsRepository.SaveThread(thread)
	if err != nil {
		return nil, err
	}

	threads := []models.Thread{thread}

	return threads, nil
}
