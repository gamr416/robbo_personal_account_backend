package models

import (
	"gorm.io/gorm"
	"strconv"
)

type RobboGroupCore struct {
	Id           string
	LastModified string
	Name         string
	RobboUnitId  string
	Students     []*StudentCore
}

type RobboGroupDB struct {
	gorm.Model
	RobboUnitId string `gorm:"not null"`
	Name        string `gorm:"size:256;not null"`
}
type RobboGroupHttp struct {
	Id           string         `json:"id"`
	LastModified string         `json:"lastModified"`
	Name         string         `json:"name"`
	RobboUnitId  string         `json:"robboUnitId"`
	Students     []*StudentHTTP `json:"students"`
}

func (em *RobboGroupDB) ToCore() *RobboGroupCore {
	return &RobboGroupCore{
		Id:           strconv.FormatUint(uint64(em.ID), 10),
		LastModified: em.UpdatedAt.String(),
		Name:         em.Name,
		RobboUnitId:  em.RobboUnitId,
	}
}

func (em *RobboGroupDB) FromCore(robboGroup *RobboGroupCore) {
	id, _ := strconv.ParseUint(robboGroup.Id, 10, 64)
	em.ID = uint(id)
	em.Name = robboGroup.Name
	em.RobboUnitId = robboGroup.RobboUnitId
}

func (ht *RobboGroupHttp) ToCore() *RobboGroupCore {
	var studentsCore []*StudentCore
	for _, studentHttp := range ht.Students {
		studentsCore = append(studentsCore, studentHttp.ToCore())
	}
	return &RobboGroupCore{
		Id:           ht.Id,
		LastModified: ht.LastModified,
		RobboUnitId:  ht.RobboUnitId,
		Name:         ht.Name,
		Students:     studentsCore,
	}
}

func (ht *RobboGroupHttp) FromCore(robboGroup *RobboGroupCore) {
	ht.Id = robboGroup.Id
	ht.LastModified = robboGroup.LastModified
	ht.Name = robboGroup.Name
	ht.RobboUnitId = robboGroup.RobboUnitId
	for _, studentCore := range robboGroup.Students {
		var studentHttpTemp StudentHTTP
		studentHttpTemp.FromCore(studentCore)
		ht.Students = append(ht.Students, &studentHttpTemp)
	}

}
