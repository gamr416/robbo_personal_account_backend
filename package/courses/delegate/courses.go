package delegate

import (
	"encoding/json"
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"log"
)

type CourseDelegateImpl struct {
	courses.UseCase
	edxApi.EdxApiCourse
	edxApi.EdxApiUser
}

type CourseDelegateModule struct {
	fx.Out
	courses.Delegate
}

func SetupCourseDelegate(usecase courses.UseCase, edxCourse edxApi.EdxApiCourse, edxUser edxApi.EdxApiUser) CourseDelegateModule {
	return CourseDelegateModule{
		Delegate: &CourseDelegateImpl{
			usecase,
			edxCourse,
			edxUser,
		},
	}
}

func (p *CourseDelegateImpl) CreateCourse(course *models.CourseHTTP, courseId string) (id string, err error) {
	body, err := p.EdxApiCourse.GetCourseContent(courseId)
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
	body, err := p.EdxApiCourse.GetCourseContent(courseId)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func (p *CourseDelegateImpl) GetCoursesByUser() (respBody []byte, err error) {
	body, err := p.EdxApiCourse.GetCoursesByUser()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (p *CourseDelegateImpl) GetEnrollments(username string) (respBody []byte, err error) {
	body, err := p.EdxApiCourse.GetEnrollments(username)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (p *CourseDelegateImpl) GetAllPublicCourses(pageNumber int) (respBody []byte, err error) {
	body, err := p.EdxApiCourse.GetAllPublicCourses(pageNumber)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (p *CourseDelegateImpl) PostEnrollment(postEnrollmentHTTP *models.PostEnrollmentHTTP) (err error) {
	_, err = p.EdxApiCourse.PostEnrollment(postEnrollmentHTTP.Message)
	if err != nil {
		log.Println(err)
		return err
	}
	return
}

func (p *CourseDelegateImpl) PostUnenroll(postUnenrollHTTP *models.PostEnrollmentHTTP) (err error) {
	_, err = p.EdxApiCourse.PostEnrollment(postUnenrollHTTP.Message)
	if err != nil {
		log.Println(err)
		return err
	}
	return
}

func (p *CourseDelegateImpl) Login(email, password string) (err error) {
	_, err = p.EdxApiUser.Login(email, password)
	if err != nil {
		log.Println(err)
		return err
	}
	return
}

func (p *CourseDelegateImpl) Registration(userForm *edxApi.RegistrationForm) (err error) {
	_, err = p.EdxApiUser.PostRegistration(*userForm)
	if err != nil {
		log.Println(err)
		return err
	}
	return
}
