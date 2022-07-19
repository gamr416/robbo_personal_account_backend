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
		err = courses.ErrCourseNotFound
		return
	}

	return nil
}

func (p *CourseUseCaseImpl) DeleteCourse(courseId string) (err error) {
	id, err := p.Gateway.DeleteCourse(courseId)
	if err != nil {
		log.Println("Error delete Course")
		err = courses.ErrCourseNotFound
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
