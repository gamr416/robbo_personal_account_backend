package edx

type RegistrationForm struct {
	Email          string
	Username       string
	Name           string
	Password       string
	TermsOfService string
}

type UserUseCase interface {
	GetUser() (respBody []byte, err error)
	PostRegistration(postMessage RegistrationForm) (respBody []byte, err error)
	Login(email, password string) (respBody []byte, err error)
}
