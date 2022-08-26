package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
)

type RobboGroupUseCaseImpl struct {
	robboGroupGateway robboGroup.Gateway
	usersGateway      users.Gateway
}

func (r *RobboGroupUseCaseImpl) CreateRobboGroup(robboGroup *models.RobboGroupCore) (robboGroupId string, err error) {
	return r.robboGroupGateway.CreateRobboGroup(robboGroup)
}

func (r *RobboGroupUseCaseImpl) DeleteRobboGroup(robboGroupId string) (err error) {
	// TODO set robboGroupId = null for student
	return r.robboGroupGateway.DeleteRobboGroup(robboGroupId)
}

func (r *RobboGroupUseCaseImpl) GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroups []*models.RobboGroupCore, err error) {
	return r.robboGroupGateway.GetRobboGroupsByRobboUnitId(robboUnitId)
}

func (r *RobboGroupUseCaseImpl) GetRobboGroupById(robboGroupId string) (robboGroup *models.RobboGroupCore, err error) {
	robboGroup, err = r.robboGroupGateway.GetRobboGroupById(robboGroupId)
	if err != nil {
		return
	}
	students, getStudentErr := r.usersGateway.GetStudentsByRobboGroupId(robboGroup.Id)
	if getStudentErr != nil {
		err = getStudentErr
		return
	}
	robboGroup.Students = students
	return
}

type RobboGroupUseCaseModule struct {
	fx.Out
	robboGroup.UseCase
}

func SetupRobboGroupUseCase(robboGroupGateway robboGroup.Gateway, usersGateway users.Gateway) RobboGroupUseCaseModule {
	return RobboGroupUseCaseModule{
		UseCase: &RobboGroupUseCaseImpl{
			robboGroupGateway,
			usersGateway,
		},
	}
}
