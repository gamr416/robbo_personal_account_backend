package models

import (
	"gorm.io/gorm"
	"strconv"
)

type CourseRelationsCore struct {
	Id           string
	LastModified string
	//параметр строка, который позволяет понять с чем связан курс (параметры: group, unit, role)
	Parameter string
	CourseId  string
	//id того с чем связь
	ObjectId string
}

type CourseRelationsDB struct {
	gorm.Model

	Parameter string
	CourseId  string
	Course    CourseDB `gorm:"foreignKey:CourseId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ObjectId  string
}

func (em *CourseRelationsDB) ToCore() *CourseRelationsCore {
	return &CourseRelationsCore{
		Id:           strconv.FormatUint(uint64(em.ID), 10),
		LastModified: em.UpdatedAt.String(),
		Parameter:    em.Parameter,
		CourseId:     em.CourseId,
		ObjectId:     em.ObjectId,
	}
}

func (em *CourseRelationsDB) FromCore(core *CourseRelationsCore) {
	id, _ := strconv.ParseUint(core.Id, 10, 64)
	em.ID = uint(id)
	em.Parameter = core.Parameter
	em.CourseId = core.CourseId
	em.ObjectId = core.ObjectId
}
