package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
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

func (r *AuthGatewayImpl) GetUser(email, password string) (user *models.UserCore, err error) {
	var userDb models.UserDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("email = ? AND  = ?", email, password).First(&userDb).Error; err != nil {
			return
		}
		return
	})
	user = userDb.ToCore()
	return
}

func (r *AuthGatewayImpl) CreateUser(user *models.UserCore) (err error) {
	userDb := models.UserDB{}
	userDb.FromCore(user)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Create(&userDb).Error; err != nil {
			return
		}
		return
	})
	return
}
