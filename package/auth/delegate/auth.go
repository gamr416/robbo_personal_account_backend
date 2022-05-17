package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"go.uber.org/fx"
)

type AuthDelegateImpl struct {
	auth.UseCase
}

type AuthDelegateModule struct {
	fx.Out
	auth.Delegate
}

func SetupAuthDelegate(usecase auth.UseCase) AuthDelegateModule {
	return AuthDelegateModule{
		Delegate: &AuthDelegateImpl{usecase},
	}
}

func (s *AuthDelegateImpl) SignIn(email, password string) (accessToken, refreshToken string, err error) {
	return s.UseCase.SignIn(email, password)
}

func (s *AuthDelegateImpl) SignUp(email, password string) (accessToken, refreshToken string, err error) {
	return s.UseCase.SignUp(email, password)
}

func (s *AuthDelegateImpl) ParseToken(token string) (id string, err error) {
	return s.UseCase.ParseToken(token)
}
func (s *AuthDelegateImpl) RefreshToken(token string) (newAccessToken, newRefreshToken string, err error) {
	return s.UseCase.RefreshToken(token)
}
