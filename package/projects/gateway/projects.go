package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"strconv"
)

type ProjectsGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type ProjectsGatewayModule struct {
	fx.Out
	projects.Gateway
}

func SetupProjectsGateway(postgresClient db_client.PostgresClient) ProjectsGatewayModule {
	return ProjectsGatewayModule{
		Gateway: &ProjectsGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *ProjectsGatewayImpl) CreateProject(project *models.ProjectCore) (id string, err error) {
	projectDb := models.ProjectDB{}
	projectDb.FromCore(project)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&projectDb).Error
		return
	})

	id = strconv.FormatUint(uint64(projectDb.ID), 10)
	return
}

func (r *ProjectsGatewayImpl) GetProject() {

}

func (r *ProjectsGatewayImpl) DeleteProject() {

}

func (r *ProjectsGatewayImpl) UpdateProject(project *models.ProjectCore) (err error) {
	projectDb := models.ProjectDB{}
	projectDb.FromCore(project)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&projectDb).Where("ID = ?", projectDb.ID).Updates(projectDb).Error
		return
	})
	return
}
