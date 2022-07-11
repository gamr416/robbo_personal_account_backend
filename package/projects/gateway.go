package projects

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateProject(project *models.ProjectCore) (id string, err error)
	DeleteProject(projectId string) (err error)
	GetProjectById(projectId string) (project *models.ProjectCore, err error)
	GetProjectsByAuthorId(authorId string) (projects []*models.ProjectCore, err error)
	UpdateProject(project *models.ProjectCore) (err error)
}
