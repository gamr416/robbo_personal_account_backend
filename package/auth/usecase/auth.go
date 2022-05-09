package usecase

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"time"
)

type AuthUseCaseImpl struct {
	auth.Gateway
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

type AuthUseCaseModule struct {
	fx.Out
	auth.UseCase
}

func SetupAuthUseCase(gateway auth.Gateway) AuthUseCaseModule {
	hashSalt := viper.GetString("auth.hash_salt")
	signingKey := []byte(viper.GetString("auth.signing_key"))
	tokenTTLTime := viper.GetDuration("auth.token_ttl")

	return AuthUseCaseModule{
		UseCase: &AuthUseCaseImpl{
			Gateway:        gateway,
			hashSalt:       hashSalt,
			signingKey:     signingKey,
			expireDuration: tokenTTLTime,
		},
	}
}

type AuthClaims struct {
	jwt.StandardClaims
	User *models.UserCore `json:"user"`
}

func (a *AuthUseCaseImpl) SignIn(email, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.Gateway.GetUser(email, password)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *AuthUseCaseImpl) SignUp(email, password string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.UserCore{
		Email:    email,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return a.Gateway.CreateUser(user)
}

func (a *AuthUseCaseImpl) ParseToken(accessToken string) (user *models.UserCore, err error) {
	return
}
