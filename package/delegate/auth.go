package delegate

import "github.com/skinnykaen/robbo_student_personal_account.git/package/usecase"

type AuthDelegate struct {
	usecase.AuthUseCase
}

func (s *AuthDelegate) SignIn(email, password string) (err error) {
	return nil
}
