package robboUnits

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateRobboUnit(robboUnit *models.RobboUnitHTTP) (robboUnitId string, err error)
	DeleteRobboUnit(robboUnitId string) (err error)
	GetAllRobboUnit() (robboUnits []*models.RobboUnitHTTP, err error)
	GetRobboUnitById(robboUnitId string) (robboUnit *models.RobboUnitHTTP, err error)
	UpdateRobboUnit(robboUnit *models.RobboUnitHTTP) (err error)
}
