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
<<<<<<< HEAD
	edx.CourseUseCase
=======
	edxApi.EdxApiCourse
	edxApi.EdxApiUser
>>>>>>> b51413c19b53a2b40776b2746be8d694d6f8e40e
}

type CourseDelegateModule struct {
	fx.Out
	courses.Delegate
}

<<<<<<< HEAD
func SetupCourseDelegate(usecase courses.UseCase, edx edx.CourseUseCase) CourseDelegateModule {
=======
func SetupCourseDelegate(usecase courses.UseCase, edxCourse edxApi.EdxApiCourse, edxUser edxApi.EdxApiUser) CourseDelegateModule {
>>>>>>> b51413c19b53a2b40776b2746be8d694d6f8e40e
	return CourseDelegateModule{
		Delegate: &CourseDelegateImpl{
			usecase,
			edxCourse,
			edxUser,
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
<<<<<<< HEAD
=======

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
>>>>>>> b51413c19b53a2b40776b2746be8d694d6f8e40e
