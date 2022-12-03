package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateRobboGroup is the resolver for the CreateRobboGroup field.
func (r *mutationResolver) CreateRobboGroup(ctx context.Context, input models.NewRobboGroup) (string, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return "", err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return "", err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return "", err
	}

	robboGroupHttp := models.RobboGroupHTTP{
		Name:        input.Name,
		RobboUnitID: input.RobboUnitID,
	}

	robboGroupHttpId, createRobboGroupErr := r.robboGroupDelegate.CreateRobboGroup(&robboGroupHttp)
	if createRobboGroupErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return robboGroupHttpId, nil
}

// UpdateRobboGroup is the resolver for the UpdateRobboGroup field.
func (r *mutationResolver) UpdateRobboGroup(ctx context.Context, input models.UpdateRobboGroup) (*models.RobboGroupHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}

	updateRobboGroupInput := &models.RobboGroupHTTP{
		ID:          input.ID,
		Name:        input.Name,
		RobboUnitID: input.RobboUnitID,
	}

	deleteRobboGroupErr := r.robboGroupDelegate.UpdateRobboGroup(updateRobboGroupInput)
	if deleteRobboGroupErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return updateRobboGroupInput, nil
}

// DeleteRobboGroup is the resolver for the DeleteRobboGroup field.
func (r *mutationResolver) DeleteRobboGroup(ctx context.Context, robboGroupID string) (string, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return "", err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return "", err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return "", err
	}

	deleteRobboGroupErr := r.robboGroupDelegate.DeleteRobboGroup(robboGroupID)
	if deleteRobboGroupErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return robboGroupID, nil
}

// GetRobboGroupByID is the resolver for the GetRobboGroupById field.
func (r *queryResolver) GetRobboGroupByID(ctx context.Context, id string) (*models.RobboGroupHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	robboGroupsHttp, getRobboGroupByIdErr := r.robboGroupDelegate.GetRobboGroupById(id)
	if getRobboGroupByIdErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return &robboGroupsHttp, nil
}

// GetRobboGroupsByTeacherID is the resolver for the GetRobboGroupsByTeacherId field.
func (r *queryResolver) GetRobboGroupsByTeacherID(ctx context.Context, teacherID string) ([]*models.RobboGroupHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, role, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	robboGroupsHttp, getRobboGroupsByTeacherIdErr := r.robboGroupDelegate.GetRobboGroupsByTeacherId(teacherID)
	if getRobboGroupsByTeacherIdErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return robboGroupsHttp, nil
}

// GetRobboGroupsByRobboUnitID is the resolver for the GetRobboGroupsByRobboUnitId field.
func (r *queryResolver) GetRobboGroupsByRobboUnitID(ctx context.Context, robboUnitID string) ([]*models.RobboGroupHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, role, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	robboGroupsHttp, getRobboGroupsByRobboUnitIdErr := r.robboGroupDelegate.GetRobboGroupsByRobboUnitId(robboUnitID)
	if getRobboGroupsByRobboUnitIdErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return robboGroupsHttp, nil
}

// GetRobboGroupsByUnitAdminID is the resolver for the GetRobboGroupsByUnitAdminID field.
func (r *queryResolver) GetRobboGroupsByUnitAdminID(ctx context.Context, unitAdminID string) ([]*models.RobboGroupHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, role, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	robboGroupsHttp, getRobboGroupsByUnitAdminIDErr := r.robboGroupDelegate.GetRobboGroupsByUnitAdminId(unitAdminID)
	if getRobboGroupsByUnitAdminIDErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return robboGroupsHttp, nil
}

// GetAllRobboGroups is the resolver for the GetAllRobboGroups field.
func (r *queryResolver) GetAllRobboGroups(ctx context.Context) ([]*models.RobboGroupHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, role, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}

	robboGroupsHttp, err := r.robboGroupDelegate.GetAllRobboGroups()
	return robboGroupsHttp, err
}

// GetRobboGroupsByAccessToken is the resolver for the GetRobboGroupsByAccessToken field.
func (r *queryResolver) GetRobboGroupsByAccessToken(ctx context.Context) ([]*models.RobboGroupHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	userId, role, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	//todo: add admins
	allowedRoles := []models.Role{models.Teacher}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	robboGroupsHttp, getRobboGroupsByTeacherIdErr := r.robboGroupDelegate.GetRobboGroupsByTeacherId(userId)
	if getRobboGroupsByTeacherIdErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return robboGroupsHttp, nil
}

// SearchGroupsByName is the resolver for the SearchGroupsByName field.
func (r *queryResolver) SearchGroupsByName(ctx context.Context, name string) ([]*models.RobboGroupHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, role, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	robboGroupsHttp, searchRobboGroupByNameErr := r.robboGroupDelegate.SearchRobboGroupByName(name)
	if searchRobboGroupByNameErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return robboGroupsHttp, nil
}
