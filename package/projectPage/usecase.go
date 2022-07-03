package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateProjectPage(authorId string) (projectId string, err error)
	DeleteProjectPage()
	GetProjectPageById()
	UpdateProjectPage(projectPage *models.ProjectPageCore) (err error)
}
