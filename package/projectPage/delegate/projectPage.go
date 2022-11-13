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

func (p *ProjectPageDelegateImpl) DeleteProjectPage(projectId string) (err error) {
	return p.UseCase.DeleteProjectPage(projectId)
}

func (p *ProjectPageDelegateImpl) UpdateProjectPage(projectPage *models.ProjectPageHTTP) (err error) {
	projectPageCore := projectPage.ToCore()
	return p.UseCase.UpdateProjectPage(projectPageCore)
}

func (p *ProjectPageDelegateImpl) GetProjectPageById(projectPageId string) (projectPage models.ProjectPageHTTP, err error) {
	projectPageCore, err := p.UseCase.GetProjectPageById(projectPageId)
	if err != nil {
		return
	}
	projectPage.FromCore(projectPageCore)
	return
}

func (p *ProjectPageDelegateImpl) GetAllProjectPagesByUserId(authorId string) (projectPages []*models.ProjectPageHTTP, err error) {
	projectPagesCore, err := p.UseCase.GetAllProjectPageByUserId(authorId)
	if err != nil {
		return
	}
	for _, projectPageCore := range projectPagesCore {
		var projectPageHttp models.ProjectPageHTTP
		projectPageHttp.FromCore(projectPageCore)
		projectPages = append(projectPages, &projectPageHttp)
	}
	return
}
