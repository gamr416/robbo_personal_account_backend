package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup"
	"go.uber.org/fx"
)

type RobboGroupDelegateImpl struct {
	UseCase robboGroup.UseCase
}

func (r *RobboGroupDelegateImpl) UpdateRobboGroup(robboGroup *models.RobboGroupHTTP) (err error) {
	robboGroupCore := robboGroup.ToCore()
	return r.UseCase.UpdateRobboGroup(robboGroupCore)
}

func (r *RobboGroupDelegateImpl) GetRobboGroupsByTeacherId(teacherId string) (robboGroups []*models.RobboGroupHTTP, err error) {
	robboGroupsCore, err := r.UseCase.GetRobboGroupsByTeacherId(teacherId)
	for _, robboGroupCore := range robboGroupsCore {
		var robboGroupTemp models.RobboGroupHTTP
		robboGroupTemp.FromCore(robboGroupCore)
		robboGroups = append(robboGroups, &robboGroupTemp)
	}
	return
}

func (r *RobboGroupDelegateImpl) SearchRobboGroupByName(name string) (robboGroups []*models.RobboGroupHTTP, err error) {
	robboGroupsCore, err := r.UseCase.SearchRobboGroupsByTitle(name)
	for _, robboGroupCore := range robboGroupsCore {
		var robboGroupTemp models.RobboGroupHTTP
		robboGroupTemp.FromCore(robboGroupCore)
		robboGroups = append(robboGroups, &robboGroupTemp)
	}
	return
}

func (r *RobboGroupDelegateImpl) SetTeacherForRobboGroup(teacherId, robboGroupId string) (err error) {
	return r.UseCase.SetTeacherForRobboGroup(teacherId, robboGroupId)
}

func (r *RobboGroupDelegateImpl) DeleteTeacherForRobboGroup(teacherId, robboGroupId string) (err error) {
	return r.UseCase.DeleteTeacherForRobboGroup(teacherId, robboGroupId)
}

func (r *RobboGroupDelegateImpl) CreateRobboGroup(robboGroup *models.RobboGroupHTTP) (robboGroupId string, err error) {
	robboGroupCore := robboGroup.ToCore()
	return r.UseCase.CreateRobboGroup(robboGroupCore)
}

func (r *RobboGroupDelegateImpl) DeleteRobboGroup(robboGroupId string) (err error) {
	return r.UseCase.DeleteRobboGroup(robboGroupId)
}

func (r *RobboGroupDelegateImpl) GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroups []*models.RobboGroupHTTP, err error) {
	robboGroupsCore, err := r.UseCase.GetRobboGroupsByRobboUnitId(robboUnitId)
	if err != nil {
		return
	}
	for _, robboGroupCore := range robboGroupsCore {
		var robboGroupTemp models.RobboGroupHTTP
		robboGroupTemp.FromCore(robboGroupCore)
		robboGroups = append(robboGroups, &robboGroupTemp)
	}
	return
}

func (r *RobboGroupDelegateImpl) GetRobboGroupById(robboGroupId string) (robboGroup models.RobboGroupHTTP, err error) {
	robboGroupCore, err := r.UseCase.GetRobboGroupById(robboGroupId)
	if err != nil {
		return
	}
	robboGroup.FromCore(robboGroupCore)
	return
}

type RobboGroupDelegateModule struct {
	fx.Out
	robboGroup.Delegate
}

func SetupRobboGroupDelegate(usecase robboGroup.UseCase) RobboGroupDelegateModule {
	return RobboGroupDelegateModule{
		Delegate: &RobboGroupDelegateImpl{
			usecase,
		},
	}
}
