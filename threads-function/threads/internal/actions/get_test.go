package actions

import (
	"testing"
	"threads/internal/repositories/mock"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestThreadsGetterRun_1件のスレッドが返却される(t *testing.T) {
	threadIdParam := map[string]string{
		"threadId": "001",
	}
	request := events.APIGatewayProxyRequest{
		PathParameters: threadIdParam,
	}
	mock := &ThreadsGetter{
		ThreadsRepository: &mock.RepositoryMock{},
	}

	threads, err := mock.Run(request)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(threads))
}

func TestThreadsGetterRun_1件以上のスレッドが返却される(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	mock := &ThreadsGetter{
		ThreadsRepository: &mock.RepositoryMock{},
	}

	threads, err := mock.Run(request)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(threads))
}
