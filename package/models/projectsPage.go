package models

import (
	"gorm.io/gorm"
)

type ProjectPageCore struct {
	LastModified string
	Title        string
	ProjectId    string
	Instruction  string
	Notes        string
	Preview      string
	LinkScratch  string
	IsShared     bool
}

type ProjectPageDB struct {
	gorm.Model

	ProjectId   string
	Project     ProjectDB `gorm:"foreignKey:ProjectId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Instruction string    `gorm:"size:256;not null"`
	Notes       string    `gorm:"size:256;not null"`
	Preview     string    `gorm:"size:256;not null"`
	LinkScratch string    `gorm:"size:256;not null"`
	Title       string    `gorm:"size:256;not null"`
	IsShared    bool
}

func (em *ProjectPageDB) ToCore() *ProjectPageCore {
	return &ProjectPageCore{
		LastModified: em.UpdatedAt.String(),
		Title:        em.Title,
		ProjectId:    em.ProjectId,
		Instruction:  em.Instruction,
		Notes:        em.Notes,
		Preview:      em.Preview,
		LinkScratch:  em.LinkScratch,
		IsShared:     em.IsShared,
	}
}

func (em *ProjectPageDB) FromCore(pp *ProjectPageCore) {
	em.ProjectId = pp.ProjectId
	em.Instruction = pp.Instruction
	em.Notes = pp.Notes
	em.Preview = pp.Preview
	em.LinkScratch = pp.LinkScratch
	em.Title = pp.Title
	em.IsShared = pp.IsShared
}

func (ht *ProjectPageHTTP) ToCore() *ProjectPageCore {
	return &ProjectPageCore{
		LastModified: ht.LastModified,
		Title:        ht.Title,
		ProjectId:    ht.ProjectID,
		Instruction:  ht.Instruction,
		Notes:        ht.Notes,
		Preview:      ht.Preview,
		LinkScratch:  ht.LinkScratch,
		IsShared:     ht.IsShared,
	}
}

func (ht *ProjectPageHTTP) FromCore(pp *ProjectPageCore) {
	ht.LastModified = pp.LastModified
	ht.ProjectID = pp.ProjectId
	ht.Instruction = pp.Instruction
	ht.Notes = pp.Notes
	ht.Preview = pp.Preview
	ht.LinkScratch = pp.LinkScratch
	ht.Title = pp.Title
	ht.IsShared = pp.IsShared
}
