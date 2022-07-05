package edx_api

type EdxApi interface {
	GetCourses() string
	GetAllCourses(pageNumber int) (string, int)
	GetWithAuth(funcName, url, token string) (string, int)
	GetEnrollment(username, token string) (string, int)
	GetUser(token string) (string, int)
	GetCourse(courseId, token string) (string, int)
	PostEnrollment(url, token string, message map[string]interface{}) (string, int)
	RefreshToken(token *string)
	PostRegistration(urlAddr string) (string, int)
	TestApi()
}
