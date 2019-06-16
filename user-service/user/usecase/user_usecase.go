package usecase

import (
	"github.com/hotdog132/graphql-todo-practice/user-service/models"
	"github.com/hotdog132/graphql-todo-practice/user-service/user"
)

// UserUsecase user use case implementation
type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase(ur user.Repository) user.Usecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) Store(u *models.User) error {
	return uu.userRepo.Store(u)
}

func (uu *userUsecase) Fetch(ID int) (*models.User, error) {
	return uu.userRepo.Fetch(ID)
}

func (uu *userUsecase) FetchAll() ([]*models.User, error) {
	return uu.userRepo.FetchAll()
}
