package courses

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateCourse(course *models.CourseHTTP, courseId string) (id string, err error)
	DeleteCourse(courseId string) (err error)
	UpdateCourse(course *models.CourseHTTP) (err error)
	GetCourseContent(courseId string) (respBody []byte, err error)
	GetCoursesByUser() (respBody []byte, err error)
	GetAllPublicCourses(pageNumber int) (respBody []byte, err error)
	GetEnrollments(username string) (respBody []byte, err error)
}
