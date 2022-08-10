package models

import (
	"gorm.io/gorm"
)

type RobboUnitCore struct {
	LastModified string
	Name         string
	City         string
}

type RobboUnitDB struct {
	gorm.Model

	Name string `gorm:"size:256;not null"`
	City string `gorm:"size:256;not null"`
}

type RobboUnitHTTP struct {
	LastModified string `json:"lastModified"`
	Name         string `json:"name"`
	City         string `json:"city"`
}

func (em *RobboUnitDB) ToCore() *RobboUnitCore {
	return &RobboUnitCore{
		LastModified: em.UpdatedAt.String(),
		Name:         em.Name,
		City:         em.City,
	}
}

func (em *RobboUnitDB) FromCore(robboUnit *RobboUnitCore) {
	em.Name = robboUnit.Name
	em.City = robboUnit.City
}

func (ht *RobboUnitHTTP) ToCore() *RobboUnitCore {
	return &RobboUnitCore{
		LastModified: ht.LastModified,
		Name:         ht.Name,
		City:         ht.City,
	}
}

func (ht *RobboUnitHTTP) FromCore(robboUnit *RobboUnitCore) {
	ht.LastModified = robboUnit.LastModified
	ht.Name = robboUnit.Name
	ht.City = robboUnit.City
}
