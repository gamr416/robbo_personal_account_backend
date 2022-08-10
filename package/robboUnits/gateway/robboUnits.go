package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	"go.uber.org/fx"
)

type RobboUnitsGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type RobboUnitsGatewayModule struct {
	fx.Out
	robboUnits.Gateway
}

func SetupProjectPageGateway(postgresClient db_client.PostgresClient) RobboUnitsGatewayModule {
	return RobboUnitsGatewayModule{
		Gateway: &RobboUnitsGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *RobboUnitsGatewayImpl) CreateRobboUnit() (robboUnitId string, err error) {
	return
}
func (r *RobboUnitsGatewayImpl) DeleteRobboUnit(projectId string) (err error) {
	return
}

func (r *RobboUnitsGatewayImpl) GetAllRobboUnit(authorId string) (robboUnits []*models.RobboUnitCore, err error) {
	return
}

func (r *RobboUnitsGatewayImpl) GetRobboUnitById(robboUnitId string) (robboUnit *models.RobboUnitCore, err error) {
	return
}

func (r *RobboUnitsGatewayImpl) UpdateRobboUnit(projectPage *models.RobboUnitCore) (err error) {
	return
}