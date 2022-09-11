package models

import (
	"gorm.io/gorm"
	"strconv"
)

type UnitAdminHTTP struct {
	UserHTTP `json:"userHttp"`
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
		UserCore: em.UserDB.ToCore(),
	}
}

func (em *UnitAdminDB) FromCore(unitAdmin *UnitAdminCore) {
	id, _ := strconv.ParseUint(unitAdmin.Id, 10, 64)
	em.ID = uint(id)
	em.UserDB.FromCore(&unitAdmin.UserCore)
}

func (ht *UnitAdminHTTP) ToCore() *UnitAdminCore {
	return &UnitAdminCore{
		UserCore: ht.UserHTTP.ToCore(),
	}
}

func (ht *UnitAdminHTTP) FromCore(unitAdmin *UnitAdminCore) {
	ht.UserHTTP.FromCore(&unitAdmin.UserCore)
}
