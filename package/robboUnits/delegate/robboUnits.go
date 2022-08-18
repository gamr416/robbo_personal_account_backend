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

func SetupRobboUnitsDelegate(usecase robboUnits.UseCase) RobboUnitsDelegateModule {
	return RobboUnitsDelegateModule{
		Delegate: &RobboUnitsDelegateImpl{
			usecase,
		},
	}
}

func (r RobboUnitsDelegateImpl) GetAllRobboUnit() (robboUnits []*models.RobboUnitHTTP, err error) {
	robboUnitsCore, getRobboUnits := r.UseCase.GetAllRobboUnit()
	if getRobboUnits != nil {
		err = getRobboUnits
		return
	}
	for _, robboUnitCore := range robboUnitsCore {
		var robboUnitTemp models.RobboUnitHTTP
		robboUnitTemp.FromCore(robboUnitCore)
		robboUnits = append(robboUnits, &robboUnitTemp)
	}
	return
}

func (r RobboUnitsDelegateImpl) GetRobboUnitById(robboUnitId string) (robboUnit models.RobboUnitHTTP, err error) {
	robboUnitCore, getRobboUnitErr := r.UseCase.GetRobboUnitById(robboUnitId)
	if getRobboUnitErr != nil {
		err = getRobboUnitErr
		return
	}
	robboUnit.FromCore(robboUnitCore)
	return
}

func (r RobboUnitsDelegateImpl) UpdateRobboUnit(robboUnit *models.RobboUnitHTTP) (err error) {
	robboUnitCore := robboUnit.ToCore()
	return r.UseCase.UpdateRobboUnit(robboUnitCore)
}

func (r RobboUnitsDelegateImpl) CreateRobboUnit(robboUnit *models.RobboUnitHTTP) (robboUnitId string, err error) {
	robboUnitCore := robboUnit.ToCore()
	return r.UseCase.CreateRobboUnit(robboUnitCore)
}
func (r RobboUnitsDelegateImpl) DeleteRobboUnit(robboUnitId string) (err error) {
	return r.UseCase.DeleteRobboUnit(robboUnitId)
}
