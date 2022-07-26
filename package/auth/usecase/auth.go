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

func (a *AuthUseCaseImpl) SignIn(email, passwordIn string, role uint) (accessToken, refreshToken string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(passwordIn))
	pwd.Write([]byte(a.hashSalt))
	password := fmt.Sprintf("%x", pwd.Sum(nil))
	var userToken models.UserToken

	switch role {
	case 0:
		user, err := a.usersUseCase.GetStudent(email, password)
		if err != nil {
			return "", "", err
		}
		userToken.ID = user.ID
		userToken.Role = user.Role
	case 1:
		user, err := a.usersUseCase.GetTeacher(email, password)
		if err != nil {
			return "", "", err
		}
		userToken.ID = user.ID
		userToken.Role = user.Role
	case 2:
		user, err := a.usersUseCase.GetParent(email, password)
		if err != nil {
			return "", "", err
		}
		userToken.ID = user.ID
		userToken.Role = user.Role
	case 3:
		return "", "", errors.New("error no permissions")
	case 4:
		user, err := a.usersUseCase.GetUnitAdmin(email, password)
		if err != nil {
			return "", "", err
		}
		userToken.ID = user.ID
		userToken.Role = user.Role
	case 5:
		user, err := a.usersUseCase.GetSuperAdmin(email, password)
		if err != nil {
			return "", "", err
		}
		userToken.ID = user.ID
		userToken.Role = user.Role
	default:
		fmt.Println("error")
	}

	if err != nil {
		return
	}
	accessToken, err = a.GenerateToken(&userToken, a.accessExpireDuration, a.accessSigningKey)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = a.GenerateToken(&userToken, a.refreshExpireDuration, a.refreshSigningKey)
	if err != nil {
		return "", "", err
	}

	return
}

func (a *AuthUseCaseImpl) SignUp(userCore *models.User, role models.Role) (accessToken, refreshToken string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(userCore.Password))
	pwd.Write([]byte(a.hashSalt))

	userCore.Password = fmt.Sprintf("%x", pwd.Sum(nil))
<<<<<<< HEAD

	id, err := a.Gateway.CreateUser(userCore)
=======
	var userToken models.UserToken
	switch userCore.Role {
	case 0:
		student := &models.StudentCore{
			User: *userCore,
		}
		id, err := a.usersUseCase.CreateStudent(student)
		if err != nil {
			return "", "", err
		}
		userToken.ID = id
		userToken.Role = userCore.Role
	case 1:
		teacher := &models.TeacherCore{
			User: *userCore,
		}
		id, err := a.usersUseCase.CreateTeacher(teacher)
		if err != nil {
			return "", "", err
		}
		userToken.ID = id
		userToken.Role = userCore.Role
	case 2:
		parent := &models.ParentCore{
			User: *userCore,
		}
		id, err := a.usersUseCase.CreateParent(parent)
		if err != nil {
			return "", "", err
		}
		userToken.ID = id
		userToken.Role = userCore.Role
	case 3:
		return "", "", errors.New("error no permissions")
	case 4:
		roleAdmin := models.Role(5)
		if role == roleAdmin {
			unitAdmin := &models.UnitAdminCore{
				User: *userCore,
			}
			id, err := a.usersUseCase.CreateUnitAdmin(unitAdmin)
			if err != nil {
				return "", "", err
			}
			userToken.ID = id
			userToken.Role = userCore.Role
		} else {
			return "", "", errors.New("error no permissions")
		}
	default:
		fmt.Println("error")
	}
>>>>>>> b51413c19b53a2b40776b2746be8d694d6f8e40e
	if err != nil {
		return "", "", err
	}

	accessToken, err = a.GenerateToken(&userToken, a.accessExpireDuration, a.accessSigningKey)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = a.GenerateToken(&userToken, a.refreshExpireDuration, a.refreshSigningKey)

	return
}

func (a *AuthUseCaseImpl) ParseToken(token string, key []byte) (claims *models.UserClaims, err error) {
	data, err := jwt.ParseWithClaims(token, &models.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})

	if err != nil {
		return &models.UserClaims{}, err
	}

	claims, ok := data.Claims.(*models.UserClaims)
	if !ok {
		return &models.UserClaims{}, auth.ErrInvalidTypeClaims
	}
	return
}

func (a *AuthUseCaseImpl) RefreshToken(token string) (newAccessToken string, err error) {
	claims, err := a.ParseToken(token, a.refreshSigningKey)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	user := &models.UserToken{
		ID:   claims.Id,
		Role: claims.Role,
	}

	newAccessToken, err = a.GenerateToken(user, a.accessExpireDuration, a.accessSigningKey)
	if err != nil {
		return "", err
	}

	return
}

func (a *AuthUseCaseImpl) GenerateToken(user *models.UserToken, duration time.Duration, signingKey []byte) (token string, err error) {
	claims := models.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(duration * time.Second)),
		},
		Id:   user.ID,
		Role: user.Role,
	}
	ss := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = ss.SignedString(signingKey)
	if err != nil {
		fmt.Println(err)
	}
	return
}
