package models

import (
	"gorm.io/gorm"
	"strconv"
)

type UserCore struct {
	ID       string
	Email    string
	Password string
}

type UserDB struct {
	gorm.Model

	Email    string `gorm:"not null;size:256"`
	Password string `gorm:"not null;size:256"`
}

func (em *UserDB) ToCore() *UserCore {
	return &UserCore{
		ID:       strconv.FormatUint(uint64(em.ID), 10),
		Email:    em.Email,
		Password: em.Password,
	}
}

func (em *UserDB) FromCore(user *UserCore) {
	id, _ := strconv.ParseUint(user.ID, 10, 64)
	em.ID = uint(id)
	em.Email = user.Email
	em.Password = user.Password
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
