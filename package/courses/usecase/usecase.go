package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"log"
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

func (p *CourseUseCaseImpl) GetAccessCourseRelationsStudents() (courseRelations []*models.CourseRelationCore, err error) {
	return p.Gateway.GetAccessCourseRelationsStudents()
}

func (p *CourseUseCaseImpl) GetAccessCourseRelationsTeachers() (courseRelations []*models.CourseRelationCore, err error) {
	return p.Gateway.GetAccessCourseRelationsTeachers()
}

func (p *CourseUseCaseImpl) GetAccessCourseRelationsUnitAdmins() (courseRelations []*models.CourseRelationCore, err error) {
	return p.Gateway.GetAccessCourseRelationsUnitAdmins()
}

func (p *CourseUseCaseImpl) GetAccessCourseRelationsRobboUnits() (courseRelations []*models.CourseRelationCore, err error) {
	return p.Gateway.GetAccessCourseRelationsRobboUnits()
}

func (p *CourseUseCaseImpl) GetAccessCourseRelationsRobboGroups() (courseRelations []*models.CourseRelationCore, err error) {
	return p.Gateway.GetAccessCourseRelationsRobboGroups()
}

func (p *CourseUseCaseImpl) GetAccessCourseRelationsByStudentId(studentId string) (courseRelations []*models.CourseRelationCore, err error) {
	return p.Gateway.GetAccessCourseRelationsByStudentId(studentId)
}

func (p *CourseUseCaseImpl) GetAccessCourseRelationsByTeacherId(teacherId string) (courseRelations []*models.CourseRelationCore, err error) {
	return p.Gateway.GetAccessCourseRelationsByTeacherId(teacherId)
}

func (p *CourseUseCaseImpl) GetAccessCourseRelationsByUnitAdminId(unitAdminId string) (courseRelations []*models.CourseRelationCore, err error) {
	return p.Gateway.GetAccessCourseRelationsByUnitAdminId(unitAdminId)
}

func (p *CourseUseCaseImpl) GetAccessCourseRelationsByRobboGroupId(robboGroupId string) (courseRelations []*models.CourseRelationCore, err error) {
	return p.Gateway.GetAccessCourseRelationsByRobboGroupId(robboGroupId)
}

func (p *CourseUseCaseImpl) GetAccessCourseRelationsByRobboUnitId(robboUnitId string) (courseRelations []*models.CourseRelationCore, err error) {
	return p.Gateway.GetAccessCourseRelationsByRobboUnitId(robboUnitId)
}

func (p *CourseUseCaseImpl) GetAccessCourseRelationsByCourseId(courseId string) (courseRelations []*models.CourseRelationCore, err error) {
	return p.Gateway.GetAccessCourseRelationsByCourseId(courseId)
}

func (p *CourseUseCaseImpl) CreateAccessCourseRelationRobboGroup(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error) {
	courseRelation.Parameter = "robbo_group"
	return p.Gateway.CreateAccessCourseRelation(courseRelation)
}

func (p *CourseUseCaseImpl) CreateAccessCourseRelationRobboUnit(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error) {
	courseRelation.Parameter = "robbo_unit"
	return p.Gateway.CreateAccessCourseRelation(courseRelation)
}

func (p *CourseUseCaseImpl) CreateAccessCourseRelationStudent(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error) {
	courseRelation.Parameter = "student"
	return p.Gateway.CreateAccessCourseRelation(courseRelation)
}

func (p *CourseUseCaseImpl) CreateAccessCourseRelationTeacher(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error) {
	courseRelation.Parameter = "teacher"
	return p.Gateway.CreateAccessCourseRelation(courseRelation)
}

func (p *CourseUseCaseImpl) CreateAccessCourseRelationUnitAdmin(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error) {
	courseRelation.Parameter = "unit_admin"
	return p.Gateway.CreateAccessCourseRelation(courseRelation)
}

func (p *CourseUseCaseImpl) DeleteAccessCourseRelationById(courseRelationId string) (id string, err error) {
	return p.Gateway.DeleteAccessCourseRelationById(courseRelationId)
}

func (p *CourseUseCaseImpl) DeleteAccessCourseRelationsByStudentId(studentId string) (err error) {
	return p.Gateway.DeleteAccessCourseRelationsByStudentId(studentId)
}

func (p *CourseUseCaseImpl) DeleteAccessCourseRelationsByTeacherId(teacherId string) (err error) {
	return p.Gateway.DeleteAccessCourseRelationsByTeacherId(teacherId)
}

func (p *CourseUseCaseImpl) DeleteAccessCourseRelationsByUnitAdminId(unitAdminId string) (err error) {
	return p.Gateway.DeleteAccessCourseRelationsByUnitAdminId(unitAdminId)
}

func (p *CourseUseCaseImpl) DeleteAccessCourseRelationsByRobboGroupId(robboGroupId string) (err error) {
	return p.Gateway.DeleteAccessCourseRelationsByRobboGroupId(robboGroupId)
}

func (p *CourseUseCaseImpl) DeleteAccessCourseRelationsByRobboUnitId(robboUnitId string) (err error) {
	return p.Gateway.DeleteAccessCourseRelationsByRobboUnitId(robboUnitId)
}

func (p *CourseUseCaseImpl) CreateCourse(course *models.CourseCore) (id string, err error) {
	CourseId, err := p.Gateway.CreateCourse(course)
	if err != nil {
		log.Println("Error create Course")
		return "", err
	}

	mediaCore := &models.CourseApiMediaCollectionCore{
		CourseID: CourseId,
	}
	MediaId, err := p.Gateway.CreateCourseApiMediaCollection(mediaCore)
	if err != nil {
		log.Println("Error create CourseApiMediaCollection")
		return "", err
	}

	bannerImage := &models.AbsoluteMediaCore{
		Uri:                        course.Media.BannerImage.Uri,
		UriAbsolute:                course.Media.BannerImage.UriAbsolute,
		CourseApiMediaCollectionID: MediaId,
	}
	_, err = p.Gateway.CreateAbsoluteMedia(bannerImage)
	if err != nil {
		log.Println("Error create AbsoluteMedia")
		return "", err
	}

	courseImage := &models.MediaCore{
		Uri:                        course.Media.CourseImage.Uri,
		CourseApiMediaCollectionID: MediaId,
	}
	_, err = p.Gateway.CreateMedia(courseImage)
	if err != nil {
		log.Println("Error create Media")
		return "", err
	}

	courseVideo := &models.MediaCore{
		Uri:                        course.Media.CourseVideo.Uri,
		CourseApiMediaCollectionID: MediaId,
	}
	_, err = p.Gateway.CreateMedia(courseVideo)
	if err != nil {
		log.Println("Error create Media")
		return "", err
	}

	image := &models.ImageCore{
		Raw:                        course.Media.Image.Raw,
		Small:                      course.Media.Image.Small,
		Large:                      course.Media.Image.Large,
		CourseApiMediaCollectionID: MediaId,
	}
	_, err = p.Gateway.CreateImage(image)
	if err != nil {
		log.Println("Error create Image")
		return "", err
	}

	return CourseId, nil
}

func (p *CourseUseCaseImpl) UpdateCourse(course *models.CourseCore) (err error) {
	err = p.Gateway.UpdateAbsoluteMedia(&course.Media.BannerImage)
	if err != nil {
		log.Println("Error update AbsoluteMedia")
		return
	}

	err = p.Gateway.UpdateMedia(&course.Media.CourseImage)
	if err != nil {
		log.Println("Error update Media")
		return
	}

	err = p.Gateway.UpdateMedia(&course.Media.CourseVideo)
	if err != nil {
		log.Println("Error update Media")
		return
	}

	err = p.Gateway.UpdateImage(&course.Media.Image)
	if err != nil {
		log.Println("Error update Image")
		return
	}

	err = p.Gateway.UpdateCourseApiMediaCollection(&course.Media)
	if err != nil {
		log.Println("Error update CourseApiMediaCollection")
		return
	}

	err = p.Gateway.UpdateCourse(course)
	if err != nil {
		log.Println("Error update Course")
		return
	}

	return nil
}

func (p *CourseUseCaseImpl) DeleteCourse(courseId string) (err error) {
	id, err := p.Gateway.DeleteCourse(courseId)
	if err != nil {
		log.Println("Error delete Course")
		return
	}

	courseApiMediaCollectionId, err := p.Gateway.DeleteCourseApiMediaCollection(id)
	if err != nil {
		log.Println("Error delete CourseApiMediaCollection")
		return
	}

	err = p.Gateway.DeleteAbsoluteMedia(courseApiMediaCollectionId)
	if err != nil {
		log.Println("Error delete AbsoluteMedia")
		return
	}

	err = p.Gateway.DeleteMedia(courseApiMediaCollectionId)
	if err != nil {
		log.Println("Error delete Media")
		return
	}

	err = p.Gateway.DeleteImage(courseApiMediaCollectionId)
	if err != nil {
		log.Println("Error delete Image")
		return
	}

	return nil
}
