package models

import (
	"gorm.io/gorm"
	"strconv"
)

type StudentCore struct {
	ID string
	User
	ParentID   uint
	GroupsID   []uint
	TeachersID []uint
	ProjectsID []uint
}

type StudentHTTP struct {
	ID string `json:"id"`
	User
	ParentID   uint   `json:"parent_id"`
	GroupsID   []uint `json:"groups_id"`
	TeachersID []uint `json:"teachers_id"`
	ProjectsID []uint `json:"projects_id"`
}

type StudentDB struct {
	gorm.Model
	User
	//ParentID uint
	//Parent   ParentDB `gorm:"foreignKey:ParentID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
}

type StudentGroup struct {
	StudentID uint
	Student   StudentDB `gorm:"foreignKey:StudentID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
	GroupID   uint
	Group     GroupDB `gorm:"foreignKey:GroupID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
}

type StudentTeacher struct {
	StudentID uint
	Student   StudentDB `gorm:"foreignKey:StudentID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
	TeacherID uint
	Teacher   TeacherDB `gorm:"foreignKey:TeacherID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
}

type StudentProject struct {
	StudentID uint
	Student   StudentDB `gorm:"foreignKey:StudentID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
	ProjectID uint
	Project   ProjectDB `gorm:"foreignKey:ProjectID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
}

func (em *StudentDB) ToCore() *StudentCore {
	return &StudentCore{
		ID:   strconv.FormatUint(uint64(em.ID), 10),
		User: em.User,
		//ParentID: em.ParentID,
	}
}

func (em *StudentDB) FromCore(student *StudentCore) {
	id, _ := strconv.ParseUint(student.ID, 10, 64)
	em.ID = uint(id)
	em.User = student.User
	//	em.ParentID = student.ParentID
}

func (ht *StudentHTTP) ToCore() *StudentCore {
	return &StudentCore{
		ID:         ht.ID,
		User:       ht.User,
		ParentID:   ht.ParentID,
		GroupsID:   ht.GroupsID,
		ProjectsID: ht.ProjectsID,
		TeachersID: ht.TeachersID,
	}
}

func (ht *StudentHTTP) FromCore(student *StudentCore) {
	ht.ID = student.ID
	ht.User = student.User
	ht.ParentID = student.ParentID
	ht.GroupsID = student.GroupsID
	ht.ProjectsID = student.ProjectsID
	ht.TeachersID = student.TeachersID
}
