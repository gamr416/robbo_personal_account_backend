package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
	"log"
)

type UsersUseCaseImpl struct {
	users.Gateway
}

type UsersUseCaseModule struct {
	fx.Out
	users.UseCase
}

func SetupUsersUseCase(gateway users.Gateway) UsersUseCaseModule {
	return UsersUseCaseModule{
		UseCase: &UsersUseCaseImpl{
			Gateway: gateway,
		},
	}
}

func (p *UsersUseCaseImpl) GetUsersByRole(role *models.Role) (user *[]models.UserCore, err error) {
	return p.Gateway.GetUsersByRole(role)
}

func (p *UsersUseCaseImpl) UpdateUser(user *models.UserCore) (err error) {
	err = p.Gateway.UpdateUser(user)
	if err != nil {
		log.Println("Error update User")
		return
	}
	return
}

func (p *UsersUseCaseImpl) DeleteUser(userId int) (err error) {
	return p.Gateway.DeleteUser(userId)
}

func (p *UsersUseCaseImpl) GetUsersById(userId int) (user *models.UserCore, err error) {
	user, err = p.Gateway.GetUserById(userId)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (p *UsersUseCaseImpl) CreateUser(user *models.UserCore) (id string, err error) {
	return p.Gateway.CreateUser(user)
}
