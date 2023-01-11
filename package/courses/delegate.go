package courses

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	GetCourseRelationsByCourseId(courseId string) (courseRelations []*models.CourseRelationHTTP, err error)
	GetCourseRelationsByRobboUnitId(robboUnitId string) (courseRelations []*models.CourseRelationHTTP, err error)
	GetCourseRelationsByRobboGroupId(robboGroupId string) (courseRelations []*models.CourseRelationHTTP, err error)
	GetCourseRelationsByStudentId(studentId string) (courseRelations []*models.CourseRelationHTTP, err error)
	GetCourseRelationsByTeacherId(teacherId string) (courseRelations []*models.CourseRelationHTTP, err error)
	GetCourseRelationsByUnitAdminId(unitAdminId string) (courseRelations []*models.CourseRelationHTTP, err error)
	GetCourseRelationsUnits() (courseRelations []*models.CourseRelationHTTP, err error)
	GetCourseRelationsGroups() (courseRelations []*models.CourseRelationHTTP, err error)
	GetCourseRelationsStudents() (courseRelations []*models.CourseRelationHTTP, err error)
	GetCourseRelationsTeachers() (courseRelations []*models.CourseRelationHTTP, err error)
	GetCourseRelationsUnitAdmins() (courseRelations []*models.CourseRelationHTTP, err error)
	CreateCourseRelationStudent(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error)
	CreateCourseRelationTeacher(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error)
	CreateCourseRelationUnitAdmin(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error)
	CreateCourseRelationUnit(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error)
	CreateCourseRelationGroup(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error)
	CreateCourse(course *models.CourseHTTP, courseId string) (id string, err error)
	DeleteCourse(courseId string) (err error)
	DeleteCourseRelationById(courseRelationId string) (id string, err error)
	UpdateCourse(course *models.CourseHTTP) (err error)
	GetCourseContent(courseId string) (courseHTTP *models.CourseHTTP, err error)
	GetCoursesByUser() (coursesListHTTP *models.CoursesListHTTP, err error)
	GetAllPublicCourses(pageNumber string) (coursesListHTTP *models.CoursesListHTTP, err error)
	GetEnrollments(username string) (enrollmentListHTTP *models.EnrollmentsListHTTP, err error)
	PostUnenroll(postUnenrollHTTP *models.PostEnrollmentHTTP) (err error)
	Login(email string, password string) (err error)
	Registration(userForm *edx.RegistrationForm) (err error)
}
