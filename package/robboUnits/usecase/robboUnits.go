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

func (p *RobboUnitsUseCaseImpl) SearchRobboUnitsByName(name string) (robboUnits []*models.RobboUnitCore, err error) {
	nameCondition := name + "%"
	return p.robboUnitsGateway.SearchRobboUnitByName(nameCondition)

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

func (p *RobboUnitsUseCaseImpl) CreateRobboUnit(robboUnit *models.RobboUnitCore) (newRobboUnit *models.RobboUnitCore, err error) {
	return p.robboUnitsGateway.CreateRobboUnit(robboUnit)
}

func (p *RobboUnitsUseCaseImpl) DeleteRobboUnit(robboUnitId string) (err error) {
	return p.robboUnitsGateway.DeleteRobboUnit(robboUnitId)
}

func (p *RobboUnitsUseCaseImpl) GetAllRobboUnit(page, pageSize int) (robboUnits []*models.RobboUnitCore, countRows int64, err error) {
	return p.robboUnitsGateway.GetAllRobboUnit(page, pageSize)
}

func (p *RobboUnitsUseCaseImpl) GetRobboUnitById(robboUnitId string) (robboUnit *models.RobboUnitCore, err error) {
	return p.robboUnitsGateway.GetRobboUnitById(robboUnitId)
}

func (p *RobboUnitsUseCaseImpl) GetRobboUnitsByUnitAdminId(unitAdminId string, page, pageSize int) (
	robboUnits []*models.RobboUnitCore,
	countRows int64,
	err error,
) {
	relations, countRows, getRelationErr := p.usersGateway.GetRelationByUnitAdminId(unitAdminId, page, pageSize)
	if getRelationErr != nil {
		err = getRelationErr
		return
	}

	for _, relation := range relations {
		robboUnit, getRobboUnitErr := p.robboUnitsGateway.GetRobboUnitById(relation.RobboUnitId)
		if getRobboUnitErr != nil {
			return []*models.RobboUnitCore{}, 0, getRobboUnitErr
		}
		robboUnits = append(robboUnits, robboUnit)
	}
	return
}

func (p *RobboUnitsUseCaseImpl) UpdateRobboUnit(robboUnit *models.RobboUnitCore) (robboUnitUpdated *models.RobboUnitCore, err error) {
	return p.robboUnitsGateway.UpdateRobboUnit(robboUnit)
}
