package courses

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateCourse(course *models.CourseCore, courseId string) (id string, err error)
	DeleteCourse(course *models.CourseCore) (err error)
	UpdateCourse(course *models.CourseCore) (err error)
	//GetCoursesByUser() (respBody string, err error)
	//GetAllPublicCourses(body string) (respBody string, err error)
	//GetCourseContent(course *models.CourseCore) (respBody string, err error)
}
