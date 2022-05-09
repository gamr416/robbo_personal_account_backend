package auth

type Gateway interface {
	GetUser(email, password string) (err error)
}
