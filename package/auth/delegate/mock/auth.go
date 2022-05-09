package mock

import (
	"github.com/stretchr/testify/mock"
)

type AuthDelegateMock struct {
	mock.Mock
}

func (m *AuthDelegateMock) SignUp(email, password string) error {
	args := m.Called(email, password)

	return args.Error(0)
}

func (m *AuthDelegateMock) SignIn(email, password string) (string, error) {
	args := m.Called(email, password)

	return args.Get(0).(string), args.Error(1)
}
