package user

import "github.com/hotdog132/graphql-todo-practice/user-service/models"

// Usecase usecase of user
type Usecase interface {
	Store(*models.User) error
	Fetch(ID int) (*models.User, error)
	FetchAll() ([]*models.User, error)
}
