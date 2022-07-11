package models

import (
	"gorm.io/gorm"
	"strconv"
)

type MediaCore struct {
	ID                         string
	Uri                        string
	CourseApiMediaCollectionID string
}

type MediaDB struct {
	gorm.Model
	Uri                        string
	CourseApiMediaCollectionID uint
	CourseApiMediaCollection   CourseApiMediaCollectionDB `gorm:"foreignKey:CourseApiMediaCollectionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type MediaHTTP struct {
	ID  string `json:"id"`
	Uri string `json:"uri"`
}

func (em *MediaDB) ToCore() *MediaCore {
	return &MediaCore{
		ID:                         strconv.FormatUint(uint64(em.ID), 10),
		Uri:                        em.Uri,
		CourseApiMediaCollectionID: strconv.FormatUint(uint64(em.CourseApiMediaCollectionID), 10),
	}
}

func (em *MediaDB) FromCore(media *MediaCore) {
	id, _ := strconv.ParseUint(media.ID, 10, 64)
	courseApiMediaCollectionId, _ := strconv.ParseUint(media.CourseApiMediaCollectionID, 10, 64)
	em.ID = uint(id)
	em.Uri = media.Uri
	em.CourseApiMediaCollectionID = uint(courseApiMediaCollectionId)
}

func (ht *MediaHTTP) ToCore() *MediaCore {
	return &MediaCore{
		ID:  ht.ID,
		Uri: ht.Uri,
	}
}

func (ht *MediaHTTP) FromCore(media *MediaCore) {
	ht.ID = media.ID
	ht.Uri = media.Uri
}
