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

type ProjectPageCore struct {
	ProjectsCore []*ProjectCore
	Description  string
	Preview      string
	LinkScratch  string
}

type ProjectDB struct {
	gorm.Model

	Name   string `gorm:"not null;size:256"`
	Author string `gorm:"not null;size:256"`
	Date   string `gorm:"not null;size:256"`
	Json   string `gorm:"not null;size:65535"`
}

type ProgectPageDB struct {
	gorm.Model

	PPId        uint
	PP          ProjectDB `gorm:"foreignKey:PPId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Description string    `gorm:"size:256;not null"`
	Preview     string    `gorm:"size:256;not null"`
	LinkScratch string    `gorm:"size:256;not null"`
}

type ProjectHTTP struct {
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Author string    `json:"author"`
	Date   time.Time `json:"date"`
	Json   string    `json:"json"`
}

type ProjectPageHTTP struct {
	ProjectsHTTP []*ProjectHTTP `json:"projects"`
	Description  string		`json:"description"`
	Preview      string		`json:"preview"`
	LinkScratch  string		`json:"link"`
}


func (em *ProjectDB) ToCore() *ProjectCore {
	return &ProjectCore{
		ID:     strconv.FormatUint(uint64(em.ID), 10),
		Name:   em.Name,
		Author: em.Author,
		Date:   em.Date,
		Json:   em.Json,
	}
}

func (em *ProjectDB) FromCore(project *ProjectCore) {
	id, _ := strconv.ParseUint(project.ID, 10, 64)
	em.ID = uint(id)
	em.Name = project.Name
	em.Author = project.Author
	em.Date = project.Date
	em.Json = project.Json
}

func (em *ProjectPageDB) ToCore(projects []*ProjectDB) ProjectPageCore {
	var coreProjects []*ProjectCore
	for _, projectDB := range projects {
		coreProjects = append(coreProjects, projectDB.ToCore())
	}
	return ProjectPageCore{
		Description: em.Description,
		Preview:     em.Preview,
		LinkScratch: em.LinkScratch,
		ProjectsCore:    coreProjects,
	}
}

func (em *ProjectPageDB) FromCore(pp *ProjectPageCore) {
	em.Description = pp.Description
	em.Preview = pp.Preview
	em.LinkScratch = pp.LinkScratch
}
