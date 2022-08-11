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

func SetupProjectPageGateway(postgresClient db_client.PostgresClient) RobboUnitsGatewayModule {
	return RobboUnitsGatewayModule{
		Gateway: &RobboUnitsGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *RobboUnitsGatewayImpl) CreateRobboUnit(robboUnit *models.RobboUnitCore) (robboUnitId string, err error) {
	robboUnitDb := models.RobboUnitDB{}
	robboUnitDb.FromCore(robboUnit)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&robboUnitDb).Error
		return
	})

	robboUnitId = strconv.FormatUint(uint64(robboUnitDb.ID), 10)

	return
}
func (r *RobboUnitsGatewayImpl) DeleteRobboUnit(robboUnitId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.RobboUnitDB{}, robboUnitId).Error
		return
	})
	return
}

func (r *RobboUnitsGatewayImpl) GetAllRobboUnit() (robboUnits []*models.RobboUnitCore, err error) {
	var robboUnitsDB []*models.RobboUnitDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Find(&robboUnitsDB).Error; err != nil {
			return
		}
		return
	})

	for _, robboUnitDb := range robboUnitsDB {
		robboUnits = append(robboUnits, robboUnitDb.ToCore())
	}
	return
}

func (r *RobboUnitsGatewayImpl) GetRobboUnitById(robboUnitId string) (robboUnit *models.RobboUnitCore, err error) {
	var robboUnitDb models.RobboUnitDB

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", robboUnitId).First(&robboUnitDb).Error; err != nil {
			// TODO init err robboUnit not found
			log.Println(err)
			return
		}
		return
	})
	robboUnit = robboUnitDb.ToCore()
	return
}

func (r *RobboUnitsGatewayImpl) UpdateRobboUnit(robboUnit *models.RobboUnitCore) (err error) {
	robboUnitDb := models.RobboUnitDB{}
	robboUnitDb.FromCore(robboUnit)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&robboUnitDb).Where("id = ?", robboUnitDb.ID).Updates(robboUnitDb).Error
		return
	})
	return
}
