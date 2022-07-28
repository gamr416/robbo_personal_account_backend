package models

import (
	"gorm.io/gorm"
	"strconv"
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
			Id:         strconv.FormatUint(uint64(em.ID), 10),
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
	id, _ := strconv.ParseUint(superAdmin.Id, 10, 64)
	em.ID = uint(id)
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
			Id:         ht.Id,
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
	ht.Id = superAdmin.Id
	ht.CreatedAt = superAdmin.CreatedAt
	ht.Email = superAdmin.Email
	ht.Password = superAdmin.Password
	ht.Role = uint(superAdmin.Role)
	ht.Nickname = superAdmin.Nickname
	ht.Firstname = superAdmin.Firstname
	ht.Lastname = superAdmin.Lastname
	ht.Middlename = superAdmin.Middlename
}
