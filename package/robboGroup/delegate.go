package robboGroup

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateRobboGroup(robboGroup *models.RobboGroupHttp) (robboGroupId string, err error)
	DeleteRobboGroup(robboGroupId string) (err error)
	GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroups []*models.RobboGroupHttp, err error)
	GetRobboGroupById(robboGroupId string) (robboGroup models.RobboGroupHttp, err error)
}
