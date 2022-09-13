package models

import (
	"gorm.io/gorm"
	"strconv"
)

type SuperAdminDB struct {
	gorm.Model
	UserDB
}

type SuperAdminCore struct {
	UserCore
}

func (em *SuperAdminDB) ToCore() *SuperAdminCore {
	return &SuperAdminCore{
		UserCore: em.UserDB.ToCore(),
	}
}

func (em *SuperAdminDB) FromCore(superAdmin *SuperAdminCore) {
	id, _ := strconv.ParseUint(superAdmin.Id, 10, 64)
	em.ID = uint(id)
	em.UserDB.ToCore()
}

func (ht *SuperAdminHTTP) ToCore() *SuperAdminCore {
	return &SuperAdminCore{
		UserCore: ht.UserHTTP.ToCore(),
	}
}

func (ht *SuperAdminHTTP) FromCore(superAdmin *SuperAdminCore) {
	ht.UserHTTP.FromCore(&superAdmin.UserCore)
}
