package users

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	GetStudentById(studentId uint) (student *models.StudentCore, err error)
	//GetStudent(email, password string) (student *models.StudentCore, err error)
	CreateStudent(student *models.StudentCore, parentId string) (id string, err error)
	DeleteStudent(studentId uint) (err error)
	UpdateStudent(student *models.StudentCore) (err error)

	//GetTeacher(email, password string) (teacher *models.TeacherCore, err error)
	GetTeacherById(teacherId uint) (teacher *models.TeacherCore, err error)
	GetAllTeachers() (teachers []*models.TeacherCore, err error)
	CreateTeacher(teacher *models.TeacherCore) (id string, err error)
	UpdateTeacher(teacher *models.TeacherCore) (err error)
	DeleteTeacher(teacherId uint) (err error)

	//GetParent(email, password string) (parent *models.ParentCore, err error)
	GetParentById(parentId uint) (parent *models.ParentCore, err error)
	GetAllParent() (parents []*models.ParentCore, err error)
	CreateParent(parent *models.ParentCore) (id string, err error)
	UpdateParent(parent *models.ParentCore) (err error)
	DeleteParent(parentId uint) (err error)

	//GetFreeListener(email, password string) (freeListener *models.FreeListenerCore, err error)
	GetFreeListenerById(freeListenerId uint) (freeListener *models.FreeListenerCore, err error)
	CreateFreeListener(freeListener *models.FreeListenerCore) (id string, err error)
	UpdateFreeListener(freeListener *models.FreeListenerCore) (err error)
	DeleteFreeListener(freeListenerId uint) (err error)

	GetUnitAdminById(unitAdminId uint) (unitAdmin *models.UnitAdminCore, err error)
	//GetUnitAdmin(email, password string) (unitAdmin *models.UnitAdminCore, err error)
	UpdateUnitAdmin(unitAdmin *models.UnitAdminCore) (err error)
	CreateUnitAdmin(unitAdmin *models.UnitAdminCore) (id string, err error)
	DeleteUnitAdmin(unitAdminId uint) (err error)

	//GetSuperAdmin(email, password string) (superAdmin *models.SuperAdminCore, err error)
	GetSuperAdminById(superAdminId uint) (superAdmin *models.SuperAdminCore, err error)
	UpdateSuperAdmin(superAdmin *models.SuperAdminCore) (err error)
	DeleteSuperAdmin(superAdminId uint) (err error)

	CreateRelation(parentId, childrenId string) (err error)
}
