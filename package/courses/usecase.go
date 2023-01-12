package courses

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	GetAccessCourseRelationsStudents() (courseRelations []*models.CourseRelationCore, err error)
	GetAccessCourseRelationsTeachers() (courseRelations []*models.CourseRelationCore, err error)
	GetAccessCourseRelationsUnitAdmins() (courseRelations []*models.CourseRelationCore, err error)
	GetAccessCourseRelationsRobboUnits() (courseRelations []*models.CourseRelationCore, err error)
	GetAccessCourseRelationsRobboGroups() (courseRelations []*models.CourseRelationCore, err error)
	GetAccessCourseRelationsByStudentId(studentId string) (courseRelations []*models.CourseRelationCore, err error)
	GetAccessCourseRelationsByTeacherId(teacherId string) (courseRelations []*models.CourseRelationCore, err error)
	GetAccessCourseRelationsByUnitAdminId(unitAdminId string) (courseRelations []*models.CourseRelationCore, err error)
	GetAccessCourseRelationsByRobboGroupId(robboGroupId string) (courseRelations []*models.CourseRelationCore, err error)
	GetAccessCourseRelationsByRobboUnitId(robboUnitId string) (courseRelations []*models.CourseRelationCore, err error)
	GetAccessCourseRelationsByCourseId(courseId string) (courseRelations []*models.CourseRelationCore, err error)
	DeleteAccessCourseRelationById(courseRelationId string) (id string, err error)
	DeleteAccessCourseRelationsByStudentId(studentId string) (err error)
	DeleteAccessCourseRelationsByTeacherId(teacherId string) (err error)
	DeleteAccessCourseRelationsByUnitAdminId(unitAdminId string) (err error)
	DeleteAccessCourseRelationsByRobboGroupId(robboGroupId string) (err error)
	DeleteAccessCourseRelationsByRobboUnitId(robboUnitId string) (err error)
	CreateAccessCourseRelationRobboGroup(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateAccessCourseRelationRobboUnit(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateAccessCourseRelationStudent(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateAccessCourseRelationTeacher(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateAccessCourseRelationUnitAdmin(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateCourse(course *models.CourseCore) (id string, err error)
	DeleteCourse(courseId string) (err error)
	UpdateCourse(course *models.CourseCore) (err error)
}
