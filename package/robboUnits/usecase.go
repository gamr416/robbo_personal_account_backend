package robboUnits

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateRobboUnit(robboUnit *models.RobboUnitCore) (robboUnitId string, err error)
	DeleteRobboUnit(robboUnitId string) (err error)
	GetAllRobboUnit() (robboUnits []*models.RobboUnitCore, err error)
	GetRobboUnitById(robboUnitId string) (robboUnit *models.RobboUnitCore, err error)
	GetRobboUnitsByUnitAdminId(unitAdminId string) (robboUnits []*models.RobboUnitCore, err error)
	UpdateRobboUnit(robboUnit *models.RobboUnitCore) (err error)
}
