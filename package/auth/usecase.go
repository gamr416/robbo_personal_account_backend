package auth

type UseCase interface {
	GetUser(email, password string) (err error)
}
