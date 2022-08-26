package edx

type CohortUseCase interface {
	CreateCohort(courseId string, cohortParams map[string]interface{}) (respBody []byte, err error)
	AddStudent(username, courseId string, cohortId int) (respBody []byte, err error)
}
