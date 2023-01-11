package courses

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	GetCourseRelationsStudents() (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsTeachers() (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsUnitAdmins() (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsUnits() (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsGroups() (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByStudentId(studentId string) (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByTeacherId(teacherId string) (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByUnitAdminId(unitAdminId string) (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByRobboGroupId(robboGroupId string) (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByRobboUnitId(robboUnitId string) (courseRelations []*models.CourseRelationCore, err error)
	GetCourseRelationsByCourseId(courseId string) (courseRelations []*models.CourseRelationCore, err error)
	DeleteCourseRelationById(courseRelationId string) (id string, err error)
	DeleteCourseRelationsByStudentId(studentId string) (err error)
	DeleteCourseRelationsByTeacherId(teacherId string) (err error)
	DeleteCourseRelationsByUnitAdminId(unitAdminId string) (err error)
	DeleteCourseRelationsByRobboGroupId(robboGroupId string) (err error)
	DeleteCourseRelationsByRobboUnitId(robboUnitId string) (err error)
	CreateCourseRelationGroup(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateCourseRelationUnit(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateCourseRelationStudent(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateCourseRelationTeacher(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateCourseRelationUnitAdmin(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error)
	CreateCourse(course *models.CourseCore) (id string, err error)
	DeleteCourse(courseId string) (err error)
	UpdateCourse(course *models.CourseCore) (err error)
}
