package models

// Thread is thread model.
type Thread struct {
	ThreadID        string `json:"threadId" dynamodbav:"PartitionKey"`
	CreatedDate     string `json:"createdDate" dynamodbav:"CreatedDate"`
	CreatedTime     string `json:"createdTime" dynamodbav:"CreatedTime"`
	CreatedUserName string `json:"createdUserName" dynamodbav:"CreatedUserName"`
	Theme1          string `json:"theme1" dynamodbav:"Theme1"`
	Theme2          string `json:"theme2" dynamodbav:"Theme2"`
}

type ThreadID string
