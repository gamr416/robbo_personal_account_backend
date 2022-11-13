package robboGroup

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateRobboGroup(robboGroup *models.RobboGroupHTTP) (robboGroupId string, err error)
	DeleteRobboGroup(robboGroupId string) (err error)
	GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroups []*models.RobboGroupHTTP, err error)
	GetRobboGroupById(robboGroupId string) (robboGroup models.RobboGroupHTTP, err error)
	GetRobboGroupsByTeacherId(teacherId string) (robboGroups []*models.RobboGroupHTTP, err error)
	SetTeacherForRobboGroup(teacherId, robboGroupId string) (err error)
	DeleteTeacherForRobboGroup(teacherId, robboGroupId string) (err error)
	SearchRobboGroupByName(name string) (robboGroups []*models.RobboGroupHTTP, err error)
}
