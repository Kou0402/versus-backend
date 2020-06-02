package repositories

import (
	"posts/internal/models"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	tableName = "versus"
)

type PostsRepository interface {
	FetchPosts(threadId models.ThreadID) ([]models.Post, error)
}

type PostsRepositoryImpl struct{}

func getDynamoSess() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(sess)
}
