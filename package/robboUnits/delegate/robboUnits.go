package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	"go.uber.org/fx"
	"log"
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

func (r RobboUnitsDelegateImpl) GetRobboUnitsByUnitAdminId(unitAdminId string) (robboUnits []*models.RobboUnitHTTP, err error) {
	robboUnitsCore, getRobboUnits := r.UseCase.GetRobboUnitsByUnitAdminId(unitAdminId)
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

func (r RobboUnitsDelegateImpl) UpdateRobboUnit(robboUnit *models.RobboUnitHTTP) (robboUnitUpdated models.RobboUnitHTTP, err error) {
	robboUnitCore := robboUnit.ToCore()
	robboUnitUpdatedCore, err := r.UseCase.UpdateRobboUnit(robboUnitCore)
	if err != nil {
		log.Println(err)
		return
	}
	robboUnitUpdated.FromCore(robboUnitUpdatedCore)
	return
}

func (r RobboUnitsDelegateImpl) CreateRobboUnit(robboUnit *models.RobboUnitHTTP) (newRobboUnit models.RobboUnitHTTP, err error) {
	robboUnitCore := robboUnit.ToCore()
	newRobboUnitCore, err := r.UseCase.CreateRobboUnit(robboUnitCore)
	if err != nil {
		log.Println(err)
		return
	}
	newRobboUnit.FromCore(newRobboUnitCore)
	return
}
func (r RobboUnitsDelegateImpl) DeleteRobboUnit(robboUnitId string) (err error) {
	return r.UseCase.DeleteRobboUnit(robboUnitId)
}
