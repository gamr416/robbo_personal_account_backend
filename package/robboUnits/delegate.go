package robboUnits

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateRobboUnit(robboUnit *models.RobboUnitHTTP) (newRobboUnit models.RobboUnitHTTP, err error)
	UpdateRobboUnit(robboUnit *models.RobboUnitHTTP) (robboUnitUpdated models.RobboUnitHTTP, err error)
	DeleteRobboUnit(robboUnitId string) (err error)
	GetAllRobboUnit() (robboUnits []*models.RobboUnitHTTP, err error)
	GetRobboUnitById(robboUnitId string) (robboUnit models.RobboUnitHTTP, err error)
	GetRobboUnitsByUnitAdminId(unitAdminId string) (robboUnits []*models.RobboUnitHTTP, err error)
}
