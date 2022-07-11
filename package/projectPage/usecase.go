package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateProjectPage(authorId string) (projectId string, err error)
	DeleteProjectPage(projectId string) (err error)
	GetAllProjectPage(authorId string) (projectPages []*models.ProjectPageCore, err error)
	GetProjectPageById(projectId string) (projectPage *models.ProjectPageCore, err error)
	UpdateProjectPage(projectPage *models.ProjectPageCore) (err error)
}
