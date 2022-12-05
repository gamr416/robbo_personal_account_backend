package robboGroup

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateRobboGroup(robboGroupCore *models.RobboGroupCore) (newRobboGroup *models.RobboGroupCore, err error)
	UpdateRobboGroup(robboGroupCore *models.RobboGroupCore) (robboGroupUpdated *models.RobboGroupCore, err error)
	DeleteRobboGroup(robboGroupId string) (err error)
	GetAllRobboGroups() (robboGroupsCore []*models.RobboGroupCore, err error)
	GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroupsCore []*models.RobboGroupCore, err error)
	GetRobboGroupById(robboGroupId string) (robboGroupCore *models.RobboGroupCore, err error)
	SetTeacherForRobboGroup(relation *models.TeachersRobboGroupsCore) (err error)
	DeleteTeacherForRobboGroup(relation *models.TeachersRobboGroupsCore) (err error)
	DeleteRelationByRobboGroupId(robboGroupId string) (err error)
	DeleteRelationByTeacherId(teacherId string) (err error)
	GetRelationByRobboGroupId(robboGroupId string) (relations []*models.TeachersRobboGroupsCore, err error)
	GetRelationByTeacherId(teacherId string) (relations []*models.TeachersRobboGroupsCore, err error)
	SearchRobboGroupsByTitle(title string) (robboGroupsCore []*models.RobboGroupCore, err error)
}
