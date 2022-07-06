package courses

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateCourse(course *models.CourseHTTP, courseId string) (id string, statusCode int, err error)
	DeleteCourse(course *models.CourseHTTP) (err error)
	UpdateCourse(course *models.CourseHTTP) (err error)
	GetCoursesForUser() (respBody string, statusCode int, err error)
	GetAllPublicCourses(pageNumber int) (respBody string, statusCode int, err error)
	GetEnrollments(username string) (respBody string, statusCode int, err error)
	GetUser() (respBody string, statusCode int, err error)
	GetCourseContent(courseId string) (respBody string, statusCode int, err error)
	PostEnrollment(message map[string]interface{}) (respBody string, statusCode int, err error)
}
