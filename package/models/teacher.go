package models

import (
	"gorm.io/gorm"
	"strconv"
)

type TeacherCore struct {
	ID string
	User
	TeachersID []uint
	CoursesID  []uint
	GroupsID   []uint
}

type TeacherDB struct {
	gorm.Model
	User
}

type TeacherHTTP struct {
	ID string `json:"id"`
	User
	TeachersID []uint `json:"teachers_id"`
	CoursesID  []uint `json:"courses_id"`
	GroupsID   []uint `json:"groups_id"`
}

type TeacherGroup struct {
	TeacherID uint
	Teacher   TeacherDB `gorm:"foreignKey:TeacherID;references:ID"`
	GroupID   uint
	Group     GroupDB `gorm:"foreignKey:GroupID;references:ID"`
}

type TeacherCourse struct {
	TeacherID uint
	Teacher   TeacherDB `gorm:"foreignKey:TeacherID;references:ID"`
	CourseID  uint
	Course    CourseDB `gorm:"foreignKey:CourseID;references:ID"`
}

type GroupDB struct {
	gorm.Model
	Name string
}

type GroupHTTP struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GroupCore struct {
	ID   string
	Name string
}

func (em *TeacherDB) ToCore() *TeacherCore {
	return &TeacherCore{
		ID:   strconv.FormatUint(uint64(em.ID), 10),
		User: em.User,
	}
}

func (em *TeacherDB) FromCore(teacher *TeacherCore) {
	id, _ := strconv.ParseUint(teacher.ID, 10, 64)
	em.ID = uint(id)
	em.User = teacher.User
}

func (ht *TeacherHTTP) ToCore() *TeacherCore {
	return &TeacherCore{
		ID:         ht.ID,
		User:       ht.User,
		TeachersID: ht.TeachersID,
		CoursesID:  ht.CoursesID,
	}
}

func (ht *TeacherHTTP) FromCore(teacher *TeacherCore) {
	ht.ID = teacher.ID
	ht.User = teacher.User
	ht.TeachersID = teacher.TeachersID
	ht.CoursesID = teacher.CoursesID
}
