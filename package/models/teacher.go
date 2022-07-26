package models

import (
	"gorm.io/gorm"
)

type TeacherCore struct {
	UserCore
	//TeachersID []uint
	//CoursesID  []uint
	//GroupsID   []uint
}

type TeacherDB struct {
	gorm.Model
	UserDB
}

type TeacherHTTP struct {
	UserHttp
	//TeachersID []uint `json:"teachers_id"`
	//CoursesID  []uint `json:"courses_id"`
	//GroupsID   []uint `json:"groups_id"`
}

//type TeacherGroup struct {
//	TeacherID uint
//	Teacher   TeacherDB `gorm:"foreignKey:TeacherID;references:ID"`
//	GroupID   uint
//	Group     GroupDB `gorm:"foreignKey:GroupID;references:ID"`
//}
//
//type TeacherCourse struct {
//	TeacherID uint
//	Teacher   TeacherDB `gorm:"foreignKey:TeacherID;references:ID"`
//	CourseID  uint
//	Course    CourseDB `gorm:"foreignKey:CourseID;references:ID"`
//}
//
//type GroupDB struct {
//	gorm.Model
//	Name string
//}
//
//type GroupHTTP struct {
//	ID   string `json:"id"`
//	Name string `json:"name"`
//}
//
//type GroupCore struct {
//	ID   string
//	Name string
//}

func (em *TeacherDB) ToCore() *TeacherCore {
	return &TeacherCore{
		UserCore{
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

func (em *TeacherDB) FromCore(teacher *TeacherCore) {
	em.Email = teacher.Email
	em.Password = teacher.Password
	em.Role = uint(teacher.Role)
	em.Nickname = teacher.Nickname
	em.Firstname = teacher.Firstname
	em.Lastname = teacher.Lastname
	em.Middlename = teacher.Middlename
}

func (ht *TeacherHTTP) ToCore() *TeacherCore {
	return &TeacherCore{
		UserCore{
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

func (ht *TeacherHTTP) FromCore(teacher *TeacherCore) {
	ht.CreatedAt = teacher.CreatedAt
	ht.Email = teacher.Email
	ht.Password = teacher.Password
	ht.Role = uint(teacher.Role)
	ht.Nickname = teacher.Nickname
	ht.Firstname = teacher.Firstname
	ht.Lastname = teacher.Lastname
	ht.Middlename = teacher.Middlename
}
