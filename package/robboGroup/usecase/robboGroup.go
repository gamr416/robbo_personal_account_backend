package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup"
	"go.uber.org/fx"
)

type RobboGroupUseCaseImpl struct {
	Gateway robboGroup.Gateway
}

func (r *RobboGroupUseCaseImpl) CreateRobboGroup(robboGroup *models.RobboGroupCore) (robboGroupId string, err error) {
	return r.Gateway.CreateRobboGroup(robboGroup)
}

func (r *RobboGroupUseCaseImpl) DeleteRobboGroup(robboGroupId string) (err error) {
	// TODO set robboGroupId = null for student
	return r.Gateway.DeleteRobboGroup(robboGroupId)
}

func (r *RobboGroupUseCaseImpl) GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroups []*models.RobboGroupCore, err error) {
	return r.Gateway.GetRobboGroupsByRobboUnitId(robboUnitId)
}

func (r *RobboGroupUseCaseImpl) GetRobboGroupById(robboGroupId string) (robboGroup *models.RobboGroupCore, err error) {
	robboGroup, err = r.Gateway.GetRobboGroupById(robboGroupId)
	if err != nil {
		return
	}

	//TODO get student by robbogroupid
	return
}

type RobboGroupUseCaseModule struct {
	fx.Out
	robboGroup.UseCase
}

func SetupRobboGroupUseCase(robboGroupGateway robboGroup.Gateway) RobboGroupUseCaseModule {
	return RobboGroupUseCaseModule{
		UseCase: &RobboGroupUseCaseImpl{
			robboGroupGateway,
		},
	}
}
