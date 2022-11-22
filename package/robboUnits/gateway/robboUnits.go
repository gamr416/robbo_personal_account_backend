package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type RobboUnitsGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type RobboUnitsGatewayModule struct {
	fx.Out
	robboUnits.Gateway
}

func SetupRobboUnitsGateway(postgresClient db_client.PostgresClient) RobboUnitsGatewayModule {
	return RobboUnitsGatewayModule{
		Gateway: &RobboUnitsGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *RobboUnitsGatewayImpl) CreateRobboUnit(robboUnitCore *models.RobboUnitCore) (robboUnitId string, err error) {
	robboUnitDb := models.RobboUnitDB{}
	robboUnitDb.FromCore(robboUnitCore)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&robboUnitDb).Error
		return
	})

	robboUnitId = strconv.FormatUint(uint64(robboUnitDb.ID), 10)

	return
}
func (r *RobboUnitsGatewayImpl) DeleteRobboUnit(robboUnitId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&models.RobboUnitDB{}).Where("id = ?", robboUnitId).First(&models.RobboUnitDB{}).Delete(&models.RobboUnitDB{}).Error; err != nil {
			err = robboUnits.ErrRobboUnitNotFound
			return
		}
		return
	})
	return
}

func (r *RobboUnitsGatewayImpl) GetAllRobboUnit() (robboUnitsCore []*models.RobboUnitCore, err error) {
	var robboUnitsDB []*models.RobboUnitDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Find(&robboUnitsDB).Error; err != nil {
			return
		}
		return
	})

	for _, robboUnitDb := range robboUnitsDB {
		robboUnitsCore = append(robboUnitsCore, robboUnitDb.ToCore())
	}
	return
}

func (r *RobboUnitsGatewayImpl) GetRobboUnitById(robboUnitId string) (robboUnitCore *models.RobboUnitCore, err error) {
	var robboUnitDb models.RobboUnitDB

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", robboUnitId).First(&robboUnitDb).Error; err != nil {
			err = robboUnits.ErrRobboUnitNotFound
			log.Println(err)
			return
		}
		return
	})
	robboUnitCore = robboUnitDb.ToCore()
	return
}

func (r *RobboUnitsGatewayImpl) UpdateRobboUnit(robboUnitCore *models.RobboUnitCore) (err error) {
	robboUnitDb := models.RobboUnitDB{}
	robboUnitDb.FromCore(robboUnitCore)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&robboUnitDb).Where("id = ?", robboUnitDb.ID).First(&models.RobboUnitDB{}).Updates(robboUnitDb).Error; err != nil {
			err = robboUnits.ErrRobboUnitNotFound
			return
		}
		return
	})
	return
}
