package gateway

//
//import (
//	"github.com/skinnykaen/robbo_student_personal_account.git/package/api"
//	"github.com/skinnykaen/robbo_student_personal_account.git/package/api/gateway/mock"
//	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestGetUser(t *testing.T) {
//	g := new(mock.UserGatewayMock)
//	user := &models.UserCore{
//		ID:       "id",
//		Email:    "email",
//		Password: "password",
//	}
//
//	err := g.CreateUser(user)
//	assert.NoError(t, err)
//
//	returnedUser, err := g.GetUser("user", "password")
//	assert.NoError(t, err)
//	assert.Equal(t, user, returnedUser)
//
//	returnedUser, err = g.GetUser("user", "")
//	assert.Error(t, err)
//	assert.Equal(t, err, api.ErrUserNotFound)
//}
