package repositories

import (
	"posts/internal/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

func (t *PostsRepositoryImpl) SavePost(post models.Post) error {
	db := getDynamoSess()

	u, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	post.ThreadID = "THREAD-" + post.ThreadID
	post.PostID = "POST-" + u.String()

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
