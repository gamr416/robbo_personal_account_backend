package edx

type UseCase interface {
	AuthUseCase
	UserUseCase
	CohortUseCase
	CourseUseCase
}
