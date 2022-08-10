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

func SetupProjectPageUseCase(robboUnitsGateway robboUnits.Gateway) RobboUnitsUseCaseModule {
	return RobboUnitsUseCaseModule{
		UseCase: &RobboUnitsUseCaseImpl{
			robboUnitsGateway,
		},
	}
}

func (p *RobboUnitsUseCaseImpl) CreateRobboUnit() (robboUnitId string, err error) {
	return
}

func (p *RobboUnitsUseCaseImpl) DeleteRobboUnit(projectId string) (err error) {
	return
}

func (p *RobboUnitsUseCaseImpl) GetAllRobboUnit(authorId string) (robboUnits []*models.RobboUnitCore, err error) {
	return
}

func (p *RobboUnitsUseCaseImpl) GetRobboUnitById(robboUnitId string) (robboUnit *models.RobboUnitCore, err error) {
	return
}

func (p *RobboUnitsUseCaseImpl) UpdateRobboUnit(projectPage *models.RobboUnitCore) (err error) {
	return
}
