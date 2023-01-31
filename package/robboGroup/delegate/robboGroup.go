package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup"
	"go.uber.org/fx"
	"log"
	"strconv"
)

type RobboGroupDelegateImpl struct {
	UseCase robboGroup.UseCase
}

func (r *RobboGroupDelegateImpl) UpdateRobboGroup(robboGroup *models.RobboGroupHTTP) (robboGroupUpdated models.RobboGroupHTTP, err error) {
	robboGroupCore := robboGroup.ToCore()
	robboGroupUpdatedCore, err := r.UseCase.UpdateRobboGroup(robboGroupCore)
	if err != nil {
		log.Println(err)
		return
	}
	robboGroupUpdated.FromCore(robboGroupUpdatedCore)
	return
}

func (r *RobboGroupDelegateImpl) GetRobboGroupsByTeacherId(teacherId, page, pageSize string) (
	robboGroups []*models.RobboGroupHTTP,
	countRows int,
	err error,
) {
	pageInt32, _ := strconv.ParseInt(page, 10, 32)
	pageSizeInt32, _ := strconv.ParseInt(pageSize, 10, 32)
	robboGroupsCore, countRowsInt64, err := r.UseCase.GetRobboGroupsByTeacherId(
		teacherId,
		int(pageInt32),
		int(pageSizeInt32),
	)
	if err != nil {
		return
	}
	countRows = int(countRowsInt64)
	for _, robboGroupCore := range robboGroupsCore {
		var robboGroupTemp models.RobboGroupHTTP
		robboGroupTemp.FromCore(robboGroupCore)
		robboGroups = append(robboGroups, &robboGroupTemp)
	}
	return
}

func (r *RobboGroupDelegateImpl) GetAllRobboGroups(page, pageSize string) (
	robboGroups []*models.RobboGroupHTTP,
	countRows int,
	err error,
) {
	pageInt32, _ := strconv.ParseInt(page, 10, 32)
	pageSizeInt32, _ := strconv.ParseInt(pageSize, 10, 32)
	robboGroupsCore, countRowsInt64, err := r.UseCase.GetAllRobboGroups(
		int(pageInt32),
		int(pageSizeInt32),
	)
	if err != nil {
		return
	}
	countRows = int(countRowsInt64)
	for _, robboGroupCore := range robboGroupsCore {
		var robboGroupTemp models.RobboGroupHTTP
		robboGroupTemp.FromCore(robboGroupCore)
		robboGroups = append(robboGroups, &robboGroupTemp)
	}
	return
}

func (r *RobboGroupDelegateImpl) SearchRobboGroupByName(name, page, pageSize string) (
	robboGroups []*models.RobboGroupHTTP,
	countRows int,
	err error,
) {
	pageInt32, _ := strconv.ParseInt(page, 10, 32)
	pageSizeInt32, _ := strconv.ParseInt(pageSize, 10, 32)

	robboGroupsCore, countRowsInt64, err := r.UseCase.SearchRobboGroupsByTitle(
		name,
		int(pageInt32),
		int(pageSizeInt32),
	)
	if err != nil {
		return
	}
	countRows = int(countRowsInt64)
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

func (r *RobboGroupDelegateImpl) CreateRobboGroup(robboGroup *models.RobboGroupHTTP) (newRobboGroup models.RobboGroupHTTP, err error) {
	robboGroupCore := robboGroup.ToCore()
	newRobboGroupCore, err := r.UseCase.CreateRobboGroup(robboGroupCore)
	if err != nil {
		log.Println(err)
		return
	}
	newRobboGroup.FromCore(newRobboGroupCore)
	return
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

func (r *RobboGroupDelegateImpl) GetRobboGroupsByUnitAdminId(unitAdminId, page, pageSize string) (
	robboGroups []*models.RobboGroupHTTP,
	countRows int,
	err error,
) {
	pageInt32, _ := strconv.ParseInt(page, 10, 32)
	pageSizeInt32, _ := strconv.ParseInt(pageSize, 10, 32)
	robboGroupsCore, countRowsInt64, err := r.UseCase.GetRobboGroupsByUnitAdminId(
		unitAdminId,
		int(pageInt32),
		int(pageSizeInt32),
	)
	if err != nil {
		return
	}
	countRows = int(countRowsInt64)
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
