package robboGroup

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateRobboGroup(robboGroup *models.RobboGroupCore) (robboGroupId string, err error)
	DeleteRobboGroup(robboGroupId string) (err error)
	GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroups []*models.RobboGroupCore, err error)
	GetRobboGroupById(robboGroupId string) (robboGroup *models.RobboGroupCore, err error)
	//UpdateRobboUnit(robboUnit *models.RobboGroupCore) (err error)
	SetTeacherForRobboGroup(relation *models.TeachersRobboGroupsCore) (err error)
	DeleteTeacherForRobboGroup(relation *models.TeachersRobboGroupsCore) (err error)
	DeleteRelationByRobboGroupId(robboGroupId string) (err error)
	DeleteRelationByTeacherId(teacherId string) (err error)
	GetRelationByRobboGroupId(robboGroupId string) (relations []*models.TeachersRobboGroupsCore, err error)
	GetRelationByTeacherId(teacherId string) (relations []*models.TeachersRobboGroupsCore, err error)
	SearchRobboGroupsByTitle(title string) (robboGroups []*models.RobboGroupCore, err error)
}
