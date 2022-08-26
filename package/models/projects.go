package models

import (
	"gorm.io/gorm"
	"strconv"
)

type ProjectCore struct {
	ID       string
	Name     string
	AuthorId string
	Json     string
}

type ProjectDB struct {
	gorm.Model

	Name     string `gorm:"not null;size:256"`
	AuthorId string `gorm:"not null;size:256"`
	Json     string `gorm:"not null;size:65535"`
}

type ProjectHTTP struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	AuthorId string `json:"authorId"`
	Json     string `json:"json"`
}

func (em *ProjectDB) ToCore() *ProjectCore {
	return &ProjectCore{
		ID:       strconv.FormatUint(uint64(em.ID), 10),
		Name:     em.Name,
		AuthorId: em.AuthorId,
		Json:     em.Json,
	}
}

func (em *ProjectDB) FromCore(project *ProjectCore) {
	id, _ := strconv.ParseUint(project.ID, 10, 64)
	em.ID = uint(id)
	em.Name = project.Name
	em.AuthorId = project.AuthorId
	em.Json = project.Json
}

func (ht *ProjectHTTP) ToCore() *ProjectCore {
	return &ProjectCore{
		ID:       ht.ID,
		Name:     ht.Name,
		AuthorId: ht.AuthorId,
		Json:     ht.Json,
	}
}

func (ht *ProjectHTTP) FromCore(project *ProjectCore) {
	ht.ID = project.ID
	ht.Name = project.Name
	ht.AuthorId = project.AuthorId
	ht.Json = project.Json
}
