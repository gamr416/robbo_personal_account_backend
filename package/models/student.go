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

//type StudentHTTP struct {
//	UserHttp     `json:"userHttp"`
//	RobboGroupId string `json:"robboGroupId"`
//	RobboUnitId  string `json:"robboUnitId"`
//}

type StudentDB struct {
	gorm.Model
	UserDB
	RobboGroupId uint
	RobboUnitId  uint
}

func (em *StudentDB) ToCore() *StudentCore {
	return &StudentCore{
		UserCore:     em.UserDB.ToCore(),
		RobboGroupId: strconv.FormatUint(uint64(em.RobboGroupId), 10),
		RobboUnitId:  strconv.FormatUint(uint64(em.RobboUnitId), 10),
	}
}

func (em *StudentDB) FromCore(student *StudentCore) {
	id, _ := strconv.ParseUint(student.Id, 10, 64)
	robboUnitId, _ := strconv.ParseUint(student.RobboUnitId, 10, 64)
	robboGroupId, _ := strconv.ParseUint(student.RobboGroupId, 10, 64)
	em.UserDB.FromCore(&student.UserCore)
	em.ID = uint(id)
	em.RobboGroupId = uint(robboUnitId)
	em.RobboUnitId = uint(robboGroupId)
}

func (ht *StudentHTTP) ToCore() *StudentCore {
	return &StudentCore{
		UserCore:     ht.UserHTTP.ToCore(),
		RobboGroupId: ht.RobboGroupID,
		RobboUnitId:  ht.RobboUnitID,
	}
}

func (ht *StudentHTTP) FromCore(student *StudentCore) {
	ht.UserHTTP.FromCore(&student.UserCore)
	ht.RobboGroupID = student.RobboGroupId
	ht.RobboUnitID = student.RobboUnitId
}
