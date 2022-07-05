package usecase

import (
	"encoding/json"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx_api"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"log"
	"net/http"
)

type CourseUseCaseImpl struct {
	courses.Gateway
}

type CourseUseCaseModule struct {
	fx.Out
	courses.UseCase
}

func SetupCourseUseCase(gateway courses.Gateway) CourseUseCaseModule {
	return CourseUseCaseModule{
		UseCase: &CourseUseCaseImpl{
			Gateway: gateway,
		},
	}
}

func (p *CourseUseCaseImpl) CreateCourse(course *models.CourseHTTP, courseId string) (id string, err error) {
	var token string
	edx_api.RefreshToken(&token)
	body, statusCode := edx_api.GetCourse(courseId, token)
	if statusCode != http.StatusOK {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(body), &course)
	if err != nil {
		log.Fatalln(err)
	}
	courseCore := course.ToCore()
	return p.Gateway.CreateCourse(courseCore, courseId)
}

func (p *CourseUseCaseImpl) GetCoursesByUser(username string) (body string, err error) {
	var token string
	edx_api.RefreshToken(&token)
	body, statusCode := edx_api.GetEnrollment(username, token)
	if statusCode == http.StatusOK {
		return body, nil
	}
	return "", err
}

func (p *CourseUseCaseImpl) GetCourseContent(courseId string) (body string, err error) {
	var token string
	edx_api.RefreshToken(&token)
	body, statusCode := edx_api.GetCourse(courseId, token)
	if statusCode == http.StatusOK {
		return body, nil
	}
	return "", err
}

func (p *CourseUseCaseImpl) UpdateCourse(course *models.CourseCore) (err error) {
	return p.Gateway.UpdateCourse(course)
}

func (p *CourseUseCaseImpl) DeleteCourse(course *models.CourseCore) (err error) {
	return p.Gateway.DeleteCourse(course)
}

func (p *CourseUseCaseImpl) GetAllPublicCourses(pageNumber int) (body string, err error) {
	body, statusCode := edx_api.GetAllCourses(pageNumber)
	if statusCode == http.StatusOK {
		return body, nil
	}
	return "", err

}

func (p *CourseUseCaseImpl) GetAllCourses() (body string, err error) {
	body, statusCode := edx_api.GetCourses()
	if statusCode == http.StatusOK {
		return body, nil
	}
	return "", err
}
