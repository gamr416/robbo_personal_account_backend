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

func (r *RobboGroupGatewayImpl) SearchRobboGroupsByTitle(title string) (robboGroupsCore []*models.RobboGroupCore, err error) {
	var robboGroupsDB []*models.RobboGroupDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Limit(10).Where("name LIKE ?", title).Find(&robboGroupsDB).Error; err != nil {
			return
		}
		return
	})
	for _, robboGroupDB := range robboGroupsDB {
		robboGroupsCore = append(robboGroupsCore, robboGroupDB.ToCore())
	}
	return
}

func (r *RobboGroupGatewayImpl) CreateRobboGroup(robboGroupCore *models.RobboGroupCore) (robboGroupId string, err error) {
	robboGroupDb := models.RobboGroupDB{}
	robboGroupDb.FromCore(robboGroupCore)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&robboGroupDb).Error
		return
	})

	robboGroupId = strconv.FormatUint(uint64(robboGroupDb.ID), 10)

	return
}

func (r *RobboGroupGatewayImpl) DeleteRobboGroup(robboGroupId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&models.RobboGroupDB{}).Where("id = ?", robboGroupId).First(&models.RobboGroupDB{}).Delete(&models.RobboGroupDB{}).Error; err != nil {
			err = robboGroup.ErrRobboGroupNotFound
			return
		}
		return
	})
	return
}

func (r *RobboGroupGatewayImpl) GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroupsCore []*models.RobboGroupCore, err error) {
	var robboGroupsDB []*models.RobboGroupDB

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("robbo_unit_id = ?", robboUnitId).Find(&robboGroupsDB).Error; err != nil {
			return
		}
		return
	})

	for _, robboGroupDb := range robboGroupsDB {
		robboGroupsCore = append(robboGroupsCore, robboGroupDb.ToCore())
	}
	return
}

func (r *RobboGroupGatewayImpl) GetRobboGroupById(robboGroupId string) (robboGroupCore *models.RobboGroupCore, err error) {
	var robboGroupDB models.RobboGroupDB

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", robboGroupId).First(&robboGroupDB).Error; err != nil {
			err = robboGroup.ErrRobboGroupNotFound
			log.Println(err)
			return
		}
		return
	})
	robboGroupCore = robboGroupDB.ToCore()
	return
}

func (r *RobboGroupGatewayImpl) SetTeacherForRobboGroup(relation *models.TeachersRobboGroupsCore) (err error) {
	relationDb := models.TeachersRobboGroupsDB{}
	relationDb.FromCore(relation)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&relationDb).Error
		return
	})

	return
}

func (r *RobboGroupGatewayImpl) DeleteTeacherForRobboGroup(relation *models.TeachersRobboGroupsCore) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.TeachersRobboGroupsDB{}, relation).Error
		return
	})
	return
}

func (r *RobboGroupGatewayImpl) DeleteRelationByRobboGroupId(robboGroupId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("robbo_group_id = ?", robboGroupId).Delete(&models.TeachersRobboGroupsDB{}).Error
		return
	})
	return
}

func (r *RobboGroupGatewayImpl) DeleteRelationByTeacherId(teacherId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("teacher_id = ?", teacherId).Delete(&models.TeachersRobboGroupsDB{}).Error
		return
	})
	return
}

func (r *RobboGroupGatewayImpl) GetRelationByRobboGroupId(robboGroupId string) (relations []*models.TeachersRobboGroupsCore, err error) {
	var relationsDB []*models.TeachersRobboGroupsDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("robbo_group_id = ?", robboGroupId).Find(&relationsDB).Error; err != nil {
			return
		}
		return
	})

	for _, relationDB := range relationsDB {
		relations = append(relations, relationDB.ToCore())
	}
	return
}

func (r *RobboGroupGatewayImpl) GetRelationByTeacherId(teacherId string) (relations []*models.TeachersRobboGroupsCore, err error) {
	var relationsDB []*models.TeachersRobboGroupsDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("teacher_id = ?", teacherId).Find(&relationsDB).Error; err != nil {
			return
		}
		return
	})

	for _, relationDB := range relationsDB {
		relations = append(relations, relationDB.ToCore())
	}
	return
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
