package usecase

import (
	"crypto/sha1"
	"errors"
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
	hashSalt              string
	accessSigningKey      []byte
	refreshSigningKey     []byte
	accessExpireDuration  time.Duration
	refreshExpireDuration time.Duration
}

type AuthUseCaseModule struct {
	fx.Out
	auth.UseCase
}

func SetupAuthUseCase(gateway auth.Gateway) AuthUseCaseModule {
	hashSalt := viper.GetString("auth.hash_salt")
	accessSigningKey := []byte(viper.GetString("auth.access_signing_key"))
	refreshSigningKey := []byte(viper.GetString("auth.refresh_signing_key"))
	accessTokenTTLTime := viper.GetDuration("auth.access_token_ttl")
	refreshTokenTTLTime := viper.GetDuration("auth.refresh_token_ttl")

	return AuthUseCaseModule{
		UseCase: &AuthUseCaseImpl{
			Gateway:               gateway,
			hashSalt:              hashSalt,
			accessSigningKey:      accessSigningKey,
			refreshSigningKey:     refreshSigningKey,
			accessExpireDuration:  accessTokenTTLTime,
			refreshExpireDuration: refreshTokenTTLTime,
		},
	}
}

func (a *AuthUseCaseImpl) SignIn(email, password string) (accessToken, refreshToken string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.Gateway.GetUser(email, password)
	if err != nil {
		return "", "", auth.ErrUserNotFound
	}

	accessToken, err = a.GenerateToken(user.ID, a.accessExpireDuration, a.accessSigningKey)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = a.GenerateToken(user.ID, a.refreshExpireDuration, a.refreshSigningKey)
	if err != nil {
		return "", "", err
	}

	return
}

func (a *AuthUseCaseImpl) SignUp(email, password string) (accessToken, refreshToken string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.UserCore{
		Email:    email,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	id, err := a.Gateway.CreateUser(user)
	if err != nil {
		return "", "", err
	}

	accessToken, err = a.GenerateToken(id, a.accessExpireDuration, a.accessSigningKey)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = a.GenerateToken(id, a.refreshExpireDuration, a.refreshSigningKey)

	return
}

func (a *AuthUseCaseImpl) ParseToken(token string) (id string, err error) {
	data, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(a.refreshSigningKey), nil
		})
	if err != nil {
		return "0", err
	}

	claims, ok := data.Claims.(*jwt.StandardClaims)
	if !ok {
		return "0", errors.New("token claims are not of type *StandardClaims")
	}

	return claims.Subject, nil
}

func (a *AuthUseCaseImpl) RefreshToken(token string) (newAccessToken, newRefreshToken string, err error) {
	id, err := a.ParseToken(token)

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	newAccessToken, err = a.GenerateToken(id, a.accessExpireDuration, a.accessSigningKey)
	if err != nil {
		return "", "", err
	}
	newRefreshToken, err = a.GenerateToken(id, a.refreshExpireDuration, a.refreshSigningKey)
	if err != nil {
		return "", "", err
	}

	return
}

func (a *AuthUseCaseImpl) GenerateToken(id string, duration time.Duration, signingKey []byte) (token string, err error) {
	claims := jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(duration * time.Second)),
		Subject:   id,
	}
	ss := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = ss.SignedString(signingKey)
	if err != nil {
		fmt.Println(err)
	}
	return
}