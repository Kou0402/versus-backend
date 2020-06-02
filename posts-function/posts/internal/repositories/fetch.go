package repositories

import (
	"posts/internal/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (t *PostsRepositoryImpl) FetchPosts(threadId models.ThreadID) ([]models.Post, error) {
	db := getDynamoSess()

	partitionKey := "THREAD#" + "001"
	sortKey := "POST#"
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":threadId": {
				S: aws.String(partitionKey),
			},
			":postId": {
				S: aws.String(sortKey),
			},
		},
		KeyConditionExpression: aws.String("PartitionKey = :threadId AND begins_with(SortKey, :postId)"),
	}
	result, err := db.Query(queryInput)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
