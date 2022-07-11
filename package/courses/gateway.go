package courses

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateAbsoluteMedia(absoluteMedia *models.AbsoluteMediaCore) (id string, err error)
	CreateMedia(media *models.MediaCore) (id string, err error)
	CreateImage(image *models.ImageCore) (id string, err error)
	CreateCourseApiMediaCollection(courseApiMediaCollection *models.CourseApiMediaCollectionCore) (id string, err error)
	CreateCourse(course *models.CourseCore) (id string, err error)
	DeleteCourse(courseId string) (err error)
	UpdateCourse(course *models.CourseCore) (err error)
}
