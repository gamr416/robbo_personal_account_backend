package auth

type Delegate interface {
	SignIn(email, password string) (token string, err error)
	SignUp(email, password string) (err error)
}
