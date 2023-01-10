package courses

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	GetCourseRelationsRoles() (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsUnits() (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsGroups() (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByRoleId(roleId string) (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByRobboGroupId(robboGroupId string) (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByRobboUnitId(robboUnitId string) (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByCourseId(courseId string) (courseRelations []*models.CourseRelationCore, err error)
	DeleteCourseRelationById(courseRelationId string) (id string, err error)
	DeleteCourseRelationsByRoleId(roleId string) (err error)
	DeleteCourseRelationsByRobboGroupId(robboGroupId string) (err error)
	DeleteCourseRelationsByRobboUnitId(robboUnitId string) (err error)
	CreateCourseRelationGroup(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateCourseRelationUnit(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateCourseRelationRole(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateCourse(course *models.CourseCore) (id string, err error)
	DeleteCourse(courseId string) (err error)
	UpdateCourse(course *models.CourseCore) (err error)
}
