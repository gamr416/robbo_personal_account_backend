package models

import (
	"gorm.io/gorm"
	"strconv"
)

type StudentCore struct {
	UserCore
	RobboGroupId string
	RobboUnitId  string
}

type StudentHTTP struct {
	UserHttp     `json:"userHttp"`
	RobboGroupId string `json:"robboGroupId"`
	RobboUnitId  string `json:"robboUnitId"`
}

type StudentDB struct {
	gorm.Model
	UserDB
	RobboGroupId uint
	//RobboGroup   RobboGroupDB `gorm:"foreignKey:RobboGroupId;references:ID;"`
	RobboUnitId uint
	//RobboUnit    RobboUnitDB `gorm:"foreignKey:RobboUnitId;references:ID;"`
}

func (em *StudentDB) ToCore() *StudentCore {
	return &StudentCore{
		UserCore: UserCore{
			Id:         strconv.FormatUint(uint64(em.ID), 10),
			Email:      em.Email,
			Password:   em.Password,
			Role:       Role(em.Role),
			Nickname:   em.Nickname,
			Firstname:  em.Firstname,
			Lastname:   em.Lastname,
			Middlename: em.Middlename,
			CreatedAt:  em.CreatedAt.String(),
		},
		RobboGroupId: strconv.FormatUint(uint64(em.RobboGroupId), 10),
		RobboUnitId:  strconv.FormatUint(uint64(em.RobboUnitId), 10),
	}
}

func (em *StudentDB) FromCore(student *StudentCore) {
	id, _ := strconv.ParseUint(student.Id, 10, 64)
	robboUnitId, _ := strconv.ParseUint(student.RobboUnitId, 10, 64)
	robboGroupId, _ := strconv.ParseUint(student.RobboGroupId, 10, 64)
	em.ID = uint(id)
	em.Email = student.Email
	em.Password = student.Password
	em.Role = uint(student.Role)
	em.Nickname = student.Nickname
	em.Firstname = student.Firstname
	em.Lastname = student.Lastname
	em.Middlename = student.Middlename
	em.RobboGroupId = uint(robboUnitId)
	em.RobboUnitId = uint(robboGroupId)
}

func (ht *StudentHTTP) ToCore() *StudentCore {
	return &StudentCore{
		UserCore: UserCore{
			Id:         ht.Id,
			Email:      ht.Email,
			Password:   ht.Password,
			Role:       Role(ht.Role),
			Nickname:   ht.Nickname,
			Firstname:  ht.Firstname,
			Lastname:   ht.Lastname,
			Middlename: ht.Middlename,
		},
		RobboGroupId: ht.RobboGroupId,
		RobboUnitId:  ht.RobboUnitId,
	}
}

func (ht *StudentHTTP) FromCore(student *StudentCore) {
	ht.Id = student.Id
	ht.CreatedAt = student.CreatedAt
	ht.Email = student.Email
	ht.Password = student.Password
	ht.Role = uint(student.Role)
	ht.Nickname = student.Nickname
	ht.Firstname = student.Firstname
	ht.Lastname = student.Lastname
	ht.Middlename = student.Middlename
	ht.RobboGroupId = student.RobboGroupId
	ht.RobboUnitId = student.RobboUnitId
}
