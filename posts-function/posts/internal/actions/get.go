package actions

import (
	"encoding/json"
	"posts/internal/models"
	"posts/internal/repositories"

	"github.com/aws/aws-lambda-go/events"
)

type PostsGetter struct {
	PostsRepository repositories.PostsRepository
}

func NewPostsGetter() Action {
	return &PostsGetter{
		PostsRepository: &repositories.PostsRepositoryImpl{},
	}
}

func (g *PostsGetter) Run(request events.APIGatewayProxyRequest) (JSONData, error) {
	var posts []models.Post
	var err error

	posts, err = g.PostsRepository.FetchPosts(models.ThreadID(request.QueryStringParameters["threadId"]))

	resultJSON, _ := json.Marshal(posts)
	return JSONData(resultJSON), err
}
