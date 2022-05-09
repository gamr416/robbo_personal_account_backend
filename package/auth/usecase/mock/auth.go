package mock

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) SignIn(email, password string) (string, error) {
	args := m.Called(email, password)

	return args.Get(0).(string), args.Error(1)
}

func (m *AuthUseCaseMock) SignUp(email, password string) error {
	args := m.Called(email, password)

	return args.Error(0)
}

func (m *AuthUseCaseMock) ParseToken(accessToken string) (*models.UserCore, error) {
	args := m.Called(accessToken)

	return args.Get(0).(*models.UserCore), args.Error(1)
}
