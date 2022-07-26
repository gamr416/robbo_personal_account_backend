package models

import (
	"gorm.io/gorm"
	"strconv"
)

type ParentCore struct {
	ID string
	User
	StudentsID []uint
}

type ParentDB struct {
	gorm.Model
	User
}

type ParentHTTP struct {
	ID string `json:"id"`
	User
	StudentsID []uint `json:"students_id"`
}

type ParentStudent struct {
	ParentID  uint
	Parent    ParentDB `gorm:"foreignKey:ParentID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
	StudentID uint
	Student   StudentDB `gorm:"foreignKey:StudentID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
}

func (em *ParentDB) ToCore() *ParentCore {
	return &ParentCore{
		ID:   strconv.FormatUint(uint64(em.ID), 10),
		User: em.User,
	}
}

func (em *ParentDB) FromCore(parent *ParentCore) {
	id, _ := strconv.ParseUint(parent.ID, 10, 64)
	em.ID = uint(id)
	em.User = parent.User
}

func (ht *ParentHTTP) ToCore() *ParentCore {
	return &ParentCore{
		ID:         ht.ID,
		User:       ht.User,
		StudentsID: ht.StudentsID,
	}
}

func (ht *ParentHTTP) FromCore(parent *ParentCore) {
	ht.ID = parent.ID
	ht.User = parent.User
	ht.StudentsID = parent.StudentsID
}
