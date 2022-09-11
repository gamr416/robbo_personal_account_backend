package models

import (
	"gorm.io/gorm"
	"strconv"
)

type FreeListenerCore struct {
	UserCore
}

type FreeListenerDB struct {
	gorm.Model
	UserDB
}

type FreeListenerHttp struct {
	UserHTTP `json:"userHttp"`
}

func (em *FreeListenerDB) ToCore() *FreeListenerCore {
	return &FreeListenerCore{
		UserCore: em.UserDB.ToCore(),
	}
}

func (em *FreeListenerDB) FromCore(freeListener *FreeListenerCore) {
	id, _ := strconv.ParseUint(freeListener.Id, 10, 64)
	em.ID = uint(id)
	em.UserDB.ToCore()
}

func (ht *FreeListenerHttp) ToCore() *FreeListenerCore {
	return &FreeListenerCore{
		UserCore: ht.UserHTTP.ToCore(),
	}
}

func (ht *FreeListenerHttp) FromCore(freeLister *FreeListenerCore) {
	ht.UserHTTP.FromCore(&freeLister.UserCore)
}
