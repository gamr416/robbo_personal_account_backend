package models

import (
	"gorm.io/gorm"
	"time"
)

type ProjectCore struct {
	ID     string
	Name   string
	Author string
	Date   time.Time
	Json   string
}

type ProjectDB struct {
	gorm.Model

	Name   string `gorm:"not null;size:256"`
	Author string `gorm:"not null;size:256"`
	Date   string `gorm:"not null;size:256"`
	Json   string `gorm:"not null;size:65535"`
}

type ProjectHTTP struct {
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Author string    `json:"author"`
	Date   time.Time `json:"date"`
	Json   string    `json:"json"`
}

//func (em *ProjectDB) ToCore() *ProjectCore {
//	return &ProjectCore{
//		ID:     strconv.FormatUint(uint64(em.ID), 10),
//		Name:   em.Name,
//		Author: em.Author,
//		Date:   em.Date,
//		Json:   em.Json,
//	}
//}
//
//func (em *ProjectDB) FromCore(project *ProjectCore) {
//	id, _ := strconv.ParseUint(project.ID, 10, 64)
//	em.ID = uint(id)
//	em.Name = project.Name
//	em.Author = project.Author
//	em.Date = project.Date
//	em.Json = project.Json
//}

// TODO mapping Http
