package gateway

import "go.uber.org/fx"

type Module struct {
	fx.Out
	AuthGateway
}

type AuthGateway interface {
	GetUser(email, password string) (err error)
}

func Setup(postgresClient PostgresClient) Module {
	return Module{
		AuthGateway: &AuthGatewayImpl{postgresClient: &postgresClient},
	}
}
