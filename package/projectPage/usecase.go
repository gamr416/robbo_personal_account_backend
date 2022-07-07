package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateProjectPage(projectPage *models.ProjectPageCore) (id string, err error)
	DeleteProjectPage(projectPage *models.ProjectPageCore) (err error)
	GetProjectPageByID()
	UpdateProjectPage(projectPage *models.ProjectPageCore) (err error)
}
