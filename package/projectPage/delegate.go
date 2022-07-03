package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateProjectPage()
	DeleteProjectPage()
	GetProjectPage()
	UpdateProjectPage(projectPage *models.ProjectPageHTTP) (err error)
}
