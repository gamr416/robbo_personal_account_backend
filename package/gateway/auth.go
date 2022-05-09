package gateway

type AuthGatewayImpl struct {
	postgresClient *PostgresClient
}

func (r *AuthGatewayImpl) GetUser(email, password string) (err error) {
	// TODO implement me
	return nil
}
