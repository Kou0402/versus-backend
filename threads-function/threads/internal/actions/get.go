package actions

import (
	"threads/internal/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type FetchThreads struct{}

func (f *FetchThreads) Run(request events.APIGatewayProxyRequest) ([]models.Thread, error) {
	db := getDynamoSess()

	result, err := db.Scan(&dynamodb.ScanInput{
		TableName: aws.String("threads"),
	})

	if err != nil {
		return []models.Thread{}, err
	}

	var threads []models.Thread
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &threads)
	if err != nil {
		return []models.Thread{}, err
	}

	return threads, nil
}

// fetchThread は DynamoDB の Threads テーブルのデータを取得する。
func FetchThread(threadID models.ThreadID) ([]models.Thread, error) {
	db := getDynamoSess()

	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("threads"),
		Key: map[string]*dynamodb.AttributeValue{
			"threadId": {
				S: aws.String(string(threadID)),
			},
		},
	})
	if err != nil {
		return []models.Thread{}, err
	}

	items := []map[string]*dynamodb.AttributeValue{result.Item}
	var threads []models.Thread
	err = dynamodbattribute.UnmarshalListOfMaps(items, &threads)
	if err != nil {
		return []models.Thread{}, err
	}

	return threads, nil
}

func getDynamoSess() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(sess)
}
