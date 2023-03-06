package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"go.uber.org/fx"
	"log"
	"strconv"
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

func (p *ProjectPageDelegateImpl) CreateProjectPage(authorId string) (newProjectPage models.ProjectPageHTTP, err error) {
	newProjectPageCore, err := p.UseCase.CreateProjectPage(authorId)
	if err != nil {
		log.Println(err)
		return
	}
	newProjectPage.FromCore(newProjectPageCore)
	return
}

func (p *ProjectPageDelegateImpl) DeleteProjectPage(projectId string) (err error) {
	return p.UseCase.DeleteProjectPage(projectId)
}

func (p *ProjectPageDelegateImpl) UpdateProjectPage(projectPage *models.ProjectPageHTTP) (projectPageUpdated models.ProjectPageHTTP, err error) {
	projectPageCore := projectPage.ToCore()
	projectPageUpdatedCore, err := p.UseCase.UpdateProjectPage(projectPageCore)
	if err != nil {
		log.Println(err)
		return
	}
	projectPageUpdated.FromCore(projectPageUpdatedCore)
	return
}

func (p *ProjectPageDelegateImpl) GetProjectPageById(projectPageId string) (projectPage models.ProjectPageHTTP, err error) {
	projectPageCore, err := p.UseCase.GetProjectPageById(projectPageId)
	if err != nil {
		return
	}
	projectPage.FromCore(projectPageCore)
	return
}

func (p *ProjectPageDelegateImpl) GetAllProjectPagesByUserId(authorId, page, pageSize string) (
	projectPages []*models.ProjectPageHTTP,
	countRows int,
	err error,
) {
	pageInt32, _ := strconv.ParseInt(page, 10, 32)
	pageSizeInt32, _ := strconv.ParseInt(pageSize, 10, 32)
	projectPagesCore, countRowsInt64, err := p.UseCase.GetAllProjectPageByUserId(
		authorId,
		int(pageInt32),
		int(pageSizeInt32),
	)
	if err != nil {
		return
	}
	countRows = int(countRowsInt64)
	for _, projectPageCore := range projectPagesCore {
		var projectPageHttp models.ProjectPageHTTP
		projectPageHttp.FromCore(projectPageCore)
		projectPages = append(projectPages, &projectPageHttp)
	}
	return
}
