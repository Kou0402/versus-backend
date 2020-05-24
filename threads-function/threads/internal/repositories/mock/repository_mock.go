package mock

import (
	"threads/internal/models"
	"threads/internal/repositories"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
	ThreadsRepository repositories.ThreadsRepository
}

func (t *RepositoryMock) FetchThread(threadID models.ThreadID) ([]models.Thread, error) {
	thread := models.Thread{
		"001",
		"2020-01-01",
		"19:00:00",
		"test-user",
		"theme1",
		"theme2",
	}
	threads := []models.Thread{thread}
	return threads, nil
}

func (t *RepositoryMock) FetchThreads() ([]models.Thread, error) {
	thread := models.Thread{
		"001",
		"2020-01-01",
		"19:00:00",
		"test-user",
		"theme1",
		"theme2",
	}
	threads := []models.Thread{thread, thread}
	return threads, nil
}

func (t *RepositoryMock) SaveThread(models.Thread) error {
	return nil
}
