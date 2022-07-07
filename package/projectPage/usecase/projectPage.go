package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"go.uber.org/fx"
)

type ProjectPageUseCaseImpl struct {
	projectPage.Gateway
}

type ProjectPageUseCaseModule struct {
	fx.Out
	projectPage.UseCase
}

func SetupProjectPageUseCase(gateway projectPage.Gateway) ProjectPageUseCaseModule {
	return ProjectPageUseCaseModule{
		UseCase: &ProjectPageUseCaseImpl{
			Gateway: gateway,
		},
	}
}

func (p *ProjectPageUseCaseImpl) CreateProjectPage(projectPage *models.ProjectPageCore) (id string, err error) {
	return p.Gateway.CreateProjectPage(projectPage)
}

func (p *ProjectPageUseCaseImpl) UpdateProjectPage(projectPage *models.ProjectPageCore) (err error) {
	return p.Gateway.UpdateProjectPage(projectPage)
}

func (p *ProjectPageUseCaseImpl) DeleteProjectPage(projectPage *models.ProjectPageCore) (err error) {
	return p.Gateway.UpdateProjectPage(projectPage)
}

func (p *ProjectPageUseCaseImpl) GetProjectPageByID() {

}
