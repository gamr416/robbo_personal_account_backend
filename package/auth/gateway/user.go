package gateway

import (
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"strconv"
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
		if err = tx.Where("email = ? AND  password = ?", email, password).First(&userDb).Error; err != nil {
			return
		}
		return
	})
	user = userDb.ToCore()
	return
}

func (r *AuthGatewayImpl) CreateUser(user *models.UserCore) (id string, err error) {
	userDb := models.UserDB{}
	userDb.FromCore(user)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if tx.Where(models.UserDB{Email: user.Email}).Take(&models.UserDB{}).Error == nil {
			return fmt.Errorf("A user already exists with the given email address: %v", user.Email)
		}
		err = tx.Create(&userDb).Error
		return
	})

	id = strconv.FormatUint(uint64(userDb.ID), 10)
	return
}
