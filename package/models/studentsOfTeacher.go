package models

import "gorm.io/gorm"

/*
	вспомогательная структура
   	для хранения связи между ребенком и учителем
*/

type StudentsOfTeacherCore struct {
	StudentId string
	TeacherId string
}

type StudentsOfTeacherDB struct {
	gorm.Model

	StudentId string
	Student   StudentDB `gorm:"foreignKey:StudentId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TeacherId string
	Teacher   TeacherDB `gorm:"foreignKey:TeacherId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (em *StudentsOfTeacherDB) FromCore(core *StudentsOfTeacherCore) {
	em.StudentId = core.StudentId
	em.TeacherId = core.TeacherId
}

func (em *StudentsOfTeacherDB) ToCore() *StudentsOfTeacherCore {
	return &StudentsOfTeacherCore{
		StudentId: em.StudentId,
		TeacherId: em.TeacherId,
	}
}
