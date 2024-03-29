package repositories

import (
	"threads/internal/models"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	tableName = "versus"
	indexName = "SortKey-index"
)

type ThreadsRepository interface {
	FetchThread(threadID models.ThreadID) ([]models.Thread, error)
	FetchThreads() ([]models.Thread, error)
	SaveThread(thread models.Thread) error
}

type ThreadsRepositoryImpl struct{}

func getDynamoSess() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(sess)
}
