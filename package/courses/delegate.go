package courses

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateCourse(course *models.CourseHTTP, courseId string) (id string, err error)
	DeleteCourse(course *models.CourseHTTP) (err error)
	UpdateCourse(course *models.CourseHTTP) (err error)
	GetCourseContent(courseId string) (respBody string, err error)
	/*GetCoursesByUser() (respBody string,  err error)
	GetAllPublicCourses(pageNumber int) (respBody string,  err error)
	*/
}
