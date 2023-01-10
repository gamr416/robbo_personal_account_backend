package courses

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	GetCourseRelationsRoles() (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsUnits() (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsGroups() (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByRoleId(roleId string) (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByRobboGroupId(robboGroupId string) (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByRobboUnitId(robboUnitId string) (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByCourseId(courseId string) (courseRelations []*models.CourseRelationCore, err error)
	DeleteCourseRelationsByRoleId(roleId string) (err error)
	DeleteCourseRelationsByRobboGroupId(robboGroupId string) (err error)
	DeleteCourseRelationsByRobboUnitId(robboUnitId string) (err error)
	DeleteCourseRelationById(courseRelationId string) (id string, err error)
	CreateCourseRelation(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateAbsoluteMedia(absoluteMedia *models.AbsoluteMediaCore) (id string, err error)
	CreateMedia(media *models.MediaCore) (id string, err error)
	CreateImage(image *models.ImageCore) (id string, err error)
	CreateCourseApiMediaCollection(courseApiMediaCollection *models.CourseApiMediaCollectionCore) (id string, err error)
	CreateCourse(course *models.CourseCore) (id string, err error)
	DeleteAbsoluteMedia(courseApiMediaCollectionId string) (err error)
	DeleteMedia(courseApiMediaCollectionId string) (err error)
	DeleteImage(courseApiMediaCollectionId string) (err error)
	DeleteCourseApiMediaCollection(courseId string) (id string, err error)
	DeleteCourse(courseId string) (id string, err error)
	UpdateAbsoluteMedia(absoluteMedia *models.AbsoluteMediaCore) (err error)
	UpdateMedia(media *models.MediaCore) (err error)
	UpdateImage(image *models.ImageCore) (err error)
	UpdateCourseApiMediaCollection(courseApiMediaCollection *models.CourseApiMediaCollectionCore) (err error)
	UpdateCourse(course *models.CourseCore) (err error)
}
