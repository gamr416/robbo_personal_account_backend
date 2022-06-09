package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"go.uber.org/fx"
)

type ProjectDelegateImpl struct {
	projects.UseCase
}

type ProjectDelegateModule struct {
	fx.Out
	projects.Delegate
}

func SetupProjectDelegate(usecase projects.UseCase) ProjectDelegateModule {
	return ProjectDelegateModule{
		Delegate: &ProjectDelegateImpl{usecase},
	}
}

func (p *ProjectDelegateImpl) CreateProject() {

}

func (p *ProjectDelegateImpl) DeleteProject() {

}

func (p *ProjectDelegateImpl) UpdateProject() {

}

func (p *ProjectDelegateImpl) GetProject() {

}
