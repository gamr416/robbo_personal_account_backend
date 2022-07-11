package auth

type Delegate interface {
	SignIn(email, password string) (string, string, error)
	SignUp(email, password string) (string, string, error)
	ParseToken(token string, key []byte) (id string, err error)
	RefreshToken(token string) (newAccessToken string, err error)
}
