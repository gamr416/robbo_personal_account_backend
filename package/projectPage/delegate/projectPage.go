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

func SetupProjectDelegate(usecase projectPage.UseCase) ProjectPageDelegateModule {
	return ProjectPageDelegateModule{
		Delegate: &ProjectPageDelegateImpl{usecase},
	}
}

func (p *ProjectPageDelegateImpl) CreateProjectPage(projectPage *models.ProjectPageHTTP, projects []*models.ProjectHTTP) (id string, err error) {
	projectPageCore := projectPage.ToCore(projects)
	return p.UseCase.CreateProjectPage(&projectPageCore)
}

func (p *ProjectPageDelegateImpl) DeleteProjectPage() {

}

func (p *ProjectPageDelegateImpl) UpdateProjectPage(projectPage *models.ProjectPageHTTP, projects []*models.ProjectHTTP) (err error) {
	projectPageCore := projectPage.ToCore(projects)
	return p.UseCase.UpdateProjectPage(&projectPageCore)
}

func (p *ProjectPageDelegateImpl) GetProjectPage() {

}
