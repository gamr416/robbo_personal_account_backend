package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
	"log"
)

type UsersUseCaseImpl struct {
	users.Gateway
}

type UsersUseCaseModule struct {
	fx.Out
	users.UseCase
}

func SetupUsersUseCase(gateway users.Gateway) UsersUseCaseModule {
	return UsersUseCaseModule{
		UseCase: &UsersUseCaseImpl{
			Gateway: gateway,
		},
	}
}

func (p *UsersUseCaseImpl) GetUnitAdminById(unitAdminId uint) (unitAdmin *models.UnitAdminCore, err error) {
	return p.Gateway.GetUnitAdminById(unitAdminId)
}

func (p *UsersUseCaseImpl) GetTeacherByID(teacherId uint) (teacher *models.TeacherCore, err error) {
	return p.Gateway.GetTeacherById(teacherId)
}

func (p *UsersUseCaseImpl) GetTeacher(email, password string) (teacher *models.TeacherCore, err error) {
	return p.Gateway.GetTeacher(email, password)
}

func (p *UsersUseCaseImpl) GetParent(email, password string) (parent *models.ParentCore, err error) {
	return p.Gateway.GetParent(email, password)
}
func (p *UsersUseCaseImpl) GetSuperAdmin(email, password string) (superAdmin *models.SuperAdminCore, err error) {
	return p.Gateway.GetSuperAdmin(email, password)
}
func (p *UsersUseCaseImpl) GetUnitAdmin(email, password string) (unitAdmin *models.UnitAdminCore, err error) {
	return p.Gateway.GetUnitAdmin(email, password)
}
func (p *UsersUseCaseImpl) GetParentById(parentId uint) (parent *models.ParentCore, err error) {
	return p.Gateway.GetParentById(parentId)
}

func (p *UsersUseCaseImpl) UpdateUnitAdmin(unitAdmin *models.UnitAdminCore) (err error) {
	err = p.Gateway.UpdateUnitAdmin(unitAdmin)
	if err != nil {
		log.Println("Error update Unit Admin")
		return
	}
	return
}

func (p *UsersUseCaseImpl) UpdateParent(parent *models.ParentCore) (err error) {
	err = p.Gateway.UpdateParent(parent)
	if err != nil {
		log.Println("Error update Parent")
		return
	}
	return
}

func (p *UsersUseCaseImpl) UpdateTeacher(teacher *models.TeacherCore) (err error) {
	err = p.Gateway.UpdateTeacher(teacher)
	if err != nil {
		log.Println("Error update Teacher")
		return
	}
	return
}

func (p *UsersUseCaseImpl) CreateUnitAdmin(unitAdmin *models.UnitAdminCore) (id string, err error) {
	return p.Gateway.CreateUnitAdmin(unitAdmin)
}

func (p *UsersUseCaseImpl) DeleteUnitAdmin(unitAdminId uint) (err error) {
	return p.Gateway.DeleteUnitAdmin(unitAdminId)
}

func (p *UsersUseCaseImpl) CreateTeacher(teacher *models.TeacherCore) (id string, err error) {
	return p.Gateway.CreateTeacher(teacher)
}

func (p *UsersUseCaseImpl) DeleteTeacher(teacherId uint) (err error) {
	return p.Gateway.DeleteTeacher(teacherId)
}

func (p *UsersUseCaseImpl) CreateParent(parent *models.ParentCore) (id string, err error) {
	return p.Gateway.CreateParent(parent)
}

func (p *UsersUseCaseImpl) DeleteParent(parentId uint) (err error) {
	return p.Gateway.DeleteParent(parentId)
}

func (p *UsersUseCaseImpl) GetTeacherById(teacherId uint) (teacher *models.TeacherCore, err error) {
	return p.Gateway.GetTeacherById(teacherId)
}

func (p *UsersUseCaseImpl) CreateStudent(student *models.StudentCore) (id string, err error) {
	return p.Gateway.CreateStudent(student)
}

func (p *UsersUseCaseImpl) DeleteStudent(studentId uint) (err error) {
	return p.Gateway.DeleteStudent(studentId)
}

func (p *UsersUseCaseImpl) GetStudentById(studentId uint) (student *models.StudentCore, err error) {
	return p.Gateway.GetStudentById(studentId)
}

func (p *UsersUseCaseImpl) GetSuperAdminById(superAdminId uint) (superAdmin *models.SuperAdminCore, err error) {
	return p.Gateway.GetSuperAdminById(superAdminId)
}

func (p *UsersUseCaseImpl) UpdateStudent(student *models.StudentCore) (err error) {
	err = p.Gateway.UpdateStudent(student)
	if err != nil {
		log.Println("Error update student")
		return
	}
	return
}

func (p *UsersUseCaseImpl) GetStudent(email, password string) (student *models.StudentCore, err error) {
	return p.Gateway.GetStudent(email, password)
}
