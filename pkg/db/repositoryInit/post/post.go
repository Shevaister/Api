package post

import (
	"Api/pkg/db/connection"
	"Api/pkg/models/posts"
)

var Repository posts.PostRepository

func init() {
	Repository = posts.ProvidePostRepostiory(connection.DB)
}
