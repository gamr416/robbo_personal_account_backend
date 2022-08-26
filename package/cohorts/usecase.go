package cohorts

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateCohort(cohort *models.CohortCore) (id string, err error)
}
