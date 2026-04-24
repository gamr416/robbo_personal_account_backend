package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"go.uber.org/fx"
	"gorm.io/gorm"
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

func (r *ProjectPageGatewayImpl) CreateProjectPage(projectPage *models.ProjectPageCore) (newProjectPage *models.ProjectPageCore, err error) {
	projectPageDb := models.ProjectPageDB{}
	projectPageDb.FromCore(projectPage)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&projectPageDb).Error
		return
	})

	newProjectPage = projectPageDb.ToCore()
	return
}

func (r *ProjectPageGatewayImpl) GetProjectPageById(projectPageId string) (projectPage *models.ProjectPageCore, err error) {
	var projectPageDB models.ProjectPageDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", projectPageId).First(&projectPageDB).Error; err != nil {
			return
		}
		return
	})

	projectPage = projectPageDB.ToCore()

	return
}

func (r *ProjectPageGatewayImpl) GetProjectPageByProjectId(projectId string) (projectPage *models.ProjectPageCore, err error) {
	var projectPageDB models.ProjectPageDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("project_id = ?", projectId).First(&projectPageDB).Error; err != nil {
			return
		}
		return
	})

	projectPage = projectPageDB.ToCore()

	return
}

func (r *ProjectPageGatewayImpl) DeleteProjectPage(projectId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.ProjectPageDB{}, projectId).Error
		return
	})
	return
}

func (r *ProjectPageGatewayImpl) UpdateProjectPage(core *models.ProjectPageCore) (projectPageUpdated *models.ProjectPageCore, err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		updates := map[string]interface{}{
			"instruction": core.Instruction,
			"notes":       core.Notes,
			"title":       core.Title,
			"is_shared":   core.IsShared,
		}
		res := tx.Model(&models.ProjectPageDB{}).Where("id = ?", core.ProjectPageId).Updates(updates)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return projectPage.ErrPageNotFound
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return r.GetProjectPageById(core.ProjectPageId)
}
