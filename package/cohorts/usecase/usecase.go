package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/cohorts"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
)

type CohortUseCaseImpl struct {
	cohorts.Gateway
}

type CohortUseCaseModule struct {
	fx.Out
	cohorts.UseCase
}

func SetupCohortUseCase(gateway cohorts.Gateway) CohortUseCaseModule {
	return CohortUseCaseModule{
		UseCase: &CohortUseCaseImpl{
			Gateway: gateway,
		},
	}
}

func (p *CohortUseCaseImpl) CreateCohort(cohort *models.CohortCore) (id string, err error) {
	return p.Gateway.CreateCohort(cohort)
}
