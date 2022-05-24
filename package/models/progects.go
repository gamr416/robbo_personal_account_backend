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
