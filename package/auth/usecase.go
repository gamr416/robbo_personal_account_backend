package auth

type UseCase interface {
	SignIn(email, password string) (string, string, error)
	SignUp(email, password string) (string, string, error)
	ParseToken(token string, key []byte) (string, error)
	RefreshToken(token string) (string, error)
}
