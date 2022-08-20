package gateway

import (
	"fmt"
	"log"
	"strconv"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/coursePacket"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type CoursePacketGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type CoursePacketGatewayModule struct {
	fx.Out
	coursePacket.Gateway
}

func SetupCoursePacketGateway(postgresClient db_client.PostgresClient) CoursePacketGatewayModule {
	return CoursePacketGatewayModule{
		Gateway: &CoursePacketGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *CoursePacketGatewayImpl) CreateCoursePacket(coursePacket *models.CoursePacketCore) (id string, err error) {
	coursePacketDb := models.CoursePacketDB{}
	coursePacketDb.FromCore(coursePacket)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&coursePacketDb).Error
		if err != nil {
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	id = strconv.FormatUint(uint64(coursePacketDb.ID), 10)
	return id, nil
}

func (r CoursePacketGatewayImpl) DeleteCoursePacket(coursePacketId string) (id string, err error) {
	coursePacket := models.CoursePacketDB{}
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", coursePacketId).First(&coursePacket).Error
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(coursePacket)
		err = tx.Model(&coursePacket).Where("id = ?", coursePacketId).Delete(&models.CoursePacketDB{}).Error
		if err != nil {
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	id = strconv.FormatUint(uint64(coursePacket.ID), 10)
	return
}

func (r *CoursePacketGatewayImpl) UpdateCoursePacket(coursePacket *models.CoursePacketCore) (err error) {
	coursePacketDb := models.CoursePacketDB{}
	coursePacketDb.FromCore(coursePacket)
	fmt.Println(coursePacketDb)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&coursePacketDb).Where("ID = ?", coursePacketDb.ID).Updates(coursePacketDb).Error
		if err != nil {
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *CoursePacketGatewayImpl) GetAllCoursePackets() (coursePackets []*models.CoursePacketCore, err error) {
	var coursePacketsDB []*models.CoursePacketDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Find(&coursePacketsDB).Error; err != nil {
			return
		}
		return
	})

	for _, coursePacketDB := range coursePacketsDB {
		coursePackets = append(coursePackets, coursePacketDB.ToCore())
	}
	return
}

func (r *CoursePacketGatewayImpl) GetCoursePacketById(coursePacketId string) (coursePacket *models.CoursePacketCore, err error) {
	var coursePacketDb models.CoursePacketDB

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", coursePacketId).First(&coursePacketDb).Error; err != nil {
			// TODO init err coursePacket not found
			log.Println(err)
			return
		}
		return
	})
	coursePacket = coursePacketDb.ToCore()
	return
}
