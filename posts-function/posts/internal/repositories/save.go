package repositories

import (
	"posts/internal/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (t *PostsRepositoryImpl) SavePost(post models.Post) error {
	db := getDynamoSess()

	// Convert item to dynamodb attribute.
	po, err := dynamodbattribute.MarshalMap(post)
	if err != nil {
		return err
	}

	// Create an input.
	itemInput := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      po,
	}

	// Execute.
	_, err = db.PutItem(itemInput)
	if err != nil {
		return err
	}

	return nil
}
