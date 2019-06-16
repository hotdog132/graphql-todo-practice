package event

import "github.com/hotdog132/graphql-todo-practice/event-service/models"

type Repository interface {
	Store(*models.Event) error
	Fetch(ID int) (*models.Event, error)
	FetchAll() ([]*models.Event, error)
}
