package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
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

func (p *CourseUseCaseImpl) CreateCourse(course *models.CourseCore, courseId string) (id string, err error) {
	return p.Gateway.CreateCourse(course)
}

func (p *CourseUseCaseImpl) UpdateCourse(course *models.CourseCore) (err error) {
	return p.Gateway.UpdateCourse(course)
}

func (p *CourseUseCaseImpl) DeleteCourse(course *models.CourseCore) (err error) {
	return p.Gateway.DeleteCourse(course)
}
