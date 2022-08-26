package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type RobboGroupGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type RobboGroupGatewayModule struct {
	fx.Out
	robboGroup.Gateway
}

func SetupRobboGroupGateway(postgresClient db_client.PostgresClient) RobboGroupGatewayModule {
	return RobboGroupGatewayModule{
		Gateway: &RobboGroupGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *RobboGroupGatewayImpl) CreateRobboGroup(robboGroup *models.RobboGroupCore) (robboGroupId string, err error) {
	robboGroupDb := models.RobboGroupDB{}
	robboGroupDb.FromCore(robboGroup)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&robboGroupDb).Error
		return
	})

	robboGroupId = strconv.FormatUint(uint64(robboGroupDb.ID), 10)

	return
}

func (r *RobboGroupGatewayImpl) DeleteRobboGroup(robboGroupId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.RobboGroupDB{}, robboGroupId).Error
		return
	})
	return
}

func (r *RobboGroupGatewayImpl) GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroups []*models.RobboGroupCore, err error) {
	var robboGroupsDB []*models.RobboGroupDB

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("robbo_unit_id = ?", robboUnitId).Find(&robboGroupsDB).Error; err != nil {
			return
		}
		return
	})

	for _, robboGroupDb := range robboGroupsDB {
		robboGroups = append(robboGroups, robboGroupDb.ToCore())
	}
	return
}

func (r *RobboGroupGatewayImpl) GetRobboGroupById(robboGroupId string) (robboGroup *models.RobboGroupCore, err error) {
	var robboGroupDB models.RobboGroupDB

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", robboGroupId).First(&robboGroupDB).Error; err != nil {
			// TODO init err robboGroup not found
			log.Println(err)
			return
		}
		return
	})
	robboGroup = robboGroupDB.ToCore()
	return
}
