package repositories

import (
	"threads/internal/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func FetchThreads() ([]models.Thread, error) {
	db := getDynamoSess()

	result, err := db.Scan(&dynamodb.ScanInput{
		TableName: aws.String(threadsTableName),
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

func FetchThread(threadID models.ThreadID) ([]models.Thread, error) {
	db := getDynamoSess()

	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(threadsTableName),
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
