package delegate

import (
	"encoding/json"
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
)

type CourseDelegateImpl struct {
	courses.UseCase
	edx.CourseUseCase
}

type CourseDelegateModule struct {
	fx.Out
	courses.Delegate
}

func SetupCourseDelegate(usecase courses.UseCase, edx edx.CourseUseCase) CourseDelegateModule {
	return CourseDelegateModule{
		Delegate: &CourseDelegateImpl{
			usecase,
			edx,
		},
	}
}

func (p *CourseDelegateImpl) CreateCourse(course *models.CourseHTTP, courseId string) (id string, err error) {
	body, err := p.CourseUseCase.GetCourseContent(courseId)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, course)
	fmt.Println(course)
	if err != nil {
		return "", err
	}
	courseCore := course.ToCore()
	return p.UseCase.CreateCourse(courseCore)
}

func (p *CourseDelegateImpl) DeleteCourse(courseId string) (err error) {
	return p.UseCase.DeleteCourse(courseId)
}

func (p *CourseDelegateImpl) UpdateCourse(course *models.CourseHTTP) (err error) {
	courseCore := course.ToCore()
	return p.UseCase.UpdateCourse(courseCore)
}

func (p *CourseDelegateImpl) GetCourseContent(courseId string) (respBody []byte, err error) {
	body, err := p.CourseUseCase.GetCourseContent(courseId)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func (p *CourseDelegateImpl) GetCoursesByUser() (respBody []byte, err error) {
	body, err := p.CourseUseCase.GetCoursesByUser()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (p *CourseDelegateImpl) GetEnrollments(username string) (respBody []byte, err error) {
	body, err := p.CourseUseCase.GetEnrollments(username)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (p *CourseDelegateImpl) GetAllPublicCourses(pageNumber int) (respBody []byte, err error) {
	body, err := p.CourseUseCase.GetAllPublicCourses(pageNumber)
	if err != nil {
		return nil, err
	}
	return body, nil
}
