package models

import (
	"gorm.io/gorm"
)

type StudentCore struct {
	UserCore
	ParentID uint
	//GroupsID   []uint
	//TeachersID []uint
	//ProjectsID []uint
}

type StudentHTTP struct {
	UserHttp
	ParentID uint `json:"parent_id"`
	//GroupsID   []uint `json:"groups_id"`
	//TeachersID []uint `json:"teachers_id"`
	//ProjectsID []uint `json:"projects_id"`
}

type StudentDB struct {
	gorm.Model
	UserDB
	ParentID uint
	Parent   ParentDB `gorm:"foreignKey:ParentID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
}

//type StudentGroup struct {
//	StudentID uint
//	Student   StudentDB `gorm:"foreignKey:StudentID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
//	GroupID   uint
//	Group     GroupDB `gorm:"foreignKey:GroupID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
//}
//
//type StudentTeacher struct {
//	StudentID uint
//	Student   StudentDB `gorm:"foreignKey:StudentID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
//	TeacherID uint
//	Teacher   TeacherDB `gorm:"foreignKey:TeacherID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
//}
//
//type StudentProject struct {
//	StudentID uint
//	Student   StudentDB `gorm:"foreignKey:StudentID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
//	ProjectID uint
//	Project   ProjectDB `gorm:"foreignKey:ProjectID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
//}

func (em *StudentDB) ToCore() *StudentCore {
	return &StudentCore{
		ParentID: em.ParentID,
		UserCore: UserCore{
			Email:      em.Email,
			Password:   em.Password,
			Role:       Role(em.Role),
			Nickname:   em.Nickname,
			Firstname:  em.Firstname,
			Lastname:   em.Lastname,
			Middlename: em.Middlename,
			CreatedAt:  em.CreatedAt.String(),
		},
	}
}

func (em *StudentDB) FromCore(student *StudentCore) {
	em.Email = student.Email
	em.Password = student.Password
	em.Role = uint(student.Role)
	em.Nickname = student.Nickname
	em.Firstname = student.Firstname
	em.Lastname = student.Lastname
	em.Middlename = student.Middlename
	em.ParentID = student.ParentID
}

func (ht *StudentHTTP) ToCore() *StudentCore {
	return &StudentCore{
		ParentID: ht.ParentID,
		UserCore: UserCore{
			Email:      ht.Email,
			Password:   ht.Password,
			Role:       Role(ht.Role),
			Nickname:   ht.Nickname,
			Firstname:  ht.Firstname,
			Lastname:   ht.Lastname,
			Middlename: ht.Middlename,
		},
	}
}

func (ht *StudentHTTP) FromCore(student *StudentCore) {
	ht.CreatedAt = student.CreatedAt
	ht.Email = student.Email
	ht.Password = student.Password
	ht.Role = uint(student.Role)
	ht.Nickname = student.Nickname
	ht.Firstname = student.Firstname
	ht.Lastname = student.Lastname
	ht.Middlename = student.Middlename
	ht.ParentID = student.ParentID
}
