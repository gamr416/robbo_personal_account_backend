package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"go.uber.org/fx"
<<<<<<< HEAD
	"gorm.io/gorm"
	"strconv"
=======
>>>>>>> b51413c19b53a2b40776b2746be8d694d6f8e40e
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

/*
func (r *AuthGatewayImpl) GetUser(email, password string) (user *models.UserBase, err error) {
	var userDb models.UserDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("email = ? AND  password = ?", email, password).First(&userDb).Error; err != nil {
			err = auth.ErrUserNotFound
			return
		}
		return
	})
	user = userDb.ToCore()
	return
}
<<<<<<< HEAD

func (r *AuthGatewayImpl) CreateUser(user *models.UserCore) (id string, err error) {
	userDb := models.UserDB{}
	userDb.FromCore(user)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where(models.UserDB{Email: user.Email}).Take(&models.UserDB{}).Error; err == nil {
			err = auth.ErrUserAlreadyExist
			return
		}
		err = tx.Create(&userDb).Error
		return
	})

	id = strconv.FormatUint(uint64(userDb.ID), 10)
	return
}
=======
*/
>>>>>>> b51413c19b53a2b40776b2746be8d694d6f8e40e
