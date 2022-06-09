package projects

type UseCase interface {
	CreateProject()
	DeleteProject()
	GetProject()
	UpdateProject()
}
