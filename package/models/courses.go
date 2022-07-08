package models

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type CourseCore struct {
	ID               string
	BlocksUrl        string
	Effort           string
	EnrollmentStart  time.Time
	EnrollmentEnd    time.Time
	Name             string
	Number           string
	Org              string
	ShortDescription string
	Start            time.Time
	StartDisplay     string
	StartType        string
	Pacing           string
	MobileAvailable  bool
	Hidden           bool
	InvitationOnly   bool
	CourseID         string
}

type CourseDB struct {
	gorm.Model

	BlocksUrl        string
	Effort           string
	EnrollmentStart  time.Time
	EnrollmentEnd    time.Time
	Name             string
	Number           string
	Org              string
	ShortDescription string
	Start            time.Time
	StartDisplay     string
	StartType        string
	Pacing           string
	MobileAvailable  bool
	Hidden           bool
	InvitationOnly   bool
	CourseID         string
}

type CourseHTTP struct {
	BlocksUrl        string      `json:"blocks_url"`
	Effort           string      `json:"effort"`
	End              interface{} `json:"end"`
	EnrollmentStart  time.Time   `json:"enrollment_start"`
	EnrollmentEnd    time.Time   `json:"enrollment_end"`
	ID               string      `json:"id"`
	Media            interface{} `json:"media"`
	Name             string      `json:"name"`
	Number           string      `json:"number"`
	Org              string      `json:"org"`
	ShortDescription string      `json:"short_description"`
	Start            time.Time   `json:"start"`
	StartDisplay     string      `json:"start_display"`
	StartType        string      `json:"start_type"`
	Pacing           string      `json:"pacing"`
	MobileAvailable  bool        `json:"mobile_available"`
	Hidden           bool        `json:"hidden"`
	InvitationOnly   bool        `json:"invitation_only"`
	CourseID         string      `json:"course_id"`
	Overview         interface{} `json:"overview"`
}

func (em *CourseDB) ToCore() *CourseCore {
	return &CourseCore{
		ID:               strconv.FormatUint(uint64(em.ID), 10),
		BlocksUrl:        em.BlocksUrl,
		Effort:           em.Effort,
		EnrollmentStart:  em.EnrollmentStart,
		EnrollmentEnd:    em.EnrollmentEnd,
		Name:             em.Name,
		Number:           em.Number,
		Org:              em.Org,
		ShortDescription: em.ShortDescription,
		Start:            em.Start,
		StartDisplay:     em.StartDisplay,
		StartType:        em.StartType,
		Pacing:           em.Pacing,
		MobileAvailable:  em.MobileAvailable,
		Hidden:           em.Hidden,
		InvitationOnly:   em.InvitationOnly,
		CourseID:         em.CourseID,
	}
}

func (em *CourseDB) FromCore(course *CourseCore) {
	id, _ := strconv.ParseUint(course.ID, 10, 64)
	em.ID = uint(id)
	em.BlocksUrl = course.BlocksUrl
	em.Effort = course.Effort
	em.EnrollmentStart = course.EnrollmentStart
	em.EnrollmentEnd = course.EnrollmentEnd
	em.Name = course.Name
	em.Number = course.Number
	em.Org = course.Org
	em.ShortDescription = course.ShortDescription
	em.Start = course.Start
	em.StartDisplay = course.StartDisplay
	em.StartType = course.StartType
	em.Pacing = course.Pacing
	em.MobileAvailable = course.MobileAvailable
	em.Hidden = course.Hidden
	em.InvitationOnly = course.InvitationOnly
	em.CourseID = course.CourseID
}

func (ht *CourseHTTP) FromCore(course *CourseCore) {
	ht.ID = course.ID
	ht.BlocksUrl = course.BlocksUrl
	ht.Effort = course.Effort
	ht.EnrollmentStart = course.EnrollmentStart
	ht.EnrollmentEnd = course.EnrollmentEnd
	ht.Name = course.Name
	ht.Number = course.Number
	ht.Org = course.Org
	ht.ShortDescription = course.ShortDescription
	ht.Start = course.Start
	ht.StartDisplay = course.StartDisplay
	ht.StartType = course.StartType
	ht.Pacing = course.Pacing
	ht.MobileAvailable = course.MobileAvailable
	ht.Hidden = course.Hidden
	ht.InvitationOnly = course.InvitationOnly
	ht.CourseID = course.CourseID
}

func (ht *CourseHTTP) ToCore() *CourseCore {
	return &CourseCore{
		ID:               ht.ID,
		BlocksUrl:        ht.BlocksUrl,
		Effort:           ht.Effort,
		EnrollmentStart:  ht.EnrollmentStart,
		EnrollmentEnd:    ht.EnrollmentEnd,
		Name:             ht.Name,
		Number:           ht.Number,
		Org:              ht.Org,
		ShortDescription: ht.ShortDescription,
		Start:            ht.Start,
		StartDisplay:     ht.StartDisplay,
		StartType:        ht.StartType,
		Pacing:           ht.Pacing,
		MobileAvailable:  ht.MobileAvailable,
		Hidden:           ht.Hidden,
		InvitationOnly:   ht.InvitationOnly,
		CourseID:         ht.CourseID,
	}
}
