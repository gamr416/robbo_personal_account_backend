package usecase

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"time"
)

type AuthUseCaseImpl struct {
	users.Gateway
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

func SetupAuthUseCase(gateway users.Gateway) AuthUseCaseModule {
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

	user := models.UserCore{}

	switch models.Role(role) {
	case models.Student:
		student, getStudentErr := a.Gateway.GetStudent(email, password)
		if getStudentErr != nil {
			return "", "", getStudentErr
		}
		user.Id = student.Id
		user.Role = models.Student
		return
	case models.Teacher:
		teacher, getTeacherErr := a.Gateway.GetTeacher(email, password)
		if getTeacherErr != nil {
			return "", "", getTeacherErr
		}
		user.Id = teacher.Id
		user.Role = models.Teacher
		return
	case models.Parent:
		parent, getParentErr := a.Gateway.GetParent(email, password)
		if getParentErr != nil {
			return "", "", getParentErr
		}
		user.Id = parent.Id
		user.Role = models.Parent
		return
	case models.FreeListener:
		freeListener, getFreeListenerErr := a.Gateway.GetFreeListener(email, password)
		if getFreeListenerErr != nil {
			return "", "", getFreeListenerErr
		}
		user.Id = freeListener.Id
		user.Role = models.FreeListener
		return
	case models.UnitAdmin:
		unitAdmin, getUnitAdminErr := a.Gateway.GetUnitAdmin(email, password)
		if getUnitAdminErr != nil {
			return "", "", getUnitAdminErr
		}
		user.Id = unitAdmin.Id
		user.Role = models.UnitAdmin
		return
	case models.SuperAdmin:
		superAdmin, getSuperAdminErr := a.Gateway.GetSuperAdmin(email, password)
		if err != nil {
			return "", "", getSuperAdminErr
		}
		user.Id = superAdmin.Id
		user.Role = models.SuperAdmin
		return
	default:
		err = auth.ErrUserNotFound
	}

	if err != nil {
		return "", "", err
	}
	accessToken, err = a.GenerateToken(&user, a.accessExpireDuration, a.accessSigningKey)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = a.GenerateToken(&user, a.refreshExpireDuration, a.refreshSigningKey)
	if err != nil {
		return "", "", err
	}

	return
}

func (a *AuthUseCaseImpl) SignUp(userCore *models.UserCore) (accessToken, refreshToken string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(userCore.Password))
	pwd.Write([]byte(a.hashSalt))
	userCore.Password = fmt.Sprintf("%x", pwd.Sum(nil))

	switch userCore.Role {
	case models.Student:
		student := &models.StudentCore{
			UserCore: *userCore,
		}
		id, createStudentErr := a.Gateway.CreateStudent(student)
		if createStudentErr != nil {
			return "", "", createStudentErr
		}
		userCore.Id = id
		return
	case models.Teacher:
		teacher := &models.TeacherCore{
			UserCore: *userCore,
		}
		id, createTeacherErr := a.Gateway.CreateTeacher(teacher)
		if createTeacherErr != nil {
			return "", "", createTeacherErr
		}
		userCore.Id = id
		return
	case models.Parent:
		parent := &models.ParentCore{
			UserCore: *userCore,
		}
		id, createParentErr := a.Gateway.CreateParent(parent)
		if createParentErr != nil {
			return "", "", createParentErr
		}
		userCore.Id = id
		return
	case models.FreeListener:
		freeListener := &models.FreeListenerCore{
			UserCore: *userCore,
		}
		id, createFreeListenerErr := a.Gateway.CreateFreeListener(freeListener)
		if createFreeListenerErr != nil {
			return "", "", createFreeListenerErr
		}
		userCore.Id = id
		return
	default:
		err = auth.ErrUserNotFound
	}

	if err != nil {
		return "", "", err
	}

	accessToken, err = a.GenerateToken(userCore, a.accessExpireDuration, a.accessSigningKey)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = a.GenerateToken(userCore, a.refreshExpireDuration, a.refreshSigningKey)

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

	user := &models.UserCore{
		Id:   claims.Id,
		Role: claims.Role,
	}

	newAccessToken, err = a.GenerateToken(user, a.accessExpireDuration, a.accessSigningKey)
	if err != nil {
		return "", err
	}

	return
}

func (a *AuthUseCaseImpl) GenerateToken(user *models.UserCore, duration time.Duration, signingKey []byte) (token string, err error) {
	claims := models.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(duration * time.Second)),
		},
		Id:   user.Id,
		Role: user.Role,
	}
	ss := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = ss.SignedString(signingKey)
	if err != nil {
		fmt.Println(err)
	}
	return
}
