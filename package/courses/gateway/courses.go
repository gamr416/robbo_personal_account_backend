package gateway

import (
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type CoursesGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type CoursesGatewayModule struct {
	fx.Out
	courses.Gateway
}

func SetupCoursesGateway(postgresClient db_client.PostgresClient) CoursesGatewayModule {
	return CoursesGatewayModule{
		Gateway: &CoursesGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *CoursesGatewayImpl) CreateCourse(course *models.CourseCore) (id string, err error) {
	courseDb := models.CourseDB{}
	courseDb.FromCore(course)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&courseDb).Error
		return
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	id = strconv.FormatUint(uint64(courseDb.ID), 10)
	return id, nil
}

func (r *CoursesGatewayImpl) CreateAbsoluteMedia(absoluteMedia *models.AbsoluteMediaCore) (id string, err error) {
	absoluteMediaDb := models.AbsoluteMediaDB{}
	absoluteMediaDb.FromCore(absoluteMedia)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&absoluteMediaDb).Error
		return
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	id = strconv.FormatUint(uint64(absoluteMediaDb.ID), 10)
	return id, nil
}

func (r *CoursesGatewayImpl) CreateMedia(media *models.MediaCore) (id string, err error) {
	mediaDb := models.MediaDB{}
	mediaDb.FromCore(media)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&mediaDb).Error
		return
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	id = strconv.FormatUint(uint64(mediaDb.ID), 10)
	return id, nil
}

func (r *CoursesGatewayImpl) CreateImage(image *models.ImageCore) (id string, err error) {
	imageDb := models.ImageDB{}
	imageDb.FromCore(image)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&imageDb).Error
		return
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	id = strconv.FormatUint(uint64(imageDb.ID), 10)
	return id, nil
}

func (r *CoursesGatewayImpl) CreateCourseApiMediaCollection(courseApiMediaCollection *models.CourseApiMediaCollectionCore) (id string, err error) {
	courseApiMediaCollectionDb := models.CourseApiMediaCollectionDB{}
	courseApiMediaCollectionDb.FromCore(courseApiMediaCollection)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&courseApiMediaCollectionDb).Error
		return
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	id = strconv.FormatUint(uint64(courseApiMediaCollectionDb.ID), 10)
	return id, nil
}

func (r CoursesGatewayImpl) DeleteCourse(courseId string) (err error) {

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("str_course_id = ?", courseId).Delete(&models.CourseDB{}).Error
		fmt.Println(err)
		return
	})

	return
}

func (r *CoursesGatewayImpl) UpdateCourse(course *models.CourseCore) (err error) {
	courseDb := models.CourseDB{}
	courseDb.FromCore(course)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&courseDb).Where("ID = ?", courseDb.ID).Updates(courseDb).Error
		return
	})
	return
}
