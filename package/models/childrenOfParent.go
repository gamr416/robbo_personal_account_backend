package models

import "gorm.io/gorm"

/*
	вспомогательная структура
   	для хранения связи между родителем и ребенком
*/

type ChildrenOfParentCore struct {
	ParentId string
	ChildId  string
}

type ChildrenOfParentDB struct {
	gorm.Model

	ParentId string
	Parent   ParentDB `gorm:"foreignKey:ParentId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ChildId  string
	Children StudentDB `gorm:"foreignKey:ChildId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (em *ChildrenOfParentDB) FromCore(core *ChildrenOfParentCore) {
	em.ParentId = core.ParentId
	em.ChildId = core.ChildId
}

func (em *ChildrenOfParentDB) ToCore() *ChildrenOfParentCore {
	return &ChildrenOfParentCore{
		ParentId: em.ParentId,
		ChildId:  em.ChildId,
	}
}
