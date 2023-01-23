package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
)

type RobboUnitsGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

func (r *RobboUnitsGatewayImpl) SearchRobboUnitByName(name string) (robboUnitsCore []*models.RobboUnitCore, err error) {
	var robboUnitsDb []*models.RobboUnitDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Limit(10).Where("name LIKE ?", name).Find(&robboUnitsDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	for _, robboUnitDb := range robboUnitsDb {
		robboUnitsCore = append(robboUnitsCore, robboUnitDb.ToCore())
	}
	return
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

func (r *RobboUnitsGatewayImpl) CreateRobboUnit(robboUnitCore *models.RobboUnitCore) (newRobboUnit *models.RobboUnitCore, err error) {
	robboUnitDb := models.RobboUnitDB{}
	robboUnitDb.FromCore(robboUnitCore)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&robboUnitDb).Error
		return
	})
	newRobboUnit = robboUnitDb.ToCore()
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

func (r *RobboUnitsGatewayImpl) GetAllRobboUnit(page, pageSize int) (
	robboUnitsCore []*models.RobboUnitCore,
	countRows int64,
	err error,
) {
	var robboUnitsDB []*models.RobboUnitDB
	offset := (page - 1) * pageSize
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Limit(pageSize).Offset(offset).Find(&robboUnitsDB).Error; err != nil {
			return
		}
		tx.Model(&models.RobboUnitDB{}).Count(&countRows)
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

func (r *RobboUnitsGatewayImpl) UpdateRobboUnit(robboUnitCore *models.RobboUnitCore) (robboUnitUpdated *models.RobboUnitCore, err error) {
	robboUnitDb := models.RobboUnitDB{}
	robboUnitDb.FromCore(robboUnitCore)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&robboUnitDb).Where("id = ?", robboUnitDb.ID).First(&models.RobboUnitDB{}).Updates(robboUnitDb).Error; err != nil {
			err = robboUnits.ErrRobboUnitNotFound
			return
		}
		return
	})
	robboUnitUpdated = robboUnitDb.ToCore()
	return
}
