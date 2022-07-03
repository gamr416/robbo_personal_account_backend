package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateProjectPage(authorId string) (projectId string, err error)
	DeleteProjectPage()
	GetProjectPage()
	UpdateProjectPage(projectPage *models.ProjectPageHTTP) (err error)
}
