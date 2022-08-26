package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup"
	"go.uber.org/fx"
)

type RobboGroupDelegateImpl struct {
	UseCase robboGroup.UseCase
}

func (r *RobboGroupDelegateImpl) CreateRobboGroup(robboGroup *models.RobboGroupHttp) (robboGroupId string, err error) {
	robboGroupCore := robboGroup.ToCore()
	return r.UseCase.CreateRobboGroup(robboGroupCore)
}

func (r *RobboGroupDelegateImpl) DeleteRobboGroup(robboGroupId string) (err error) {
	return r.UseCase.DeleteRobboGroup(robboGroupId)
}

func (r *RobboGroupDelegateImpl) GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroups []*models.RobboGroupHttp, err error) {
	robboGroupsCore, err := r.UseCase.GetRobboGroupsByRobboUnitId(robboUnitId)
	if err != nil {
		return
	}
	for _, robboGroupCore := range robboGroupsCore {
		var robboGroupTemp models.RobboGroupHttp
		robboGroupTemp.FromCore(robboGroupCore)
		robboGroups = append(robboGroups, &robboGroupTemp)
	}
	return
}

func (r *RobboGroupDelegateImpl) GetRobboGroupById(robboGroupId string) (robboGroup models.RobboGroupHttp, err error) {
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
