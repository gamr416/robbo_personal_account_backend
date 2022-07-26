package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx/api"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx/cohorts"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx/users"
	"go.uber.org/fx"
	"net/http"
	"net/url"
)

//go:generate mockgen -source=edx.go -destination=mocks/mock.go

type EdxApiUseCaseImpl struct {
}
type EdxApiUseCaseModule struct {
	fx.Out
	ApiCourse edx.CourseUseCase
	ApiUser   edx.UserUseCase
	ApiCohort edx.CohortUseCase
	ApiAuth   edx.AuthUseCase
}

func SetupEdxApiUseCase() EdxApiUseCaseModule {
	return EdxApiUseCaseModule{
		ApiCourse: &courses.EdxApiCourseImpl{},
		ApiUser:   &users.EdxApiUserImpl{},
		ApiCohort: &cohorts.EdxApiCohortImpl{},
		ApiAuth:   &api.EdxApiAuthImpl{},
	}
}

type myjar struct {
	jar map[string][]*http.Cookie
}

func (p *myjar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	p.jar[u.Host] = cookies
}

func (p *myjar) Cookies(u *url.URL) []*http.Cookie {
	return p.jar[u.Host]
}

func handleСookies(n []*http.Cookie) (csrfToken string, found bool) {
	for _, cookie := range n {
		if cookie.Name == "csrftoken" {
			return cookie.Value, true
		}
	}
	return "", false
}
