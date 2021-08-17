package user

import (
	"Api/pkg/db/connection"
	"Api/pkg/models/users"
)

var Repository users.UserRepository

func init() {
	Repository = users.ProvideUserRepostiory(connection.DB)
}
