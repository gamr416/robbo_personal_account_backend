package auth

type Delegate interface {
	SignIn(email, password string) (err error)
}
