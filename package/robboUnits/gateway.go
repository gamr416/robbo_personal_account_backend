package robboUnits

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateRobboUnit() (robboUnitId string, err error)
	DeleteRobboUnit(projectId string) (err error)
	GetAllRobboUnit(authorId string) (robboUnits []*models.RobboUnitCore, err error)
	GetRobboUnitById(robboUnitId string) (robboUnit *models.RobboUnitCore, err error)
	UpdateRobboUnit(projectPage *models.RobboUnitCore) (err error)
}
