package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"go.uber.org/fx"
)

type ProjectPageDelegateImpl struct {
	projectPage.UseCase
}

type ProjectPageDelegateModule struct {
	fx.Out
	projectPage.Delegate
}

func SetupProjectPageDelegate(usecase projectPage.UseCase) ProjectPageDelegateModule {
	return ProjectPageDelegateModule{
		Delegate: &ProjectPageDelegateImpl{
			UseCase: usecase,
		},
	}
}

func (p *ProjectPageDelegateImpl) CreateProjectPage(authorId string) (projectId string, err error) {
	return p.UseCase.CreateProjectPage(authorId)
}

func (p *ProjectPageDelegateImpl) DeleteProjectPage() {

}

func (p *ProjectPageDelegateImpl) UpdateProjectPage(projectPage *models.ProjectPageHTTP) (err error) {
	projectPageCore := projectPage.ToCore()
	return p.UseCase.UpdateProjectPage(&projectPageCore)
}

func (p *ProjectPageDelegateImpl) GetProjectPage() {

}
