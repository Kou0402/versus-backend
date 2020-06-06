package repositories

import (
	"strings"
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

	for i, t := range threads {
		threads[i].ThreadID = strings.Replace(t.ThreadID, "THREAD-", "", 1)
	}

	return threads, nil
}

func (t *ThreadsRepositoryImpl) FetchThread(threadID models.ThreadID) ([]models.Thread, error) {
	db := getDynamoSess()

	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"PartitionKey": {
				S: aws.String("THREAD-" + string(threadID)),
			},
			"SortKey": {
				S: aws.String("INFO"),
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

	threads[0].ThreadID = strings.Replace(threads[0].ThreadID, "THREAD-", "", 1)

	return threads, nil
}
