package models

// Post is Post model.
type Post struct {
	ThreadID        string `json:"threadId" dynamodbav:"PartitionKey"`
	PostID          string `json:"postId" dynamodbav:"SortKey"`
	CreatedDate     string `json:"createdDate" dynamodbav:"CreatedDate"`
	CreatedTime     string `json:"createdTime" dynamodbav:"CreatedTime"`
	CreatedUserName string `json:"createdUserName" dynamodbav:"CreatedUserName"`
	Position        string `json:"position" dynamodbav:"Position"`
	Side            string `json:"side" dynamodbav:"Side"`
	Content         string `json:"content" dynamodbav:"Content"`
}

type ThreadID string
