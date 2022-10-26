package robboGroup

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateRobboGroup(robboGroup *models.RobboGroupCore) (robboGroupId string, err error)
	DeleteRobboGroup(robboGroupId string) (err error)
	GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroups []*models.RobboGroupCore, err error)
	GetRobboGroupById(robboGroupId string) (robboGroup *models.RobboGroupCore, err error)

	SetTeacherForRobboGroup(teacherId, robboGroupId string) (err error)
	DeleteUnitAdminForRobboUnit(teacherId, robboGroupId string) (err error)
	GetRobboGroupsByTeacherId(teacherId string) (robboGroups []*models.RobboGroupCore, err error)
	SearchRobboGroupsByTitle(title string) (robboGroups []*models.RobboGroupCore, err error)
}
