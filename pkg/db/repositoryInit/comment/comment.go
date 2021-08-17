package comment

import (
	"Api/pkg/db/connection"
	"Api/pkg/models/comments"
)

var Repository comments.CommentRepository

func init() {
	Repository = comments.ProvideCommentRepostiory(connection.DB)
}
