package models

import "gorm.io/gorm"

/*
	вспомогательная структура
   	для хранения связи между teacher и robboGroup
*/

type TeachersRobboGroupsCore struct {
	TeacherId    string
	RobboGroupId string
}

type TeachersRobboGroupsDB struct {
	gorm.Model

	TeacherId    string
	Teacher      TeacherDB `gorm:"foreignKey:TeacherId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RobboGroupId string
	RobboGroup   RobboGroupDB `gorm:"foreignKey:RobboGroupId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (em *TeachersRobboGroupsDB) FromCore(core *TeachersRobboGroupsCore) {
	em.TeacherId = core.TeacherId
	em.RobboGroupId = core.RobboGroupId
}

func (em *TeachersRobboGroupsDB) ToCore() *TeachersRobboGroupsCore {
	return &TeachersRobboGroupsCore{
		TeacherId:    em.TeacherId,
		RobboGroupId: em.RobboGroupId,
	}
}
