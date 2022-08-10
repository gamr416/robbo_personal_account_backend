package robboUnits

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateRobboUnit() (robboUnitId string, err error)
	DeleteRobboUnit(projectId string) (err error)
	GetAllRobboUnit(authorId string) (robboUnits []*models.RobboUnitHTTP, err error)
	GetRobboUnitById(robboUnitId string) (robboUnit *models.RobboUnitHTTP, err error)
	UpdateRobboUnit(projectPage *models.RobboUnitHTTP) (err error)
}
