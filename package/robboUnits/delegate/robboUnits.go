package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	"go.uber.org/fx"
)

type RobboUnitsDelegateImpl struct {
	robboUnits.UseCase
}

type RobboUnitsDelegateModule struct {
	fx.Out
	robboUnits.Delegate
}

func SetupProjectPageDelegate(usecase robboUnits.UseCase) RobboUnitsDelegateModule {
	return RobboUnitsDelegateModule{
		Delegate: &RobboUnitsDelegateImpl{
			usecase,
		},
	}
}

func (r RobboUnitsDelegateImpl) GetAllRobboUnit(authorId string) (robboUnits []*models.RobboUnitHTTP, err error) {
	//TODO implement me
	panic("implement me")
}

func (r RobboUnitsDelegateImpl) GetRobboUnitById(robboUnitId string) (robboUnit *models.RobboUnitHTTP, err error) {
	//TODO implement me
	panic("implement me")
}

func (r RobboUnitsDelegateImpl) UpdateRobboUnit(projectPage *models.RobboUnitHTTP) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r RobboUnitsDelegateImpl) CreateRobboUnit() (robboUnitId string, err error) {
	//TODO implement me
	panic("implement me")
}
func (r RobboUnitsDelegateImpl) DeleteRobboUnit(projectId string) (err error) {
	//TODO implement me
	panic("implement me")
}
