package robboUnits

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateRobboUnit(robboUnitCore *models.RobboUnitCore) (newRobboUnit *models.RobboUnitCore, err error)
	UpdateRobboUnit(robboUnitCore *models.RobboUnitCore) (robboUnitUpdated *models.RobboUnitCore, err error)
	DeleteRobboUnit(robboUnitId string) (err error)
	GetAllRobboUnit(page, pageSize int) (robboUnitsCore []*models.RobboUnitCore, countRows int64, err error)
	GetRobboUnitById(robboUnitId string) (robboUnitCore *models.RobboUnitCore, err error)
}
