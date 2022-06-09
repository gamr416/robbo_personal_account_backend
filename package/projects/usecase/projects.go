package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"go.uber.org/fx"
)

type ProjectUseCaseImpl struct {
	projects.Gateway
}

type ProjectUseCaseModule struct {
	fx.Out
	projects.UseCase
}

func SetupProjectUseCase(gateway projects.Gateway) ProjectUseCaseModule {
	return ProjectUseCaseModule{
		UseCase: &ProjectUseCaseImpl{
			Gateway: gateway,
		},
	}
}

func (p *ProjectUseCaseImpl) CreateProject() {

}

func (p *ProjectUseCaseImpl) UpdateProject() {

}

func (p *ProjectUseCaseImpl) DeleteProject() {

}

func (p *ProjectUseCaseImpl) GetProject() {

}