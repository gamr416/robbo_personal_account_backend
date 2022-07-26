package models

import (
	"gorm.io/gorm"
)

type SuperAdminHTTP struct {
	UserHttp
}

type SuperAdminDB struct {
	gorm.Model
	UserDB
}

type SuperAdminCore struct {
	UserCore
}

func (em *SuperAdminDB) ToCore() *SuperAdminCore {
	return &SuperAdminCore{
		UserCore{
			Email:      em.Email,
			Password:   em.Password,
			Role:       Role(em.Role),
			Nickname:   em.Nickname,
			Firstname:  em.Firstname,
			Lastname:   em.Lastname,
			Middlename: em.Middlename,
			CreatedAt:  em.CreatedAt.String(),
		},
	}
}

func (em *SuperAdminDB) FromCore(superAdmin *SuperAdminCore) {
	em.Email = superAdmin.Email
	em.Password = superAdmin.Password
	em.Role = uint(superAdmin.Role)
	em.Nickname = superAdmin.Nickname
	em.Firstname = superAdmin.Firstname
	em.Lastname = superAdmin.Lastname
	em.Middlename = superAdmin.Middlename
}

func (ht *SuperAdminHTTP) ToCore() *SuperAdminCore {
	return &SuperAdminCore{
		UserCore{
			Email:      ht.Email,
			Password:   ht.Password,
			Role:       Role(ht.Role),
			Nickname:   ht.Nickname,
			Firstname:  ht.Firstname,
			Lastname:   ht.Lastname,
			Middlename: ht.Middlename,
		},
	}
}

func (ht *SuperAdminHTTP) FromCore(superAdmin *SuperAdminCore) {
	ht.CreatedAt = superAdmin.CreatedAt
	ht.Email = superAdmin.Email
	ht.Password = superAdmin.Password
	ht.Role = uint(superAdmin.Role)
	ht.Nickname = superAdmin.Nickname
	ht.Firstname = superAdmin.Firstname
	ht.Lastname = superAdmin.Lastname
	ht.Middlename = superAdmin.Middlename
}
