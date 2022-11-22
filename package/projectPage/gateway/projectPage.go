package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"strconv"
)

type ProjectPageGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type ProjectPageGatewayModule struct {
	fx.Out
	projectPage.Gateway
}

func SetupProjectPageGateway(postgresClient db_client.PostgresClient) ProjectPageGatewayModule {
	return ProjectPageGatewayModule{
		Gateway: &ProjectPageGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *ProjectPageGatewayImpl) CreateProjectPage(projectPageCore *models.ProjectPageCore) (projectPageId string, err error) {
	projectPageDb := models.ProjectPageDB{}
	projectPageDb.FromCore(projectPageCore)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&projectPageDb).Error
		return
	})

	projectPageId = strconv.FormatUint(uint64(projectPageDb.ID), 10)
	return
}

func (r *ProjectPageGatewayImpl) GetProjectPageById(projectPageId string) (projectPageCore *models.ProjectPageCore, err error) {
	var projectPageDB models.ProjectPageDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", projectPageId).First(&projectPageDB).Error; err != nil {
			err = projectPage.ErrPageNotFound
			return
		}
		return
	})

	projectPageCore = projectPageDB.ToCore()

	return
}

func (r *ProjectPageGatewayImpl) GetProjectPageByProjectId(projectId string) (projectPageCore *models.ProjectPageCore, err error) {
	var projectPageDB models.ProjectPageDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("project_id = ?", projectId).First(&projectPageDB).Error; err != nil {
			err = projectPage.ErrPageNotFound
			return
		}
		return
	})

	projectPageCore = projectPageDB.ToCore()
	return
}

func (r *ProjectPageGatewayImpl) DeleteProjectPage(projectId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&models.ProjectPageDB{}).Where("project_id = ?", projectId).First(&models.ProjectPageDB{}).Delete(&models.ProjectPageDB{}).Error; err != nil {
			err = projectPage.ErrPageNotFound
			return
		}
		return
	})
	return
}

func (r *ProjectPageGatewayImpl) UpdateProjectPage(projectPageCore *models.ProjectPageCore) (err error) {
	projectPageDb := models.ProjectPageDB{}
	projectPageDb.FromCore(projectPageCore)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&projectPageDb).Where("project_id = ?", projectPageDb.ProjectId).First(&models.ProjectPageDB{}).Updates(projectPageDb).Error; err != nil {
			err = projectPage.ErrPageNotFound
			return
		}
		return
	})
	return
}
