package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
)

func (r *RobboGroupUseCaseImpl) CreateRobboGroup(robboGroup *models.RobboGroupCore) (newRobboGroup *models.RobboGroupCore, err error) {
	return r.robboGroupGateway.CreateRobboGroup(robboGroup)
}

func (r *RobboGroupUseCaseImpl) DeleteRobboGroup(robboGroupId string) (err error) {
	if err = r.robboGroupGateway.DeleteRobboGroup(robboGroupId); err != nil {
		return
	}
	relations, err := r.robboGroupGateway.GetRelationByRobboGroupId(robboGroupId)
	if err != nil {
		return err
	}
	students, getStudentErr := r.usersGateway.GetStudentsByRobboGroupId(robboGroupId)
	if getStudentErr != nil {
		err = getStudentErr
		return
	}
	for _, student := range students {
		for _, relation := range relations {
			relationCore := &models.StudentsOfTeacherCore{
				StudentId: student.Id,
				TeacherId: relation.TeacherId,
			}
			if err = r.usersGateway.DeleteStudentTeacherRelation(relationCore); err != nil {
				return
			}
		}
		student.RobboGroupId = ""
		if _, err = r.usersGateway.UpdateStudent(student); err != nil {
			return
		}
	}
	return
}

func (r *RobboGroupUseCaseImpl) GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroups []*models.RobboGroupCore, err error) {
	return r.robboGroupGateway.GetRobboGroupsByRobboUnitId(robboUnitId)
}

func (r *RobboGroupUseCaseImpl) GetRobboGroupsByUnitAdminId(unitAdminId string, page, pageSize int) (robboGroups []*models.RobboGroupCore, countRows int64, err error) {
	relations, _, getRelationErr := r.usersGateway.GetRelationByUnitAdminId(unitAdminId, 0, 0)
	if getRelationErr != nil {
		err = getRelationErr
		return
	}
	if relations == nil {
		return
	}
	var robboUnitsIds []string
	for _, relation := range relations {
		robboUnitsIds = append(robboUnitsIds, relation.RobboUnitId)
	}
	unitRobboGroups, countRows, getRobboGroupErr := r.robboGroupGateway.GetRobboGroupsByRobboUnitsIds(robboUnitsIds, page, pageSize)
	if getRobboGroupErr != nil {
		err = getRelationErr
		return
	}
	robboGroups = append(robboGroups, unitRobboGroups...)
	return
}

func (r *RobboGroupUseCaseImpl) GetAllRobboGroups(page, pageSize int) (robboGroups []*models.RobboGroupCore, countRows int64, err error) {
	return r.robboGroupGateway.GetAllRobboGroups(page, pageSize)
}

func (r *RobboGroupUseCaseImpl) GetRobboGroupById(robboGroupId string) (robboGroup *models.RobboGroupCore, err error) {
	robboGroup, err = r.robboGroupGateway.GetRobboGroupById(robboGroupId)
	if err != nil {
		return
	}
	students, getStudentErr := r.usersGateway.GetStudentsByRobboGroupId(robboGroup.Id)
	if getStudentErr != nil {
		err = getStudentErr
		return
	}
	robboGroup.Students = students
	return
}

type RobboGroupUseCaseImpl struct {
	robboGroupGateway robboGroup.Gateway
	usersGateway      users.Gateway
}

func (r *RobboGroupUseCaseImpl) UpdateRobboGroup(robboGroup *models.RobboGroupCore) (robboGroupUpdated *models.RobboGroupCore, err error) {
	return r.robboGroupGateway.UpdateRobboGroup(robboGroup)
}

func (r *RobboGroupUseCaseImpl) SearchRobboGroupsByTitle(title string) (robboGroups []*models.RobboGroupCore, err error) {
	titleCondition := "%" + title + "%"
	return r.robboGroupGateway.SearchRobboGroupsByTitle(titleCondition)
}

func (r *RobboGroupUseCaseImpl) SetTeacherForRobboGroup(teacherId, robboGroupId string) (err error) {
	relationCore := &models.TeachersRobboGroupsCore{
		TeacherId:    teacherId,
		RobboGroupId: robboGroupId,
	}
	if err = r.robboGroupGateway.SetTeacherForRobboGroup(relationCore); err != nil {
		return
	}
	students, getStudentErr := r.usersGateway.GetStudentsByRobboGroupId(robboGroupId)
	if getStudentErr != nil {
		err = getStudentErr
		return
	}
	for _, student := range students {
		relationCore := &models.StudentsOfTeacherCore{
			StudentId: student.Id,
			TeacherId: teacherId,
		}
		if err = r.usersGateway.CreateStudentTeacherRelation(relationCore); err != nil {
			return
		}
	}
	return
}

func (r *RobboGroupUseCaseImpl) DeleteTeacherForRobboGroup(teacherId, robboGroupId string) (err error) {
	relationCore := &models.TeachersRobboGroupsCore{
		TeacherId:    teacherId,
		RobboGroupId: robboGroupId,
	}
	if err = r.robboGroupGateway.DeleteTeacherForRobboGroup(relationCore); err != nil {
		return
	}
	students, getStudentErr := r.usersGateway.GetStudentsByRobboGroupId(robboGroupId)
	if getStudentErr != nil {
		err = getStudentErr
		return
	}
	for _, student := range students {
		relationCore := &models.StudentsOfTeacherCore{
			StudentId: student.Id,
			TeacherId: teacherId,
		}
		if err = r.usersGateway.DeleteStudentTeacherRelation(relationCore); err != nil {
			return
		}
	}
	return
}

func (r *RobboGroupUseCaseImpl) GetRobboGroupsByTeacherId(teacherId string, page, pageSize int) (
	robboGroups []*models.RobboGroupCore,
	countRows int64,
	err error,
) {
	relations, countRows, getRelationsErr := r.robboGroupGateway.GetRelationByTeacherId(teacherId, page, pageSize)
	if getRelationsErr != nil {
		err = getRelationsErr
		return
	}
	for _, relation := range relations {
		robboGroup, getRobboGroupErr := r.robboGroupGateway.GetRobboGroupById(relation.RobboGroupId)
		if getRobboGroupErr != nil {
			err = getRobboGroupErr
			return
		}
		robboGroups = append(robboGroups, robboGroup)
	}
	return
}

type RobboGroupUseCaseModule struct {
	fx.Out
	robboGroup.UseCase
}

func SetupRobboGroupUseCase(robboGroupGateway robboGroup.Gateway, usersGateway users.Gateway) RobboGroupUseCaseModule {
	return RobboGroupUseCaseModule{
		UseCase: &RobboGroupUseCaseImpl{
			robboGroupGateway,
			usersGateway,
		},
	}
}
