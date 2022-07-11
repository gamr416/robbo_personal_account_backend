package models

import (
	"gorm.io/gorm"
	"strconv"
)

type ImageCore struct {
	ID                         string
	Raw                        string
	Small                      string
	Large                      string
	CourseApiMediaCollectionID string
}

type ImageDB struct {
	gorm.Model
	Raw                        string
	Small                      string
	Large                      string
	CourseApiMediaCollectionID uint
	CourseApiMediaCollection   CourseApiMediaCollectionDB `gorm:"foreignKey:CourseApiMediaCollectionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ImageHTTP struct {
	ID    string `json:"id"`
	Raw   string `json:"raw"`
	Small string `json:"small"`
	Large string `json:"large"`
}

func (em *ImageDB) ToCore() *ImageCore {
	return &ImageCore{
		ID:                         strconv.FormatUint(uint64(em.ID), 10),
		Raw:                        em.Raw,
		Small:                      em.Small,
		Large:                      em.Large,
		CourseApiMediaCollectionID: strconv.FormatUint(uint64(em.CourseApiMediaCollectionID), 10),
	}
}

func (em *ImageDB) FromCore(image *ImageCore) {
	id, _ := strconv.ParseUint(image.ID, 10, 64)
	courseApiMediaCollectionId, _ := strconv.ParseUint(image.CourseApiMediaCollectionID, 10, 64)
	em.ID = uint(id)
	em.Raw = image.Raw
	em.Small = image.Small
	em.Large = image.Large
	em.CourseApiMediaCollectionID = uint(courseApiMediaCollectionId)
}

func (ht *ImageHTTP) ToCore() *ImageCore {
	return &ImageCore{
		ID:    ht.ID,
		Raw:   ht.Raw,
		Small: ht.Small,
		Large: ht.Large,
	}
}

func (ht *ImageHTTP) FromCore(imageCore *ImageCore) {
	ht.ID = imageCore.ID
	ht.Raw = imageCore.Raw
	ht.Small = imageCore.Small
	ht.Large = imageCore.Large
}
