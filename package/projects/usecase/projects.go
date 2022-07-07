package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"go.uber.org/fx"
)

type ProjectUseCaseImpl struct {
	projects     projects.Gateway
	projectPages projectPage.Gateway
}

type ProjectUseCaseModule struct {
	fx.Out
	projects.UseCase
}

func SetupProjectUseCase(gateway projects.Gateway, PPgateway projectPage.Gateway) ProjectUseCaseModule {
	return ProjectUseCaseModule{
		UseCase: &ProjectUseCaseImpl{
			projects:     gateway,
			projectPages: PPgateway,
		},
	}
}

func (p *ProjectUseCaseImpl) CreateProject(project *models.ProjectCore) (id string, err error) {
	return p.projects.CreateProject(project)
}

func (p *ProjectUseCaseImpl) UpdateProject(project *models.ProjectCore) (err error) {
	return p.projects.UpdateProject(project)
}

func (p *ProjectUseCaseImpl) DeleteProject() {

}

func (p *ProjectUseCaseImpl) GetProjectById(projectId string) (project *models.ProjectCore, err error) {
	return p.Gateway.GetProjectById(projectId)
}
