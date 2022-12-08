package users

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	GetStudentById(studentId string) (student *models.StudentCore, err error)
	GetStudentByParentId(parentId string) (students []*models.StudentCore, err error)
	GetStudentsByRobboGroupId(robboGroupId string) (students []*models.StudentCore, err error)
	GetStudentsByRobboUnitId(robboUnitId string) (students []*models.StudentCore, err error)
	GetStudentsByTeacherId(teacherId string) (students []*models.StudentCore, err error)
	//GetStudent(email, password string) (student *models.StudentCore, err error)
	SearchStudentByEmail(email string, parentId string) (students []*models.StudentCore, err error)
	CreateStudent(student *models.StudentCore, parentId string) (id string, err error)
	DeleteStudent(studentId uint) (err error)
	UpdateStudent(student *models.StudentCore) (err error)
	AddStudentToRobboGroup(studentId string, robboGroupId string, robboUnitId string) (err error)

	//GetTeacher(email, password string) (teacher *models.TeacherCore, err error)
	GetTeacherById(teacherId string) (teacher models.TeacherCore, err error)
	GetAllTeachers() (teachers []models.TeacherCore, err error)
	CreateTeacher(teacher *models.TeacherCore) (id string, err error)
	UpdateTeacher(teacher *models.TeacherCore) (err error)
	DeleteTeacher(teacherId uint) (err error)
	GetTeacherByRobboGroupId(robboGroupId string) (teachers []*models.TeacherCore, err error)
	GetTeachersByStudentId(studentId string) (teachers []*models.TeacherCore, err error)

	//GetParent(email, password string) (parent *models.ParentCore, err error)
	GetParentById(parentId string) (parent *models.ParentCore, err error)
	GetAllParent() (parents []*models.ParentCore, err error)
	CreateParent(parent *models.ParentCore) (id string, err error)
	UpdateParent(parent *models.ParentCore) (err error)
	DeleteParent(parentId uint) (err error)

	//GetFreeListener(email, password string) (freeListener *models.FreeListenerCore, err error)
	GetFreeListenerById(freeListenerId string) (freeListener *models.FreeListenerCore, err error)
	CreateFreeListener(freeListener *models.FreeListenerCore) (id string, err error)
	UpdateFreeListener(freeListener *models.FreeListenerCore) (err error)
	DeleteFreeListener(freeListenerId uint) (err error)

	GetUnitAdminById(unitAdminId string) (unitAdmin *models.UnitAdminCore, err error)
	GetAllUnitAdmins() (unitAdmins []*models.UnitAdminCore, err error)
	GetUnitAdminByRobboUnitId(robboUnitId string) (unitAdmins []*models.UnitAdminCore, err error)
	UpdateUnitAdmin(unitAdmin *models.UnitAdminCore) (err error)
	CreateUnitAdmin(unitAdmin *models.UnitAdminCore) (id string, err error)
	DeleteUnitAdmin(unitAdminId uint) (err error)
	SearchUnitAdminByEmail(email string, robboUnitId string) (unitAdmins []*models.UnitAdminCore, err error)

	//GetSuperAdmin(email, password string) (superAdmin *models.SuperAdminCore, err error)
	GetSuperAdminById(superAdminId string) (superAdmin *models.SuperAdminCore, err error)
	UpdateSuperAdmin(superAdmin *models.SuperAdminCore) (err error)
	DeleteSuperAdmin(superAdminId uint) (err error)

	CreateRelation(parentId, childrenId string) (err error)
	SetNewUnitAdminForRobboUnit(unitAdminId, robboUnitId string) (err error)
	DeleteUnitAdminForRobboUnit(unitAdminId, robboUnitId string) (err error)

	CreateStudentTeacherRelation(studentId, teacherId string) (student *models.StudentCore, err error)
	DeleteStudentTeacherRelation(studentId, teacherId string) (student *models.StudentCore, err error)
}
