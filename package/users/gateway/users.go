package gateway

import (
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type UsersGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type UsersGatewayModule struct {
	fx.Out
	users.Gateway
}

func SetupUsersGateway(postgresClient db_client.PostgresClient) UsersGatewayModule {
	return UsersGatewayModule{
		Gateway: &UsersGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *UsersGatewayImpl) GetUsersByRole(role *models.Role) (users *[]models.UserCore, err error) {
	var usersDb []models.UserDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("role = ?", role).Find(&usersDb).Error; err != nil {
			return
		}
		log.Println(err)
		return
	})
	var usersToCore []models.UserCore
	for _, user := range usersDb {
		usersToCore = append(usersToCore, *user.ToCore())
	}
	return &usersToCore, err
}

func (r *UsersGatewayImpl) UpdateUser(user *models.UserCore) (err error) {
	userDb := models.UserDB{}
	userDb.FromCore(user)
	fmt.Println(userDb)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", userDb.ID).First(&models.UserDB{}).Error
		if err != nil {
			log.Println(err)
			return auth.ErrUserNotFound
		}
		err = tx.Model(&userDb).Where("id = ?", userDb.ID).Updates(userDb).Error
		if err != nil {
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *UsersGatewayImpl) DeleteUser(userId int) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", userId).Delete(&models.UserDB{}).Error
		if err != nil {
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *UsersGatewayImpl) GetUserById(userId int) (user *models.UserCore, err error) {
	var userDb models.UserDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", userId).First(&userDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	user = userDb.ToCore()
	return
}

func (r *UsersGatewayImpl) CreateUser(user *models.UserCore) (id string, err error) {
	userDb := models.UserDB{}
	userDb.FromCore(user)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where(models.UserDB{Email: user.Email}).Take(&models.UserDB{}).Error; err == nil {
			err = auth.ErrUserAlreadyExist
			log.Println(err)
			return
		}
		err = tx.Create(&userDb).Error
		return
	})

	id = strconv.FormatUint(uint64(userDb.ID), 10)
	return
}
