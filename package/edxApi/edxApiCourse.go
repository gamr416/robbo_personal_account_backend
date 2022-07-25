package edxApi

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type EdxApiCourse interface {
	GetCoursesByUser() (respBody []byte, err error)
	GetAllPublicCourses(pageNumber int) (respBody []byte, err error)
	GetEnrollments(username string) (respBody []byte, err error)
	GetCourseContent(courseId string) (respBody []byte, err error)
	PostEnrollment(message map[string]interface{}) (respBody []byte, err error)
}
