package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
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

func (s *AuthDelegateImpl) SignIn(userHttp *models.UserHttp) (accessToken, refreshToken string, err error) {
	return s.UseCase.SignIn(userHttp.ToCore())
}

func (s *AuthDelegateImpl) SignUp(userHttp *models.UserHttp) (accessToken, refreshToken string, err error) {
	return s.UseCase.SignUp(userHttp.ToCore())
}

func (s *AuthDelegateImpl) ParseToken(token string, key []byte) (claims *models.UserClaims, err error) {
	return s.UseCase.ParseToken(token, key)
}
func (s *AuthDelegateImpl) RefreshToken(token string) (newAccessToken string, err error) {
	return s.UseCase.RefreshToken(token)
}
