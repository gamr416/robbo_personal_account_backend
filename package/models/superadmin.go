package models

import (
	"gorm.io/gorm"
	"strconv"
)

type SuperAdminHTTP struct {
	ID string `json:"id"`
	User
}

type SuperAdminDB struct {
	gorm.Model
	User
}

type SuperAdminCore struct {
	ID string
	User
}

func (em *SuperAdminDB) ToCore() *SuperAdminCore {
	return &SuperAdminCore{
		ID:   strconv.FormatUint(uint64(em.ID), 10),
		User: em.User,
	}
}

func (em *SuperAdminDB) FromCore(superAdmin *SuperAdminCore) {
	id, _ := strconv.ParseUint(superAdmin.ID, 10, 64)
	em.ID = uint(id)
	em.User = superAdmin.User
}

func (ht *SuperAdminHTTP) ToCore() *SuperAdminCore {
	return &SuperAdminCore{
		ID:   ht.ID,
		User: ht.User,
	}
}

func (ht *SuperAdminHTTP) FromCore(superAdmin *SuperAdminCore) {
	ht.ID = superAdmin.ID
	ht.User = superAdmin.User
}
