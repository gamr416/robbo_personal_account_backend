package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth/gateway/mock"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthFlow(t *testing.T) {
	gateway := new(mock.UserGatewayMock)

	uc := SetupAuthUseCase(gateway)

	var (
		email    = "email"
		password = "pass"
		user     = &models.UserCore{
			Email:    email,
			Password: "9d4e1e23bd5b727046a9e3b4b7db57bd8d6ee684", // sha1 of pass+salt
		}
	)

	// Sign Up
	gateway.On("CreateUser", user).Return(nil)
	err := uc.SignUp(email, password)
	assert.NoError(t, err)

	// Sign In (Get Auth Token)
	gateway.On("GetUser", user.Email, user.Password).Return(user, nil)
	token, err := uc.SignIn(email, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token
	//parsedUser, err := uc.ParseToken(ctx, token)
	//assert.NoError(t, err)
	//assert.Equal(t, user, parsedUser)
}
