package models

import "gorm.io/gorm"

/*
	вспомогательная структура
   	для хранения связи между unitAdmin и robboUnit
*/

type UnitAdminsRobboUnitsCore struct {
	UnitAdminId string
	RobboUnitId string
}

type UnitAdminsRobboUnitsDB struct {
	gorm.Model

	UnitAdminId string
	UnitAdmin   UnitAdminDB `gorm:"foreignKey:UnitAdminId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RobboUnitId string
	RobboUnit   RobboUnitDB `gorm:"foreignKey:RobboUnitId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (em *UnitAdminsRobboUnitsDB) FromCore(core *UnitAdminsRobboUnitsCore) {
	em.UnitAdminId = core.UnitAdminId
	em.RobboUnitId = core.RobboUnitId
}

func (em *UnitAdminsRobboUnitsDB) ToCore() *UnitAdminsRobboUnitsCore {
	return &UnitAdminsRobboUnitsCore{
		UnitAdminId: em.UnitAdminId,
		RobboUnitId: em.RobboUnitId,
	}
}
