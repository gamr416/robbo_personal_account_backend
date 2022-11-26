package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateProjectPage(projectPageCore *models.ProjectPageCore) (projectPageId string, err error)
	DeleteProjectPage(projectId string) (err error)
	GetProjectPageById(projectPageId string) (projectPageCore *models.ProjectPageCore, err error)
	GetProjectPageByProjectId(projectId string) (projectPageCore *models.ProjectPageCore, err error)
	UpdateProjectPage(projectPageCore *models.ProjectPageCore) (err error)
}
