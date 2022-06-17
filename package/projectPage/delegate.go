package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateProjectPage(projectPage *models.ProjectPageHTTP, projects []*models.ProjectHTTP) (id string, err error)
	DeleteProjectPage()
	GetProjectPage()
	UpdateProjectPage(projectPage *models.ProjectPageHTTP, projects []*models.ProjectHTTP) (err error)
}
