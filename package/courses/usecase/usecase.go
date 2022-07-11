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

func (p *CourseUseCaseImpl) CreateCourse(course *models.CourseCore) (id string, err error) {
	CourseId, err := p.Gateway.CreateCourse(course)
	if err != nil {
		return "", err
	}

	mediaCore := &models.CourseApiMediaCollectionCore{
		CourseID: CourseId,
	}

	MediaId, err := p.Gateway.CreateCourseApiMediaCollection(mediaCore)
	if err != nil {
		return "", err
	}

	bannerImage := &models.AbsoluteMediaCore{
		Uri:                        course.Media.BannerImage.Uri,
		UriAbsolute:                course.Media.BannerImage.UriAbsolute,
		CourseApiMediaCollectionID: MediaId,
	}

	_, err = p.Gateway.CreateAbsoluteMedia(bannerImage)
	if err != nil {
		return "", err
	}

	courseImage := &models.MediaCore{
		Uri:                        course.Media.CourseImage.Uri,
		CourseApiMediaCollectionID: MediaId,
	}

	_, err = p.Gateway.CreateMedia(courseImage)
	if err != nil {
		return "", err
	}

	courseVideo := &models.MediaCore{
		Uri:                        course.Media.CourseVideo.Uri,
		CourseApiMediaCollectionID: MediaId,
	}

	_, err = p.Gateway.CreateMedia(courseVideo)
	if err != nil {
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
		return "", err
	}

	return CourseId, nil
}

func (p *CourseUseCaseImpl) UpdateCourse(course *models.CourseCore) (err error) {

	return p.Gateway.UpdateCourse(course)
}

func (p *CourseUseCaseImpl) DeleteCourse(courseId string) (err error) {
	return p.Gateway.DeleteCourse(courseId)
}
