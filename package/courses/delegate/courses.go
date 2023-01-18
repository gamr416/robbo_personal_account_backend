package delegate

import (
	"encoding/json"
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"log"
	"strconv"
)

type CourseDelegateImpl struct {
	CoursesUseCase courses.UseCase
	EdxUseCase     edx.UseCase
}

type CourseDelegateModule struct {
	fx.Out
	courses.Delegate
}

func SetupCourseDelegate(coursesUsecase courses.UseCase, edxUsecase edx.UseCase) CourseDelegateModule {
	return CourseDelegateModule{
		Delegate: &CourseDelegateImpl{
			coursesUsecase,
			edxUsecase,
		},
	}
}

func (p *CourseDelegateImpl) CreateAccessCourseRelationRobboGroup(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error) {
	courseRelationCore := courseRelation.ToCore()
	newCourseRelationCore, err := p.CoursesUseCase.CreateAccessCourseRelationRobboGroup(courseRelationCore)
	if err != nil {
		return
	}
	newCourseRelation.FromCore(newCourseRelationCore)
	return
}

func (p *CourseDelegateImpl) CreateAccessCourseRelationRobboUnit(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error) {
	courseRelationCore := courseRelation.ToCore()
	newCourseRelationCore, err := p.CoursesUseCase.CreateAccessCourseRelationRobboUnit(courseRelationCore)
	if err != nil {
		return
	}
	newCourseRelation.FromCore(newCourseRelationCore)
	return
}

func (p *CourseDelegateImpl) CreateAccessCourseRelationStudent(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error) {
	courseRelationCore := courseRelation.ToCore()
	newCourseRelationCore, err := p.CoursesUseCase.CreateAccessCourseRelationStudent(courseRelationCore)
	if err != nil {
		return
	}
	newCourseRelation.FromCore(newCourseRelationCore)
	return
}

func (p *CourseDelegateImpl) CreateAccessCourseRelationTeacher(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error) {
	courseRelationCore := courseRelation.ToCore()
	newCourseRelationCore, err := p.CoursesUseCase.CreateAccessCourseRelationTeacher(courseRelationCore)
	if err != nil {
		return
	}
	newCourseRelation.FromCore(newCourseRelationCore)
	return
}

func (p *CourseDelegateImpl) CreateAccessCourseRelationUnitAdmin(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error) {
	courseRelationCore := courseRelation.ToCore()
	newCourseRelationCore, err := p.CoursesUseCase.CreateAccessCourseRelationUnitAdmin(courseRelationCore)
	if err != nil {
		return
	}
	newCourseRelation.FromCore(newCourseRelationCore)
	return
}

func (p *CourseDelegateImpl) DeleteAccessCourseRelationById(courseRelationId string) (id string, err error) {
	return p.CoursesUseCase.DeleteAccessCourseRelationById(courseRelationId)
}

func (p *CourseDelegateImpl) GetAccessCourseRelationsByCourseId(courseId string) (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetAccessCourseRelationsByCourseId(courseId)
	if err != nil {
		return
	}
	for _, courseRelationCore := range courseRelationsCore {
		var courseRelationTemp models.CourseRelationHTTP
		courseRelationTemp.FromCore(courseRelationCore)
		courseRelations = append(courseRelations, &courseRelationTemp)
	}
	return
}

func (p *CourseDelegateImpl) GetAccessCourseRelationsByRobboUnitId(robboUnitId string) (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetAccessCourseRelationsByRobboUnitId(robboUnitId)
	if err != nil {
		return
	}
	for _, courseRelationCore := range courseRelationsCore {
		var courseRelationTemp models.CourseRelationHTTP
		courseRelationTemp.FromCore(courseRelationCore)
		courseRelations = append(courseRelations, &courseRelationTemp)
	}
	return
}

func (p *CourseDelegateImpl) GetAccessCourseRelationsByRobboGroupId(robboGroupId string) (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetAccessCourseRelationsByRobboGroupId(robboGroupId)
	if err != nil {
		return
	}
	for _, courseRelationCore := range courseRelationsCore {
		var courseRelationTemp models.CourseRelationHTTP
		courseRelationTemp.FromCore(courseRelationCore)
		courseRelations = append(courseRelations, &courseRelationTemp)
	}
	return
}

func (p *CourseDelegateImpl) GetAccessCourseRelationsByStudentId(studentId string) (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetAccessCourseRelationsByStudentId(studentId)
	if err != nil {
		return
	}
	for _, courseRelationCore := range courseRelationsCore {
		var courseRelationTemp models.CourseRelationHTTP
		courseRelationTemp.FromCore(courseRelationCore)
		courseRelations = append(courseRelations, &courseRelationTemp)
	}
	return
}

func (p *CourseDelegateImpl) GetAccessCourseRelationsByTeacherId(teacherId string) (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetAccessCourseRelationsByTeacherId(teacherId)
	if err != nil {
		return
	}
	for _, courseRelationCore := range courseRelationsCore {
		var courseRelationTemp models.CourseRelationHTTP
		courseRelationTemp.FromCore(courseRelationCore)
		courseRelations = append(courseRelations, &courseRelationTemp)
	}
	return
}

func (p *CourseDelegateImpl) GetAccessCourseRelationsByUnitAdminId(unitAdminId string) (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetAccessCourseRelationsByUnitAdminId(unitAdminId)
	if err != nil {
		return
	}
	for _, courseRelationCore := range courseRelationsCore {
		var courseRelationTemp models.CourseRelationHTTP
		courseRelationTemp.FromCore(courseRelationCore)
		courseRelations = append(courseRelations, &courseRelationTemp)
	}
	return
}

func (p *CourseDelegateImpl) GetAccessCourseRelationsRobboUnits() (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetAccessCourseRelationsRobboUnits()
	if err != nil {
		return
	}
	for _, courseRelationCore := range courseRelationsCore {
		var courseRelationTemp models.CourseRelationHTTP
		courseRelationTemp.FromCore(courseRelationCore)
		courseRelations = append(courseRelations, &courseRelationTemp)
	}
	return
}

func (p *CourseDelegateImpl) GetAccessCourseRelationsRobboGroups() (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetAccessCourseRelationsRobboGroups()
	if err != nil {
		return
	}
	for _, courseRelationCore := range courseRelationsCore {
		var courseRelationTemp models.CourseRelationHTTP
		courseRelationTemp.FromCore(courseRelationCore)
		courseRelations = append(courseRelations, &courseRelationTemp)
	}
	return
}

func (p *CourseDelegateImpl) GetAccessCourseRelationsStudents() (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetAccessCourseRelationsStudents()
	if err != nil {
		return
	}
	for _, courseRelationCore := range courseRelationsCore {
		var courseRelationTemp models.CourseRelationHTTP
		courseRelationTemp.FromCore(courseRelationCore)
		courseRelations = append(courseRelations, &courseRelationTemp)
	}
	return
}

func (p *CourseDelegateImpl) GetAccessCourseRelationsTeachers() (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetAccessCourseRelationsTeachers()
	if err != nil {
		return
	}
	for _, courseRelationCore := range courseRelationsCore {
		var courseRelationTemp models.CourseRelationHTTP
		courseRelationTemp.FromCore(courseRelationCore)
		courseRelations = append(courseRelations, &courseRelationTemp)
	}
	return
}

func (p *CourseDelegateImpl) GetAccessCourseRelationsUnitAdmins() (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetAccessCourseRelationsUnitAdmins()
	if err != nil {
		return
	}
	for _, courseRelationCore := range courseRelationsCore {
		var courseRelationTemp models.CourseRelationHTTP
		courseRelationTemp.FromCore(courseRelationCore)
		courseRelations = append(courseRelations, &courseRelationTemp)
	}
	return
}

func (p *CourseDelegateImpl) CreateCourse(course *models.CourseHTTP, courseId string) (id string, err error) {
	body, err := p.EdxUseCase.GetCourseContent(courseId)
	if err != nil {
		return "", courses.ErrBadRequest
	}
	err = json.Unmarshal(body, course)
	if err != nil {
		return "", courses.ErrInternalServerLevel
	}
	courseCore := course.ToCore()
	return p.CoursesUseCase.CreateCourse(courseCore)
}

func (p *CourseDelegateImpl) DeleteCourse(courseId string) (err error) {
	return p.CoursesUseCase.DeleteCourse(courseId)
}

func (p *CourseDelegateImpl) UpdateCourse(course *models.CourseHTTP) (err error) {
	courseCore := course.ToCore()
	return p.CoursesUseCase.UpdateCourse(courseCore)
}

func (p *CourseDelegateImpl) GetCourseContent(courseId string) (courseHTTP *models.CourseHTTP, err error) {
	body, err := p.EdxUseCase.GetCourseContent(courseId)
	if err != nil {
		return nil, courses.ErrBadRequest
	}
	err = json.Unmarshal(body, &courseHTTP)
	if err != nil {
		return nil, courses.ErrInternalServerLevel
	}
	return courseHTTP, nil
}

func (p *CourseDelegateImpl) GetCoursesByUser(userId string, role models.Role) (
	coursesListHTTP *models.CoursesListHTTP,
	err error,
) {
	var courseAccessRelations []*models.CourseRelationCore
	var errGetRelations error
	switch role {
	case models.Student:
		courseAccessRelations, errGetRelations = p.CoursesUseCase.GetAccessCourseRelationsByStudentId(userId)
	case models.Teacher:
		courseAccessRelations, err = p.CoursesUseCase.GetAccessCourseRelationsByTeacherId(userId)
	case models.UnitAdmin:
		courseAccessRelations, err = p.CoursesUseCase.GetAccessCourseRelationsByUnitAdminId(userId)
	case models.SuperAdmin:
		body, err := p.EdxUseCase.GetCoursesByUser()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		err = json.Unmarshal(body, &coursesListHTTP)
		if err != nil {
			return nil, courses.ErrInternalServerLevel
		}
		return coursesListHTTP, nil
	}
	if errGetRelations != nil {
		return nil, errGetRelations
	}
	for _, courseAccessRelation := range courseAccessRelations {
		var courseHTTP *models.CourseHTTP
		body, err := p.EdxUseCase.GetCourseContent(courseAccessRelation.CourseId)
		if err != nil {
			return nil, courses.ErrBadRequest
		}
		err = json.Unmarshal(body, &courseHTTP)
		fmt.Println(courseHTTP)
		if err != nil {
			return nil, courses.ErrInternalServerLevel
		}
		coursesListHTTP = &models.CoursesListHTTP{
			Results:    []*models.CourseHTTP{},
			Pagination: &models.Pagination{},
		}
		coursesListHTTP.Results = append(coursesListHTTP.Results, courseHTTP)
	}
	return coursesListHTTP, nil

	//body, err := p.EdxUseCase.GetCoursesByUser()
	//if err != nil {
	//	return nil, courses.ErrBadRequest
	//}
	//err = json.Unmarshal(body, &coursesListHTTP)
	//if err != nil {
	//	return nil, courses.ErrInternalServerLevel
	//}
	//return coursesListHTTP, nil
}

func (p *CourseDelegateImpl) GetEnrollments(username string) (enrollmentsListHTTP *models.EnrollmentsListHTTP, err error) {
	body, err := p.EdxUseCase.GetEnrollments(username)
	if err != nil {
		return nil, courses.ErrBadRequest
	}
	err = json.Unmarshal(body, &enrollmentsListHTTP)
	if err != nil {
		return nil, courses.ErrInternalServerLevel
	}
	return enrollmentsListHTTP, nil
}

func (p *CourseDelegateImpl) GetAllPublicCourses(pageNumber string) (coursesListHTTP *models.CoursesListHTTP, err error) {
	pN, err := strconv.Atoi(pageNumber)
	if err != nil {
		return nil, courses.ErrBadRequest
	}
	body, err := p.EdxUseCase.GetAllPublicCourses(pN)
	if err != nil {
		return nil, courses.ErrBadRequest
	}
	err = json.Unmarshal(body, &coursesListHTTP)
	if err != nil {
		return nil, courses.ErrInternalServerLevel
	}
	return coursesListHTTP, nil
}

func (p *CourseDelegateImpl) PostEnrollment(postEnrollmentHTTP *models.PostEnrollmentHTTP) (err error) {
	_, err = p.EdxUseCase.PostEnrollment(postEnrollmentHTTP.Message)
	if err != nil {
		log.Println(err)
		return courses.ErrBadRequest
	}
	return
}

func (p *CourseDelegateImpl) PostUnenroll(postUnenrollHTTP *models.PostEnrollmentHTTP) (err error) {
	_, err = p.EdxUseCase.PostEnrollment(postUnenrollHTTP.Message)
	if err != nil {
		log.Println(err)
		return courses.ErrBadRequest
	}
	return
}

func (p *CourseDelegateImpl) Login(email, password string) (err error) {
	_, err = p.EdxUseCase.Login(email, password)
	if err != nil {
		log.Println(err)
		return courses.ErrBadRequest
	}
	return
}

func (p *CourseDelegateImpl) Registration(userForm *edx.RegistrationForm) (err error) {
	_, err = p.EdxUseCase.PostRegistration(*userForm)
	if err != nil {
		log.Println(err)
		return courses.ErrBadRequest
	}
	return
}
