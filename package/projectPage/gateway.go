package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateProjectPage(projectPage *models.ProjectPageCore) (projectPageId string, err error)
	DeleteProjectPage(projectId string) (err error)
	GetProjectPageByID()
	UpdateProjectPage(projectPage *models.ProjectPageCore) (err error)
}
