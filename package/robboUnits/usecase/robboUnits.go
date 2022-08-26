package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
)

type RobboUnitsUseCaseImpl struct {
	robboUnitsGateway robboUnits.Gateway
	usersGateway      users.Gateway
}

type RobboUnitsUseCaseModule struct {
	fx.Out
	robboUnits.UseCase
}

func SetupRobboUnitsUseCase(robboUnitsGateway robboUnits.Gateway, usersGateway users.Gateway) RobboUnitsUseCaseModule {
	return RobboUnitsUseCaseModule{
		UseCase: &RobboUnitsUseCaseImpl{
			robboUnitsGateway,
			usersGateway,
		},
	}
}

func (p *RobboUnitsUseCaseImpl) CreateRobboUnit(robboUnit *models.RobboUnitCore) (robboUnitId string, err error) {
	return p.robboUnitsGateway.CreateRobboUnit(robboUnit)
}

func (p *RobboUnitsUseCaseImpl) DeleteRobboUnit(robboUnitId string) (err error) {
	return p.robboUnitsGateway.DeleteRobboUnit(robboUnitId)
}

func (p *RobboUnitsUseCaseImpl) GetAllRobboUnit() (robboUnits []*models.RobboUnitCore, err error) {
	return p.robboUnitsGateway.GetAllRobboUnit()
}

func (p *RobboUnitsUseCaseImpl) GetRobboUnitById(robboUnitId string) (robboUnit *models.RobboUnitCore, err error) {
	return p.robboUnitsGateway.GetRobboUnitById(robboUnitId)
}

func (p *RobboUnitsUseCaseImpl) GetRobboUnitsByUnitAdminId(unitAdminId string) (robboUnits []*models.RobboUnitCore, err error) {
	relations, getRelationErr := p.usersGateway.GetRelationByUnitAdminId(unitAdminId)
	if getRelationErr != nil {
		err = getRelationErr
		return
	}

	for _, relation := range relations {
		robboUnit, getRobboUnitErr := p.robboUnitsGateway.GetRobboUnitById(relation.RobboUnitId)
		if getRobboUnitErr != nil {
			err = getRelationErr
			return
		}
		robboUnits = append(robboUnits, robboUnit)
	}
	return
}

func (p *RobboUnitsUseCaseImpl) UpdateRobboUnit(robboUnit *models.RobboUnitCore) (err error) {
	return p.robboUnitsGateway.UpdateRobboUnit(robboUnit)
}
