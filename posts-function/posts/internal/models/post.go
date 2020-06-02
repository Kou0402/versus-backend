package models

// Post is Post model.
type Post struct {
	ThreadID        string `json:"threadId" dynamodbav:"PartitionKey"`
	PostID          string `json:"postId" dynamodbav:"SortKey"`
	CreatedDate     string `json:"createdDate"`
	CreatedTime     string `json:"createdTime"`
	CreatedUserName string `json:"createdUserName"`
	Position        string `json:"position"`
	Side            string `json:"side"`
	Content         string `json:"content"`
}

type ThreadID string
