package models

import (
	"gorm.io/gorm"
	"strconv"
)

type CourseApiMediaCollectionCore struct {
	ID          string
	BannerImage AbsoluteMediaCore
	CourseImage MediaCore
	CourseVideo MediaCore
	Image       ImageCore
	CourseID    string
}

type CourseApiMediaCollectionDB struct {
	gorm.Model
	CourseID uint
	Course   CourseDB `gorm:"foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (em *CourseApiMediaCollectionDB) ToCore() *CourseApiMediaCollectionCore {
	return &CourseApiMediaCollectionCore{
		ID:       strconv.FormatUint(uint64(em.ID), 10),
		CourseID: strconv.FormatUint(uint64(em.CourseID), 10),
	}
}

func (em *CourseApiMediaCollectionDB) FromCore(courseApiMediaCollection *CourseApiMediaCollectionCore) {
	id, _ := strconv.ParseUint(courseApiMediaCollection.ID, 10, 64)
	courseId, _ := strconv.ParseUint(courseApiMediaCollection.CourseID, 10, 64)
	em.ID = uint(id)
	em.CourseID = uint(courseId)
}

func (ht *CourseAPIMediaCollectionHTTP) ToCore() *CourseApiMediaCollectionCore {
	bannerImageCore := &AbsoluteMediaCore{}
	bannerImageCore = ht.BannerImage.ToCore()
	courseImageCore := &MediaCore{}
	courseImageCore = ht.CourseImage.ToCore()
	courseVideoCore := &MediaCore{}
	courseVideoCore = ht.CourseVideo.ToCore()
	imageCore := &ImageCore{}
	imageCore = ht.Image.ToCore()
	return &CourseApiMediaCollectionCore{
		ID:          ht.ID,
		BannerImage: *bannerImageCore,
		CourseImage: *courseImageCore,
		CourseVideo: *courseVideoCore,
		Image:       *imageCore,
	}
}

func (ht *CourseAPIMediaCollectionHTTP) FromCore(courseApiMediaCollection *CourseApiMediaCollectionCore) {
	ht.ID = courseApiMediaCollection.ID
	ht.BannerImage.FromCore(&courseApiMediaCollection.BannerImage)
	ht.CourseImage.FromCore(&courseApiMediaCollection.CourseImage)
	ht.CourseVideo.FromCore(&courseApiMediaCollection.CourseVideo)
	ht.Image.FromCore(&courseApiMediaCollection.Image)
}
