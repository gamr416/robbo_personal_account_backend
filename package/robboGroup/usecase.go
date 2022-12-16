package robboGroup

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateRobboGroup(robboGroup *models.RobboGroupCore) (newRobboGroup *models.RobboGroupCore, err error)
	UpdateRobboGroup(robboGroup *models.RobboGroupCore) (robboGroupUpdated *models.RobboGroupCore, err error)
	DeleteRobboGroup(robboGroupId string) (err error)
	GetAllRobboGroups(page, pageSize int) (robboGroups []*models.RobboGroupCore, countRows int64, err error)
	GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroups []*models.RobboGroupCore, err error)
	GetRobboGroupsByUnitAdminId(unitAdminId string) (robboGroups []*models.RobboGroupCore, err error)
	GetRobboGroupById(robboGroupId string) (robboGroup *models.RobboGroupCore, err error)
	GetRobboGroupsByTeacherId(teacherId string, page, pageSize int) (robboGroups []*models.RobboGroupCore, countRows int64, err error)
	SetTeacherForRobboGroup(teacherId, robboGroupId string) (err error)
	DeleteTeacherForRobboGroup(teacherId, robboGroupId string) (err error)

	SearchRobboGroupsByTitle(title string) (robboGroups []*models.RobboGroupCore, err error)
}
