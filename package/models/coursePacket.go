package models

import (
	"strconv"

	"gorm.io/gorm"
)

// type PostEnrollmentHTTP struct {
// 	Message map[string]interface{} `json:"message"`
// }

type CoursePacketCore struct {
	ID      string
	Name    string
	Level   uint
	Courses []*CourseCore
}

type CoursePacketDB struct {
	gorm.Model

	Name  string `gorm:"size:256"`
	Level uint
}

type CoursePacketHTTP struct {
	ID      string        `json:"id"`
	Name    string        `json:"name"`
	Level   uint          `json:"level"`
	Courses []*CourseHTTP `json:"courses"`
}

func (em *CoursePacketDB) ToCore() *CoursePacketCore {
	return &CoursePacketCore{
		ID:    strconv.FormatUint(uint64(em.ID), 10),
		Name:  em.Name,
		Level: em.Level,
	}
}

func (em *CoursePacketDB) FromCore(course_packet *CoursePacketCore) {
	id, _ := strconv.ParseUint(course_packet.ID, 10, 64)
	em.ID = uint(id)
	em.Name = course_packet.Name
	em.Level = course_packet.Level

}

func (ht *CoursePacketHTTP) FromCore(course_packet *CoursePacketCore) {
	ht.ID = course_packet.ID
	ht.Name = course_packet.Name
	ht.Level = course_packet.Level

}

func (ht *CoursePacketHTTP) ToCore() *CoursePacketCore {

	return &CoursePacketCore{
		ID:    ht.ID,
		Name:  ht.Name,
		Level: ht.Level,
	}
}
