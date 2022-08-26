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

func (s *AuthDelegateImpl) SignIn(email, password string, role uint) (accessToken, refreshToken string, err error) {
	return s.UseCase.SignIn(email, password, role)
}

func (s *AuthDelegateImpl) SignUp(userHttp *models.UserHttp) (accessToken, refreshToken string, err error) {
	userCore := userHttp.ToCore()
	return s.UseCase.SignUp(userCore)
}

func (s *AuthDelegateImpl) ParseToken(token string, key []byte) (claims *models.UserClaims, err error) {
	return s.UseCase.ParseToken(token, key)
}
func (s *AuthDelegateImpl) RefreshToken(refreshToken string) (newAccessToken string, err error) {
	return s.UseCase.RefreshToken(refreshToken)
}
