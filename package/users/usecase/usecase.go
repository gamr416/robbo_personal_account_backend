package usecase

import (
	"crypto/sha1"
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"strconv"
)

type UsersUseCaseImpl struct {
	Gateway users.Gateway
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

//func (p *UsersUseCaseImpl) GetStudent(email, password string) (student *models.StudentCore, err error) {
//	return p.Gateway.GetStudent(email, password)
//}

func (p *UsersUseCaseImpl) GetStudentById(studentId string) (student *models.StudentCore, err error) {
	return p.Gateway.GetStudentById(studentId)
}

func (p *UsersUseCaseImpl) SearchStudentByEmail(email string, parentId string) (students []*models.StudentCore, err error) {
	emailCondition := "%" + email + "%"
	return p.Gateway.SearchStudentsByEmail(emailCondition, parentId)
}

func (p *UsersUseCaseImpl) GetStudentByParentId(parentId string) (students []*models.StudentCore, err error) {
	relations, getRelationErr := p.Gateway.GetRelationByParentId(parentId)
	if getRelationErr != nil {
		err = getRelationErr
		return
	}
	for _, relation := range relations {
		student, getStudentErr := p.Gateway.GetStudentById(relation.ChildId)
		if getStudentErr != nil {
			err = getStudentErr
			return
		}
		students = append(students, student)
	}
	return
}

func (p *UsersUseCaseImpl) CreateStudent(student *models.StudentCore, parentId string) (id string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(student.Password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))
	passwordHash := fmt.Sprintf("%x", pwd.Sum(nil))
	student.Password = passwordHash
	id, err = p.Gateway.CreateStudent(student)
	if err != nil {
		return
	}
	relation := &models.ChildrenOfParentCore{
		ChildId:  id,
		ParentId: parentId,
	}
	p.Gateway.CreateRelation(relation)
	return
}

func (p *UsersUseCaseImpl) DeleteStudent(studentId uint) (err error) {

	deleteRelationErr := p.Gateway.DeleteRelationByChildrenId(strconv.Itoa(int(studentId)))
	if deleteRelationErr != nil {
		return deleteRelationErr
	}
	return p.Gateway.DeleteStudent(studentId)
}

func (p *UsersUseCaseImpl) UpdateStudent(student *models.StudentCore) (err error) {
	err = p.Gateway.UpdateStudent(student)
	if err != nil {
		log.Println("Error update student")
		return
	}
	return
}

func (p *UsersUseCaseImpl) AddStudentToRobboGroup(studentId string, robboGroupId string, robboUnitId string) (err error) {
	return p.Gateway.AddStudentToRobboGroup(studentId, robboGroupId, robboUnitId)
}

func (p *UsersUseCaseImpl) GetTeacherById(teacherId string) (teacher models.TeacherCore, err error) {
	return p.Gateway.GetTeacherById(teacherId)
}

//func (p *UsersUseCaseImpl) GetTeacher(email, password string) (teacher *models.TeacherCore, err error) {
//	return p.Gateway.GetTeacher(email, password)
//}

func (p *UsersUseCaseImpl) GetAllTeachers() (teachers []models.TeacherCore, err error) {
	return p.Gateway.GetAllTeachers()
}

func (p *UsersUseCaseImpl) UpdateTeacher(teacher *models.TeacherCore) (err error) {
	err = p.Gateway.UpdateTeacher(teacher)
	if err != nil {
		log.Println("Error update Teacher")
		return
	}
	return
}

func (p *UsersUseCaseImpl) CreateTeacher(teacher *models.TeacherCore) (id string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(teacher.Password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))
	passwordHash := fmt.Sprintf("%x", pwd.Sum(nil))
	teacher.Password = passwordHash
	return p.Gateway.CreateTeacher(teacher)
}

func (p *UsersUseCaseImpl) DeleteTeacher(teacherId uint) (err error) {
	return p.Gateway.DeleteTeacher(teacherId)
}

//func (p *UsersUseCaseImpl) GetParent(email, password string) (parent *models.ParentCore, err error) {
//	return p.Gateway.GetParent(email, password)
//}

func (p *UsersUseCaseImpl) GetParentById(parentId string) (parent *models.ParentCore, err error) {
	return p.Gateway.GetParentById(parentId)
}

func (p *UsersUseCaseImpl) GetAllParent() (parents []*models.ParentCore, err error) {
	parents, err = p.Gateway.GetAllParent()
	return
}

func (p *UsersUseCaseImpl) CreateParent(parent *models.ParentCore) (id string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(parent.Password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))
	passwordHash := fmt.Sprintf("%x", pwd.Sum(nil))
	parent.Password = passwordHash
	return p.Gateway.CreateParent(parent)
}

func (p *UsersUseCaseImpl) DeleteParent(parentId uint) (err error) {
	relations, getRelationsErr := p.Gateway.GetRelationByParentId(strconv.Itoa(int(parentId)))
	if getRelationsErr != nil {
		return getRelationsErr
	}

	for _, relation := range relations {
		studentId, _ := strconv.ParseUint(relation.ChildId, 10, 64)
		deleteStudentErr := p.Gateway.DeleteStudent(uint(studentId))
		if deleteStudentErr != nil {
			return deleteStudentErr
		}
	}
	deleteRelationErr := p.Gateway.DeleteRelationByParentId(strconv.Itoa(int(parentId)))
	if deleteRelationErr != nil {
		return deleteRelationErr
	}
	return p.Gateway.DeleteParent(parentId)
}

func (p *UsersUseCaseImpl) UpdateParent(parent *models.ParentCore) (err error) {
	err = p.Gateway.UpdateParent(parent)
	if err != nil {
		log.Println("Error update Parent")
		return
	}
	return
}

//func (p *UsersUseCaseImpl) GetFreeListener(email, password string) (freeListener *models.FreeListenerCore, err error) {
//	return p.Gateway.GetFreeListener(email, password)
//}

func (p *UsersUseCaseImpl) GetFreeListenerById(freeListenerId string) (freeListener *models.FreeListenerCore, err error) {
	return p.Gateway.GetFreeListenerById(freeListenerId)
}

func (p *UsersUseCaseImpl) CreateFreeListener(freeListener *models.FreeListenerCore) (id string, err error) {
	return p.Gateway.CreateFreeListener(freeListener)
}

func (p *UsersUseCaseImpl) DeleteFreeListener(freeListener uint) (err error) {
	return p.Gateway.DeleteFreeListener(freeListener)
}

func (p *UsersUseCaseImpl) UpdateFreeListener(freeListener *models.FreeListenerCore) (err error) {
	err = p.Gateway.UpdateFreeListener(freeListener)
	if err != nil {
		log.Println("Error update Parent")
		return
	}
	return
}

func (p *UsersUseCaseImpl) GetUnitAdminById(unitAdminId string) (unitAdmin *models.UnitAdminCore, err error) {
	return p.Gateway.GetUnitAdminById(unitAdminId)
}

func (p *UsersUseCaseImpl) GetAllUnitAdmins() (unitAdmins []*models.UnitAdminCore, err error) {
	return p.Gateway.GetAllUnitAdmins()
}

//func (p *UsersUseCaseImpl) GetUnitAdmin(email, password string) (unitAdmin *models.UnitAdminCore, err error) {
//	return p.Gateway.GetUnitAdmin(email, password)
//}

func (p *UsersUseCaseImpl) UpdateUnitAdmin(unitAdmin *models.UnitAdminCore) (err error) {
	err = p.Gateway.UpdateUnitAdmin(unitAdmin)
	if err != nil {
		log.Println("Error update Unit Admin")
		return
	}
	return
}

func (p *UsersUseCaseImpl) CreateUnitAdmin(unitAdmin *models.UnitAdminCore) (id string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(unitAdmin.Password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))
	passwordHash := fmt.Sprintf("%x", pwd.Sum(nil))
	unitAdmin.Password = passwordHash
	return p.Gateway.CreateUnitAdmin(unitAdmin)
}

func (p *UsersUseCaseImpl) DeleteUnitAdmin(unitAdminId uint) (err error) {
	return p.Gateway.DeleteUnitAdmin(unitAdminId)
}

func (p *UsersUseCaseImpl) SearchUnitAdminByEmail(email string, robboUnitId string) (unitAdmins []*models.UnitAdminCore, err error) {
	emailCondition := "%" + email + "%"
	return p.Gateway.SearchUnitAdminByEmail(emailCondition, robboUnitId)
}

func (p *UsersUseCaseImpl) GetUnitAdminByRobboUnitId(robboUnitId string) (unitAdmins []*models.UnitAdminCore, err error) {
	relations, getRelationErr := p.Gateway.GetRelationByRobboUnitId(robboUnitId)
	if getRelationErr != nil {
		err = getRelationErr
		return
	}

	for _, relation := range relations {
		unitAdmin, getUnitAdminErr := p.Gateway.GetUnitAdminById(relation.UnitAdminId)
		if getUnitAdminErr != nil {
			err = getRelationErr
			return
		}
		unitAdmins = append(unitAdmins, unitAdmin)
	}
	return
}

func (p *UsersUseCaseImpl) GetSuperAdminById(superAdminId string) (superAdmin *models.SuperAdminCore, err error) {
	return p.Gateway.GetSuperAdminById(superAdminId)
}

func (p *UsersUseCaseImpl) UpdateSuperAdmin(superAdmin *models.SuperAdminCore) (err error) {
	err = p.Gateway.UpdateSuperAdmin(superAdmin)
	if err != nil {
		log.Println("Error update Super Admin")
		return
	}
	return
}
func (p *UsersUseCaseImpl) DeleteSuperAdmin(superAdminId uint) (err error) {
	return p.Gateway.DeleteSuperAdmin(superAdminId)
}

func (p *UsersUseCaseImpl) CreateRelation(parentId, childrenId string) (err error) {
	relationCore := &models.ChildrenOfParentCore{
		ChildId:  childrenId,
		ParentId: parentId,
	}
	return p.Gateway.CreateRelation(relationCore)
}

func (p *UsersUseCaseImpl) SetNewUnitAdminForRobboUnit(unitAdminId, robboUnitId string) (err error) {
	relationCore := &models.UnitAdminsRobboUnitsCore{
		UnitAdminId: unitAdminId,
		RobboUnitId: robboUnitId,
	}
	return p.Gateway.SetUnitAdminForRobboUnit(relationCore)
}

func (p *UsersUseCaseImpl) DeleteUnitAdminForRobboUnit(unitAdminId, robboUnitId string) (err error) {
	relationCore := &models.UnitAdminsRobboUnitsCore{
		UnitAdminId: unitAdminId,
		RobboUnitId: robboUnitId,
	}
	return p.Gateway.DeleteUnitAdminForRobboUnit(relationCore)
}
