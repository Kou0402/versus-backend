package repositories

import (
	"threads/internal/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (t *ThreadsRepositoryImpl) FetchThreads() ([]models.Thread, error) {
	db := getDynamoSess()

	partitionKey := "INFO"
	sortKey := "THREAD"
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		IndexName: aws.String(indexName),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":partitionKey": {
				S: aws.String(partitionKey),
			},
			":sortKey": {
				S: aws.String(sortKey),
			},
		},
		KeyConditionExpression: aws.String("SortKey = :partitionKey AND begins_with(PartitionKey, :sortKey)"),
	}
	result, err := db.Query(queryInput)
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

func (t *ThreadsRepositoryImpl) FetchThread(threadID models.ThreadID) ([]models.Thread, error) {
	db := getDynamoSess()

	partitionKey := "THREAD#" + threadID
	sortKey := "INFO"

	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"PartitionKey": {
				S: aws.String(string(partitionKey)),
			},
			"SortKey": {
				S: aws.String(string(sortKey)),
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
