package mock

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/stretchr/testify/mock"
)

type UserStorageMock struct {
	mock.Mock
}

func (s *UserStorageMock) CreateUser(user *models.UserCore) error {
	args := s.Called(user)

	return args.Error(0)
}

func (s *UserStorageMock) GetUser(email, password string) (*models.UserCore, error) {
	args := s.Called(email, password)
	return args.Get(0).(*models.UserCore), args.Error(1)
}
