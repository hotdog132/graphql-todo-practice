package repository

import (
	"github.com/hotdog132/graphql-todo-practice/user-service/models"
	"github.com/hotdog132/graphql-todo-practice/user-service/user"
	"github.com/jinzhu/gorm"
)

type psqlUserRepository struct {
	db *gorm.DB
}

func NewPsqlUserRepository(conn *gorm.DB) user.Repository {
	return &psqlUserRepository{conn}
}

func (p *psqlUserRepository) Store(u *models.User) error {
	return p.db.FirstOrCreate(u, u).Error
}

func (p *psqlUserRepository) Fetch(ID int) (*models.User, error) {
	u := &models.User{}
	err := p.db.First(u, ID).Error
	return u, err
}

func (p *psqlUserRepository) FetchAll() ([]*models.User, error) {
	var users []*models.User
	err := p.db.Find(&users).Error
	return users, err
}
