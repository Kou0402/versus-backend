package repositories

import (
	"threads/internal/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (t *ThreadsRepositoryImpl) SaveThread(thread models.Thread) error {
	db := getDynamoSess()

	thread.ThreadID = "THREAD-" + thread.ThreadID

	// Convert item to dynamodb attribute.
	th, err := dynamodbattribute.MarshalMap(thread)
	if err != nil {
		return err
	}

	th["SortKey"] = &dynamodb.AttributeValue{
		S: aws.String("INFO"),
	}

	// Create an input.
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      th,
	}

	// Execute.
	_, err = db.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
