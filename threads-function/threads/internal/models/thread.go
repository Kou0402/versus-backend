package models

// Thread is thread model.
type Thread struct {
	ThreadID        string `json:"threadId"`
	CreatedDate     string `json:"createdDate"`
	CreatedTime     string `json:"createdTime"`
	CreatedUserName string `json:"createdUserName"`
	Theme1          string `json:"theme1"`
	Theme2          string `json:"theme2"`
}

type ThreadID string
