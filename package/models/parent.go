package models

import (
	"gorm.io/gorm"
)

type ParentCore struct {
	UserCore
	//StudentsID []uint
}

type ParentDB struct {
	gorm.Model
	UserDB
}

type ParentHTTP struct {
	UserHttp
	//StudentsID []uint `json:"students_id"`
}

//type ParentStudent struct {
//	ParentID  uint
//	Parent    ParentDB `gorm:"foreignKey:ParentID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
//	StudentID uint
//	Student   StudentDB `gorm:"foreignKey:StudentID;references:ID;constraint:onUpdate:CASCADE;onDELETE:CASCADE"`
//}

func (em *ParentDB) ToCore() *ParentCore {
	return &ParentCore{
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
		//StudentsID: students,
	}
}

func (em *ParentDB) FromCore(parent *ParentCore) {
	em.Email = parent.Email
	em.Password = parent.Password
	em.Role = uint(parent.Role)
	em.Nickname = parent.Nickname
	em.Firstname = parent.Firstname
	em.Lastname = parent.Lastname
	em.Middlename = parent.Middlename
}

func (ht *ParentHTTP) ToCore() *ParentCore {
	return &ParentCore{
		UserCore: UserCore{
			Email:      ht.Email,
			Password:   ht.Password,
			Role:       Role(ht.Role),
			Nickname:   ht.Nickname,
			Firstname:  ht.Firstname,
			Lastname:   ht.Lastname,
			Middlename: ht.Middlename,
		},
		//StudentsID: ht.StudentsID,
	}
}

func (ht *ParentHTTP) FromCore(parent *ParentCore) {
	ht.CreatedAt = parent.CreatedAt
	ht.Email = parent.Email
	ht.Password = parent.Password
	ht.Role = uint(parent.Role)
	ht.Nickname = parent.Nickname
	ht.Firstname = parent.Firstname
	ht.Lastname = parent.Lastname
	ht.Middlename = parent.Middlename
	//ht.StudentsID = parent.StudentsID
}
