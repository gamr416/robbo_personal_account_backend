package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	"go.uber.org/fx"
	"log"
	"strconv"
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

func (r RobboUnitsDelegateImpl) GetAllRobboUnit(page, pageSize string) (
	robboUnits []*models.RobboUnitHTTP,
	countRows int,
	err error,
) {
	pageInt32, _ := strconv.ParseInt(page, 10, 32)
	pageSizeInt32, _ := strconv.ParseInt(pageSize, 10, 32)
	robboUnitsCore, countRowsInt64, getRobboUnitsErr := r.UseCase.GetAllRobboUnit(
		int(pageInt32),
		int(pageSizeInt32),
	)
	if getRobboUnitsErr != nil {
		err = getRobboUnitsErr
		return
	}
	countRows = int(countRowsInt64)
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

func (r RobboUnitsDelegateImpl) GetRobboUnitsByUnitAdminId(unitAdminId, page, pageSize string) (
	robboUnits []*models.RobboUnitHTTP,
	countRows int,
	err error,
) {
	pageInt32, _ := strconv.ParseInt(page, 10, 32)
	pageSizeInt32, _ := strconv.ParseInt(pageSize, 10, 32)
	robboUnitsCore, countRowsInt64, getRobboUnitsErr := r.UseCase.GetRobboUnitsByUnitAdminId(
		unitAdminId,
		int(pageInt32),
		int(pageSizeInt32),
	)
	if getRobboUnitsErr != nil {
		err = getRobboUnitsErr
		return
	}
	countRows = int(countRowsInt64)
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
