package model

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

func (gm *Parent) ToCore() *models.ParentCore {
	return &models.ParentCore{
		UserCore: gm.User.ToCore(),
		// Children не нужно???
	}
}

func (gm *User) ToCore() models.UserCore {
	return models.UserCore{
		Id:         gm.ID,
		Email:      gm.Email,
		Password:   gm.Password,
		Role:       models.Role(gm.Role),
		Nickname:   gm.Nickname,
		Firstname:  gm.Firstname,
		Lastname:   gm.Lastname,
		Middlename: gm.Middlename,
		CreatedAt:  gm.CreatedAt,
	}
}

func (gm *User) FromCore(user *models.UserCore) {
	gm.ID = user.Id
	gm.Email = user.Email
	gm.Password = user.Password
	gm.Role = int(user.Role)
	gm.Nickname = user.Nickname
	gm.Firstname = user.Firstname
	gm.Lastname = user.Lastname
	gm.Middlename = user.Middlename
	gm.CreatedAt = user.CreatedAt
}
