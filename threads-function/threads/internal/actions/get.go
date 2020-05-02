package actions

import (
	"threads/internal/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type GetThreads struct{}

func (g *GetThreads) Run(request events.APIGatewayProxyRequest) ([]models.Thread, error) {
	if val, ok := request.PathParameters["threadId"]; ok {
		return fetchThread(models.ThreadID(val))
	}
	return fetchThreads()
}

func fetchThreads() ([]models.Thread, error) {
	db := getDynamoSess()

	result, err := db.Scan(&dynamodb.ScanInput{
		TableName: aws.String(ThreadsTable),
	})
	if err != nil {
		return nil, err
	}

	var threads []models.Thread
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &threads)
	if err != nil {
		return nil, err
	}

	return threads, nil
}

func fetchThread(threadID models.ThreadID) ([]models.Thread, error) {
	db := getDynamoSess()

	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(ThreadsTable),
		Key: map[string]*dynamodb.AttributeValue{
			"threadId": {
				S: aws.String(string(threadID)),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	items := []map[string]*dynamodb.AttributeValue{result.Item}
	var threads []models.Thread
	err = dynamodbattribute.UnmarshalListOfMaps(items, &threads)
	if err != nil {
		return nil, err
	}

	return threads, nil
}

func getDynamoSess() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(sess)
}
