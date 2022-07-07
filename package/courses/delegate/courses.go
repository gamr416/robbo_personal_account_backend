package delegate

import (
	"encoding/json"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx_api"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
)

type CourseDelegateImpl struct {
	courses.UseCase
	edx_api.EdxApiUseCase
}

type CourseDelegateModule struct {
	fx.Out
	courses.Delegate
}

func SetupCourseDelegate(usecase courses.UseCase, edx edx_api.EdxApiUseCase) CourseDelegateModule {
	return CourseDelegateModule{
		Delegate: &CourseDelegateImpl{
			usecase,
			edx,
		},
	}
}

func (p *CourseDelegateImpl) CreateCourse(course *models.CourseHTTP, courseId string) (id string, err error) {
	body, err := p.EdxApiUseCase.GetCourseContent(courseId)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(body), course)
	if err != nil {
		return "", err
	}
	courseCore := course.ToCore()
	return p.UseCase.CreateCourse(courseCore, courseId)
}

func (p *CourseDelegateImpl) DeleteCourse(course *models.CourseHTTP) (err error) {
	courseCore := course.ToCore()
	return p.UseCase.DeleteCourse(courseCore)
}

func (p *CourseDelegateImpl) UpdateCourse(course *models.CourseHTTP) (err error) {
	//TODO implement me
	panic("implement me")
}
func (p *CourseDelegateImpl) GetCourseContent(courseId string) (respBody string, err error) {
	body, err := p.EdxApiUseCase.GetCourseContent(courseId)
	if err != nil {
		return "", err
	}
	return body, nil
}
func (p *CourseDelegateImpl) GetCoursesByUser() (respBody string, err error) {
	body, err := p.EdxApiUseCase.GetCoursesByUser()
	if err != nil {
		return "", err
	}
	return body, nil
}

func (p *CourseDelegateImpl) GetEnrollments(username string) (respBody string, err error) {
	body, err := p.EdxApiUseCase.GetEnrollments(username)
	if err != nil {
		return "", err
	}
	return body, nil
}

func (p *CourseDelegateImpl) GetAllPublicCourses(pageNumber int) (respBody string, err error) {
	body, err := p.EdxApiUseCase.GetAllPublicCourses(pageNumber)
	if err != nil {
		return "", err
	}
	return body, nil
}
