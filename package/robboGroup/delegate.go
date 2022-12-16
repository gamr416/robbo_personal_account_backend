package robboGroup

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateRobboGroup(robboGroup *models.RobboGroupHTTP) (newRobboGroup models.RobboGroupHTTP, err error)
	UpdateRobboGroup(robboGroup *models.RobboGroupHTTP) (robboGroupUpdated models.RobboGroupHTTP, err error)
	DeleteRobboGroup(robboGroupId string) (err error)
	GetAllRobboGroups(page, pageSize string) (robboGroups []*models.RobboGroupHTTP, countRows int, err error)
	GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroups []*models.RobboGroupHTTP, err error)
	GetRobboGroupsByUnitAdminId(unitAdminId string) (robboGroups []*models.RobboGroupHTTP, err error)
	GetRobboGroupById(robboGroupId string) (robboGroup models.RobboGroupHTTP, err error)
	GetRobboGroupsByTeacherId(teacherId, page, pageSize string) (robboGroups []*models.RobboGroupHTTP, countRows int, err error)
	SetTeacherForRobboGroup(teacherId, robboGroupId string) (err error)
	DeleteTeacherForRobboGroup(teacherId, robboGroupId string) (err error)
	SearchRobboGroupByName(name string) (robboGroups []*models.RobboGroupHTTP, err error)
}
