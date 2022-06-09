package projects

type Gateway interface {
	CreateProject()
	DeleteProject()
	GetProject()
	UpdateProject()
}
