package models

import "gorm.io/gorm"

type ProjectPageCore struct {
	ProjectsCore []*ProjectCore
	Description  string
	Preview      string
	LinkScratch  string
}

type ProjectPageDB struct {
	gorm.Model

	PPId        uint
	PP          ProjectDB `gorm:"foreignKey:PPId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Description string    `gorm:"size:256;not null"`
	Preview     string    `gorm:"size:256;not null"`
	LinkScratch string    `gorm:"size:256;not null"`
}

type ProjectPageHTTP struct {
	ProjectsHTTP []*ProjectHTTP `json:"projects"`
	Description  string         `json:"description"`
	Preview      string         `json:"preview"`
	LinkScratch  string         `json:"link"`
}

func (em *ProjectPageDB) ToCore(projects []*ProjectDB) ProjectPageCore {
	var coreProjects []*ProjectCore
	for _, projectDB := range projects {
		coreProjects = append(coreProjects, projectDB.ToCore())
	}
	return ProjectPageCore{
		Description:  em.Description,
		Preview:      em.Preview,
		LinkScratch:  em.LinkScratch,
		ProjectsCore: coreProjects,
	}
}

func (em *ProjectPageDB) FromCore(pp *ProjectPageCore) {
	em.Description = pp.Description
	em.Preview = pp.Preview
	em.LinkScratch = pp.LinkScratch
}

func (ht *ProjectHTTP) ToCore() *ProjectCore {
	return &ProjectCore{
		ID:     ht.ID,
		Name:   ht.Name,
		Author: ht.Author,
		Json:   ht.Json,
	}
}

func (ht *ProjectHTTP) FromCore(project *ProjectCore) {
	ht.ID = project.ID
	ht.Name = project.Name
	ht.Author = project.Author
	ht.Json = project.Json
}

func (ht *ProjectPageHTTP) ToCore(projects []*ProjectDB) ProjectPageCore {
	var coreProjects []*ProjectCore
	for _, projectDB := range projects {
		coreProjects = append(coreProjects, projectDB.ToCore())
	}
	return ProjectPageCore{
		Description:  ht.Description,
		Preview:      ht.Preview,
		LinkScratch:  ht.LinkScratch,
		ProjectsCore: coreProjects,
	}
}

func (ht *ProjectPageHTTP) FromCore(pp *ProjectPageCore) {
	ht.Description = pp.Description
	ht.Preview = pp.Preview
	ht.LinkScratch = pp.LinkScratch
}
