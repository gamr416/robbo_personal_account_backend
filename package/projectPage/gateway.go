package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateProjectPage(projectPage *models.ProjectPageCore) (projectPageId string, err error)
	DeleteProjectPage(projectId string) (err error)
	GetProjectPageById(projectPageId string) (projectPage *models.ProjectPageCore, err error)
	GetProjectPageByProjectId(projectId string) (projectPage *models.ProjectPageCore, err error)
	UpdateProjectPage(projectPage *models.ProjectPageCore) (err error)
}
