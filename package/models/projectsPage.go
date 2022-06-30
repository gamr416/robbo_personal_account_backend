package models

import (
	"gorm.io/gorm"
)

type ProjectPageCore struct {
	LastModified string
	ProjectsCore ProjectCore
	Information  string
	Notes        string
	Preview      string
	LinkScratch  string
	IsShares     bool
}

type ProjectPageDB struct {
	gorm.Model

	PPId        uint
	PP          ProjectDB `gorm:"foreignKey:PPId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Information string    `gorm:"size:256;not null"`
	Notes       string    `gorm:"size:256;not null"`
	Preview     string    `gorm:"size:256;not null"`
	LinkScratch string    `gorm:"size:256;not null"`
	IsShares    bool
}

type ProjectPageHTTP struct {
	LastModified string      `json:"last_modified"`
	ProjectsHTTP ProjectHTTP `json:"projects"`
	Information  string      `json:"information"`
	Notes        string      `json:"notes"`
	Preview      string      `json:"preview"`
	LinkScratch  string      `json:"link"`
	IsShares     bool        `json:"isShares"`
}

func (em *ProjectPageDB) ToCore() ProjectPageCore {
	var coreProjects ProjectCore
	return ProjectPageCore{
		LastModified: em.UpdatedAt.String(),
		Information:  em.Information,
		Notes:        em.Notes,
		Preview:      em.Preview,
		LinkScratch:  em.LinkScratch,
		ProjectsCore: coreProjects,
		IsShares:     em.IsShares,
	}
}

func (em *ProjectPageDB) FromCore(pp *ProjectPageCore) {
	em.Information = pp.Information
	em.Notes = pp.Notes
	em.Preview = pp.Preview
	em.LinkScratch = pp.LinkScratch
	em.IsShares = pp.IsShares
}

func (ht *ProjectPageHTTP) ToCore() ProjectPageCore {
	var coreProjects ProjectCore
	return ProjectPageCore{
		LastModified: ht.LastModified,
		Information:  ht.Information,
		Notes:        ht.Notes,
		Preview:      ht.Preview,
		LinkScratch:  ht.LinkScratch,
		ProjectsCore: coreProjects,
		IsShares:     ht.IsShares,
	}
}

func (ht *ProjectPageHTTP) FromCore(pp *ProjectPageCore) {
	ht.LastModified = pp.LastModified
	ht.Information = pp.Information
	ht.Notes = pp.Notes
	ht.Preview = pp.Preview
	ht.LinkScratch = pp.LinkScratch
	ht.IsShares = pp.IsShares
}
