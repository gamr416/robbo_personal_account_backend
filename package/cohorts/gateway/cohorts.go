package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/cohorts"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
)

type CohortsGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type CohortsGatewayModule struct {
	fx.Out
	cohorts.Gateway
}

func SetupCohortsGateway(postgresClient db_client.PostgresClient) CohortsGatewayModule {
	return CohortsGatewayModule{
		Gateway: &CohortsGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *CohortsGatewayImpl) CreateCohort(cohort *models.CohortCore) (id string, err error) {
	return "", nil
}
