package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"go.uber.org/fx"
)

type AuthGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type AuthGatewayModule struct {
	fx.Out
	auth.Gateway
}

func SetupAuthGateway(postgresClient db_client.PostgresClient) AuthGatewayModule {
	return AuthGatewayModule{
		Gateway: &AuthGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *AuthGatewayImpl) GetUser(email, password string) (err error) {
	// TODO implement me
	return nil
}
