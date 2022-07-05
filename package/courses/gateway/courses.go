package gateway

import (
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
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

func (r *CoursesGatewayImpl) CreateCourse(course *models.CourseCore, courseId string) (id string, err error) {
	courseDb := models.CourseDB{}
	courseDb.FromCore(course)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&courseDb).Error
		return
	})

	id = strconv.FormatUint(uint64(courseDb.ID), 10)
	return
}

func (r *CoursesGatewayImpl) DeleteCourse(course *models.CourseCore) (err error) {
	courseDb := models.CourseDB{}
	courseDb.FromCore(course)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&courseDb).Where("course_id = ?", courseDb.CourseID).Select(&courseDb).Error
		fmt.Println(err)
		return
	})

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&courseDb).Where("course_id = ?", courseDb.CourseID).Delete(&courseDb).Error
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
