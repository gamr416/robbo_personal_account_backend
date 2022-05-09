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

func (s *AuthDelegateImpl) SignIn(email, password string) (err error) {
	return nil
}

func (s *AuthDelegateImpl) SignUp(email, password string) error {
	return nil
}
