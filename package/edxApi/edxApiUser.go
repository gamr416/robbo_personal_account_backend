package edxApi

type RegistrationForm struct {
	Email            string
	Username         string
	Name             string
	Password         string
	Terms_of_service string
}

type EdxApiUser interface {
	GetUser() (respBody []byte, err error)
	PostRegistration(postMessage RegistrationForm) (respBody []byte, err error)
	Login(email, password string) (respBody []byte, err error)
}
