package projects

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateProject(project *models.ProjectCore) (id string, err error)
	DeleteProject()
	GetProject()
	UpdateProject(project *models.ProjectCore) (err error)
}
