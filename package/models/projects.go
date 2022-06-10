package models

import (
	"gorm.io/gorm"
	"strconv"
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
	Json   string `gorm:"not null;size:65535"`
}

type ProjectHTTP struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Json   string `json:"json"`
}

func (em *ProjectDB) ToCore() *ProjectCore {
	return &ProjectCore{
		ID:     strconv.FormatUint(uint64(em.ID), 10),
		Name:   em.Name,
		Author: em.Author,
		Json:   em.Json,
	}
}

func (em *ProjectDB) FromCore(project *ProjectCore) {
	id, _ := strconv.ParseUint(project.ID, 10, 64)
	em.ID = uint(id)
	em.Name = project.Name
	em.Author = project.Author
	em.Json = project.Json
}

// TODO mapping Http
