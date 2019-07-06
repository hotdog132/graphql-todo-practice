package usecase_test

import (
	"errors"
	"testing"

	"github.com/hotdog132/graphql-todo-practice/user-service/user/usecase"

	"github.com/hotdog132/graphql-todo-practice/user-service/models"
	"github.com/hotdog132/graphql-todo-practice/user-service/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStore(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockUser := models.User{
		ID:   1,
		Name: "user-1",
	}

	t.Run("no user record", func(t *testing.T) {
		tempMockUser := mockUser
		mockRepo.On("Store", mock.AnythingOfType("*models.User")).Return(nil).Once()

		u := usecase.NewUserUsecase(mockRepo)

		err := u.Store(&tempMockUser)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("has one user record", func(t *testing.T) {
		tempMockUser := mockUser
		mockRepo.On("Store", mock.AnythingOfType("*models.User")).Return(errors.New("user duplicated")).Once()

		u := usecase.NewUserUsecase(mockRepo)

		err := u.Store(&tempMockUser)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestFetch(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockUser := models.User{
		ID:   1,
		Name: "user-1",
	}

	t.Run("Fetch one user", func(t *testing.T) {
		tempMockUser := mockUser
		mockRepo.On("Fetch", mock.AnythingOfType("int")).Return(&tempMockUser, nil).Once()

		u := usecase.NewUserUsecase(mockRepo)

		fetchUser, err := u.Fetch(mockUser.ID)

		assert.NoError(t, err)
		assert.Equal(t, &tempMockUser, fetchUser)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Fetch all users", func(t *testing.T) {
		tempMockUser1 := mockUser
		tempMockUser2 := mockUser
		users := []*models.User{&tempMockUser1, &tempMockUser2}

		mockRepo.On("FetchAll").Return(users, nil).Once()

		u := usecase.NewUserUsecase(mockRepo)

		fetchUsers, err := u.FetchAll()

		assert.NoError(t, err)
		assert.Equal(t, len(fetchUsers), 2)
		mockRepo.AssertExpectations(t)
	})
}
