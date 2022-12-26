package robboUnits

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateRobboUnit(robboUnit *models.RobboUnitCore) (newRobboUnit *models.RobboUnitCore, err error)
	UpdateRobboUnit(robboUnit *models.RobboUnitCore) (robboUnitUpdated *models.RobboUnitCore, err error)
	DeleteRobboUnit(robboUnitId string) (err error)
	GetAllRobboUnit(page, pageSize int) (robboUnits []*models.RobboUnitCore, countRows int64, err error)
	GetRobboUnitById(robboUnitId string) (robboUnit *models.RobboUnitCore, err error)
	GetRobboUnitsByUnitAdminId(unitAdminId string, page, pageSize int) (robboUnits []*models.RobboUnitCore, countRows int64, err error)
}
