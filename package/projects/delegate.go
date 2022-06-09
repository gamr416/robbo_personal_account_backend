package projects

type Delegate interface {
	CreateProject()
	DeleteProject()
	GetProject()
	UpdateProject()
}
