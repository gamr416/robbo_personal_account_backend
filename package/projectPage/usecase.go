package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateProjectPage(authorId string) (newProjectPage *models.ProjectPageCore, err error)
	UpdateProjectPage(projectPage *models.ProjectPageCore) (projectPageUpdated *models.ProjectPageCore, err error)
	DeleteProjectPage(projectId string) (err error)
	GetAllProjectPageByUserId(authorId string, page int, pageSize int) (projectPages []*models.ProjectPageCore, countRows int64, err error)
	GetProjectPageById(projectPageId string) (projectPage *models.ProjectPageCore, err error)
}
