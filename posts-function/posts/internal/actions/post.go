package actions

import (
	"encoding/json"
	"posts/internal/models"
	"posts/internal/repositories"

	"github.com/aws/aws-lambda-go/events"
)

type PostsPoster struct {
	PostsRepository repositories.PostsRepository
}

func NewPostsPoster() Action {
	return &PostsPoster{
		PostsRepository: &repositories.PostsRepositoryImpl{},
	}
}

func (g *PostsPoster) Run(request events.APIGatewayProxyRequest) (JSONData, error) {
	var post models.Post
	err := json.Unmarshal([]byte(request.Body), &post)
	if err != nil {
		return "", err
	}

	err = g.PostsRepository.SavePost(post)
	if err != nil {
		return "", err
	}

	posts := []models.Post{post}

	resultJSON, _ := json.Marshal(posts)
	return JSONData(resultJSON), nil
}
