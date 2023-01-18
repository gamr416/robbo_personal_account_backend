package courses

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	GetAccessCourseRelationsByCourseId(courseId string) (courseRelations []*models.CourseRelationHTTP, err error)
	GetAccessCourseRelationsByRobboUnitId(robboUnitId string) (courseRelations []*models.CourseRelationHTTP, err error)
	GetAccessCourseRelationsByRobboGroupId(robboGroupId string) (courseRelations []*models.CourseRelationHTTP, err error)
	GetAccessCourseRelationsByStudentId(studentId string) (courseRelations []*models.CourseRelationHTTP, err error)
	GetAccessCourseRelationsByTeacherId(teacherId string) (courseRelations []*models.CourseRelationHTTP, err error)
	GetAccessCourseRelationsByUnitAdminId(unitAdminId string) (courseRelations []*models.CourseRelationHTTP, err error)
	GetAccessCourseRelationsRobboUnits() (courseRelations []*models.CourseRelationHTTP, err error)
	GetAccessCourseRelationsRobboGroups() (courseRelations []*models.CourseRelationHTTP, err error)
	GetAccessCourseRelationsStudents() (courseRelations []*models.CourseRelationHTTP, err error)
	GetAccessCourseRelationsTeachers() (courseRelations []*models.CourseRelationHTTP, err error)
	GetAccessCourseRelationsUnitAdmins() (courseRelations []*models.CourseRelationHTTP, err error)
	CreateAccessCourseRelationStudent(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error)
	CreateAccessCourseRelationTeacher(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error)
	CreateAccessCourseRelationUnitAdmin(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error)
	CreateAccessCourseRelationRobboUnit(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error)
	CreateAccessCourseRelationRobboGroup(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error)
	DeleteAccessCourseRelationById(courseRelationId string) (id string, err error)
	CreateCourse(course *models.CourseHTTP, courseId string) (id string, err error)
	DeleteCourse(courseId string) (err error)
	UpdateCourse(course *models.CourseHTTP) (err error)
	GetCourseContent(courseId string) (courseHTTP *models.CourseHTTP, err error)
	GetCoursesByUser(userId string, role models.Role) (coursesListHTTP *models.CoursesListHTTP, err error)
	GetAllPublicCourses(pageNumber string) (coursesListHTTP *models.CoursesListHTTP, err error)
	GetEnrollments(username string) (enrollmentListHTTP *models.EnrollmentsListHTTP, err error)
	PostUnenroll(postUnenrollHTTP *models.PostEnrollmentHTTP) (err error)
	Login(email string, password string) (err error)
	Registration(userForm *edx.RegistrationForm) (err error)
}
