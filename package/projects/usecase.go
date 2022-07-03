package projects

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateProject(project *models.ProjectCore) (id string, err error)
	DeleteProject()
	GetProjectById(projectId string) (project *models.ProjectCore, err error)
	UpdateProject(project *models.ProjectCore) (err error)
}
