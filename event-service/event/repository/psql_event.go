package repository

import (
	"github.com/hotdog132/graphql-todo-practice/event-service/event"
	"github.com/hotdog132/graphql-todo-practice/event-service/models"
	"github.com/jinzhu/gorm"
)

type psqlEventRepository struct {
	db *gorm.DB
}

func NewPsqlEventRepository(conn *gorm.DB) event.Repository {
	return &psqlEventRepository{conn}
}

func (p *psqlEventRepository) Store(e *models.Event) error {
	return p.db.Create(e).Error
}

func (p *psqlEventRepository) Fetch(ID int) (*models.Event, error) {
	e := &models.Event{}
	err := p.db.First(e, ID).Error
	return e, err
}

func (p *psqlEventRepository) FetchAll() ([]*models.Event, error) {
	var events []*models.Event
	err := p.db.Find(&events).Error
	return events, err
}
