package edxApi

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type NewToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type RegistrationForm struct {
	Email            string
	Username         string
	Name             string
	Password         string
	Terms_of_service string
}

type EdxApiUseCase interface {
	GetCoursesByUser() (respBody []byte, err error)
	GetAllPublicCourses(pageNumber int) (respBody []byte, err error) //+
	GetEnrollments(username string) (respBody []byte, err error)     //+
	GetUser() (respBody []byte, err error)
	GetCourseContent(courseId string) (respBody []byte, err error)              //+
	PostEnrollment(message map[string]interface{}) (respBody []byte, err error) //+
	PostRegistration(postMessage RegistrationForm) (respBody []byte, err error)
	Login(email, password string) (respBody []byte, err error) //+
}
