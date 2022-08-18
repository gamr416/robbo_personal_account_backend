package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	"go.uber.org/fx"
)

type RobboUnitsUseCaseImpl struct {
	robboUnits.Gateway
}

type RobboUnitsUseCaseModule struct {
	fx.Out
	robboUnits.UseCase
}

func SetupRobboUnitsUseCase(robboUnitsGateway robboUnits.Gateway) RobboUnitsUseCaseModule {
	return RobboUnitsUseCaseModule{
		UseCase: &RobboUnitsUseCaseImpl{
			robboUnitsGateway,
		},
	}
}

func (p *RobboUnitsUseCaseImpl) CreateRobboUnit(robboUnit *models.RobboUnitCore) (robboUnitId string, err error) {
	return p.Gateway.CreateRobboUnit(robboUnit)
}

func (p *RobboUnitsUseCaseImpl) DeleteRobboUnit(robboUnitId string) (err error) {
	return p.Gateway.DeleteRobboUnit(robboUnitId)
}

func (p *RobboUnitsUseCaseImpl) GetAllRobboUnit() (robboUnits []*models.RobboUnitCore, err error) {
	return p.Gateway.GetAllRobboUnit()
}

func (p *RobboUnitsUseCaseImpl) GetRobboUnitById(robboUnitId string) (robboUnit *models.RobboUnitCore, err error) {
	return p.Gateway.GetRobboUnitById(robboUnitId)
}

func (p *RobboUnitsUseCaseImpl) UpdateRobboUnit(robboUnit *models.RobboUnitCore) (err error) {
	return p.Gateway.UpdateRobboUnit(robboUnit)
}
