package delegate

import (
	"encoding/json"
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

func (p *CourseDelegateImpl) CreateCourseRelationGroup(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error) {
	courseRelationCore := courseRelation.ToCore()
	newCourseRelationCore, err := p.CoursesUseCase.CreateCourseRelationGroup(courseRelationCore)
	if err != nil {
		return
	}
	newCourseRelation.FromCore(newCourseRelationCore)
	return
}

func (p *CourseDelegateImpl) CreateCourseRelationUnit(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error) {
	courseRelationCore := courseRelation.ToCore()
	newCourseRelationCore, err := p.CoursesUseCase.CreateCourseRelationUnit(courseRelationCore)
	if err != nil {
		return
	}
	newCourseRelation.FromCore(newCourseRelationCore)
	return
}

func (p *CourseDelegateImpl) CreateCourseRelationStudent(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error) {
	courseRelationCore := courseRelation.ToCore()
	newCourseRelationCore, err := p.CoursesUseCase.CreateCourseRelationStudent(courseRelationCore)
	if err != nil {
		return
	}
	newCourseRelation.FromCore(newCourseRelationCore)
	return
}

func (p *CourseDelegateImpl) CreateCourseRelationTeacher(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error) {
	courseRelationCore := courseRelation.ToCore()
	newCourseRelationCore, err := p.CoursesUseCase.CreateCourseRelationTeacher(courseRelationCore)
	if err != nil {
		return
	}
	newCourseRelation.FromCore(newCourseRelationCore)
	return
}

func (p *CourseDelegateImpl) CreateCourseRelationUnitAdmin(courseRelation *models.CourseRelationHTTP) (newCourseRelation models.CourseRelationHTTP, err error) {
	courseRelationCore := courseRelation.ToCore()
	newCourseRelationCore, err := p.CoursesUseCase.CreateCourseRelationUnitAdmin(courseRelationCore)
	if err != nil {
		return
	}
	newCourseRelation.FromCore(newCourseRelationCore)
	return
}

func (p *CourseDelegateImpl) DeleteCourseRelationById(courseRelationId string) (id string, err error) {
	return p.CoursesUseCase.DeleteCourseRelationById(courseRelationId)
}

func (p *CourseDelegateImpl) GetCourseRelationsByCourseId(courseId string) (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetCourseRelationsByCourseId(courseId)
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

func (p *CourseDelegateImpl) GetCourseRelationsByRobboUnitId(robboUnitId string) (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetCourseRelationsByRobboUnitId(robboUnitId)
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

func (p *CourseDelegateImpl) GetCourseRelationsByRobboGroupId(robboGroupId string) (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetCourseRelationsByRobboGroupId(robboGroupId)
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

func (p *CourseDelegateImpl) GetCourseRelationsByStudentId(studentId string) (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetCourseRelationsByStudentId(studentId)
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

func (p *CourseDelegateImpl) GetCourseRelationsByTeacherId(teacherId string) (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetCourseRelationsByTeacherId(teacherId)
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

func (p *CourseDelegateImpl) GetCourseRelationsByUnitAdminId(unitAdminId string) (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetCourseRelationsByUnitAdminId(unitAdminId)
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

func (p *CourseDelegateImpl) GetCourseRelationsUnits() (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetCourseRelationsUnits()
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

func (p *CourseDelegateImpl) GetCourseRelationsGroups() (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetCourseRelationsGroups()
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

func (p *CourseDelegateImpl) GetCourseRelationsStudents() (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetCourseRelationsStudents()
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

func (p *CourseDelegateImpl) GetCourseRelationsTeachers() (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetCourseRelationsTeachers()
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

func (p *CourseDelegateImpl) GetCourseRelationsUnitAdmins() (courseRelations []*models.CourseRelationHTTP, err error) {
	courseRelationsCore, err := p.CoursesUseCase.GetCourseRelationsUnitAdmins()
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

func (p *CourseDelegateImpl) GetCoursesByUser() (coursesListHTTP *models.CoursesListHTTP, err error) {
	body, err := p.EdxUseCase.GetCoursesByUser()
	if err != nil {
		return nil, courses.ErrBadRequest
	}
	err = json.Unmarshal(body, &coursesListHTTP)
	if err != nil {
		return nil, courses.ErrInternalServerLevel
	}
	return coursesListHTTP, nil
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
