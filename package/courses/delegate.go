package courses

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	CreateCourse(course *models.CourseHTTP, courseId string) (id string, err error)
	DeleteCourse(courseId string) (err error)
	UpdateCourse(course *models.CourseHTTP) (err error)
	GetCourseContent(courseId string) (courseHTTP *models.CourseHTTP, err error)
	GetCoursesByUser() (coursesListHTTP *models.CoursesListHTTP, err error)
	GetAllPublicCourses(pageNumber string) (coursesListHTTP *models.CoursesListHTTP, err error)
	GetEnrollments(username string) (enrollmentListHTTP *models.EnrollmentsListHTTP, err error)
	PostUnenroll(postUnenrollHTTP *models.PostEnrollmentHTTP) (err error)
	Login(email string, password string) (err error)
	Registration(userForm *edx.RegistrationForm) (err error)
}
