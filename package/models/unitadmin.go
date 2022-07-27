package models

import (
	"gorm.io/gorm"
	"strconv"
)

type UnitAdminHTTP struct {
	UserHttp
}

type UnitAdminDB struct {
	gorm.Model
	UserDB
}

type UnitAdminCore struct {
	UserCore
}

func (em *UnitAdminDB) ToCore() *UnitAdminCore {
	return &UnitAdminCore{
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

func (em *UnitAdminDB) FromCore(unitAdmin *UnitAdminCore) {
	id, _ := strconv.ParseUint(unitAdmin.Id, 10, 64)
	em.ID = uint(id)
	em.Email = unitAdmin.Email
	em.Password = unitAdmin.Password
	em.Role = uint(unitAdmin.Role)
	em.Nickname = unitAdmin.Nickname
	em.Firstname = unitAdmin.Firstname
	em.Lastname = unitAdmin.Lastname
	em.Middlename = unitAdmin.Middlename
}

func (ht *UnitAdminHTTP) ToCore() *UnitAdminCore {
	return &UnitAdminCore{
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

func (ht *UnitAdminHTTP) FromCore(unitAdmin *UnitAdminCore) {
	ht.Id = unitAdmin.Id
	ht.CreatedAt = unitAdmin.CreatedAt
	ht.Email = unitAdmin.Email
	ht.Password = unitAdmin.Password
	ht.Role = uint(unitAdmin.Role)
	ht.Nickname = unitAdmin.Nickname
	ht.Firstname = unitAdmin.Firstname
	ht.Lastname = unitAdmin.Lastname
	ht.Middlename = unitAdmin.Middlename
}
