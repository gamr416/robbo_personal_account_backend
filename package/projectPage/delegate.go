package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateProjectPage(authorId string) (projectId string, err error)
	DeleteProjectPage(projectId string) (err error)
	GetProjectPageById(projectId string) (projectPage models.ProjectPageHTTP, err error)
	GetAllProjectPages(authorId string) (projectPages []*models.ProjectPageHTTP, err error)
	UpdateProjectPage(projectPage *models.ProjectPageHTTP) (err error)
}
