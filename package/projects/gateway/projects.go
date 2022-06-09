package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"go.uber.org/fx"
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

func (r *ProjectsGatewayImpl) CreateProject() {

}

func (r *ProjectsGatewayImpl) GetProject() {

}

func (r *ProjectsGatewayImpl) DeleteProject() {

}

func (r *ProjectsGatewayImpl) UpdateProject() {

}
