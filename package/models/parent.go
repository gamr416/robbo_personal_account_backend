package models

import (
	"strconv"
)

type ParentCore struct {
	UserCore
	Children []*StudentCore
}

type ParentDB struct {
	UserDB
}

func (em *ParentDB) ToCore() *ParentCore {
	return &ParentCore{
		UserCore: em.UserDB.ToCore(),
	}
}

func (em *ParentDB) FromCore(parent *ParentCore) {
	id, _ := strconv.ParseUint(parent.Id, 10, 64)
	em.ID = uint(id)
	em.UserDB.FromCore(&parent.UserCore)
}

func (ht *ParentHTTP) ToCore() *ParentCore {
	return &ParentCore{
		UserCore: ht.UserHTTP.ToCore(),
	}
}

func (ht *ParentHTTP) FromCore(parent ParentCore) {
	ht.UserHTTP.FromCore(&parent.UserCore)
	for _, child := range parent.Children {
		var childTemp StudentHTTP
		childTemp.FromCore(child)
		ht.Children = append(ht.Children, &childTemp)
	}
}
