package models

import (
	"gorm.io/gorm"
	"strconv"
)

type ParentCore struct {
	UserCore
	Children []*StudentCore
}

type ParentDB struct {
	gorm.Model
	UserDB
}

type ParentHTTP struct {
	UserHttp `json:"userHttp"`
	Children []*StudentHTTP `json:"children"`
}

func (em *ParentDB) ToCore() *ParentCore {
	return &ParentCore{
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
	}
}

func (em *ParentDB) FromCore(parent *ParentCore) {
	id, _ := strconv.ParseUint(parent.Id, 10, 64)
	em.ID = uint(id)
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
			Id:         ht.Id,
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
	// TODO refactor user.ToCore()
	ht.Id = parent.Id
	ht.CreatedAt = parent.CreatedAt
	ht.Email = parent.Email
	ht.Password = parent.Password
	ht.Role = uint(parent.Role)
	ht.Nickname = parent.Nickname
	ht.Firstname = parent.Firstname
	ht.Lastname = parent.Lastname
	ht.Middlename = parent.Middlename
	for _, child := range parent.Children {
		var childTemp StudentHTTP
		childTemp.FromCore(child)
		ht.Children = append(ht.Children, &childTemp)
	}
}
