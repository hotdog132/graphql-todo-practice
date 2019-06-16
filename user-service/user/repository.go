package user

import "github.com/hotdog132/graphql-todo-practice/user-service/models"

// Repository user respository
type Repository interface {
	Store(*models.User) error
	Fetch(ID int) (*models.User, error)
	FetchAll() ([]*models.User, error)
}
