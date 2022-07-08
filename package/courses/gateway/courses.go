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

func (r CoursesGatewayImpl) DeleteCourse(courseId string) (err error) {
	courseDb := models.CourseDB{}

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&courseDb).Where("course_id = ?", courseId).Delete(&courseDb).Error
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
