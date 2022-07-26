package models

import (
	"gorm.io/gorm"
	"strconv"
)

type UnitAdminHTTP struct {
	ID string `json:"id"`
	User
}

type UnitAdminDB struct {
	gorm.Model
	User
}

type UnitAdminCore struct {
	ID string
	User
}

func (em *UnitAdminDB) ToCore() *UnitAdminCore {
	return &UnitAdminCore{
		ID:   strconv.FormatUint(uint64(em.ID), 10),
		User: em.User,
	}
}

func (em *UnitAdminDB) FromCore(unitAdmin *UnitAdminCore) {
	id, _ := strconv.ParseUint(unitAdmin.ID, 10, 64)
	em.ID = uint(id)
	em.User = unitAdmin.User
}

func (ht *UnitAdminHTTP) ToCore() *UnitAdminCore {
	return &UnitAdminCore{
		ID:   ht.ID,
		User: ht.User,
	}
}

func (ht *UnitAdminHTTP) FromCore(unitAdmin *UnitAdminCore) {
	ht.ID = unitAdmin.ID
	ht.User = unitAdmin.User
}
