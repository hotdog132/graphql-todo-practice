package usecase

import (
	"github.com/hotdog132/graphql-todo-practice/event-service/event"
	"github.com/hotdog132/graphql-todo-practice/event-service/models"
)

type eventUsecase struct {
	eventRepo event.Repository
}

func NewEventUsecase(er event.Repository) event.Usecase {
	return &eventUsecase{er}
}

func (eu *eventUsecase) Store(e *models.Event) error {
	return eu.eventRepo.Store(e)
}

func (eu *eventUsecase) Fetch(ID int) (*models.Event, error) {
	return eu.eventRepo.Fetch(ID)
}

func (eu *eventUsecase) FetchAll() ([]*models.Event, error) {
	return eu.eventRepo.FetchAll()
}
