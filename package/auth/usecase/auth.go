package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
)

type AuthUseCaseImpl struct {
	auth.Gateway
}

type AuthUseCaseModule struct {
	fx.Out
	auth.UseCase
}

func SetupAuthUseCase(gateway auth.Gateway) AuthUseCaseModule {
	return AuthUseCaseModule{
		UseCase: &AuthUseCaseImpl{Gateway: gateway},
	}
}

func (s *AuthUseCaseImpl) SignIn(email, password string) error {
	return s.Gateway.GetUser(email, password)
}

func (s *AuthUseCaseImpl) SignUp(user *models.UserCore) error {
	return s.Gateway.CreateUser(user)
}

func (s *AuthUseCaseImpl) ParseToken(accessToken string) (user *models.UserCore, err error) {
	return
}
