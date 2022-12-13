package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateRobboGroup is the resolver for the CreateRobboGroup field.
func (r *mutationResolver) CreateRobboGroup(ctx context.Context, input models.NewRobboGroup) (models.RobboGroupResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := identityErr
		return &models.Error{Message: err.Error()}, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}

	robboGroupHttp := models.RobboGroupHTTP{
		Name:        input.Name,
		RobboUnitID: input.RobboUnitID,
	}

	newRobboGroup, createRobboGroupErr := r.robboGroupDelegate.CreateRobboGroup(&robboGroupHttp)
	if createRobboGroupErr != nil {
		err := createRobboGroupErr
		return &models.Error{Message: err.Error()}, err
	}
	return &newRobboGroup, nil
}

// UpdateRobboGroup is the resolver for the UpdateRobboGroup field.
func (r *mutationResolver) UpdateRobboGroup(ctx context.Context, input models.UpdateRobboGroup) (models.RobboGroupResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := identityErr
		return &models.Error{Message: err.Error()}, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}

	updateRobboGroupInput := &models.RobboGroupHTTP{
		ID:          input.ID,
		Name:        input.Name,
		RobboUnitID: input.RobboUnitID,
	}

	robboGroupUpdated, updateRobboGroupErr := r.robboGroupDelegate.UpdateRobboGroup(updateRobboGroupInput)
	if updateRobboGroupErr != nil {
		err := updateRobboGroupErr
		return &models.Error{Message: err.Error()}, err
	}
	return &robboGroupUpdated, nil
}

// DeleteRobboGroup is the resolver for the DeleteRobboGroup field.
func (r *mutationResolver) DeleteRobboGroup(ctx context.Context, robboGroupID string) (*models.DeletedRobboGroup, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.DeletedRobboGroup{RobboGroupID: ""}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return &models.DeletedRobboGroup{RobboGroupID: ""}, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.DeletedRobboGroup{RobboGroupID: ""}, err
	}

	deleteRobboGroupErr := r.robboGroupDelegate.DeleteRobboGroup(robboGroupID)
	if deleteRobboGroupErr != nil {
		err := errors.New("baq request")
		return &models.DeletedRobboGroup{RobboGroupID: ""}, err
	}
	return &models.DeletedRobboGroup{RobboGroupID: robboGroupID}, nil
}

// GetRobboGroupByID is the resolver for the GetRobboGroupById field.
func (r *queryResolver) GetRobboGroupByID(ctx context.Context, id string) (models.RobboGroupResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := identityErr
		return &models.Error{Message: err.Error()}, err
	}
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	robboGroup, getRobboGroupByIdErr := r.robboGroupDelegate.GetRobboGroupById(id)
	if getRobboGroupByIdErr != nil {
		err := getRobboGroupByIdErr
		return nil, err
	}
	return &robboGroup, nil
}

// GetRobboGroupsByTeacherID is the resolver for the GetRobboGroupsByTeacherId field.
func (r *queryResolver) GetRobboGroupsByTeacherID(ctx context.Context, teacherID string) (models.RobboGroupResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := identityErr
		return &models.Error{Message: err.Error()}, err
	}
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	robboGroups, getRobboGroupsByTeacherIdErr := r.robboGroupDelegate.GetRobboGroupsByTeacherId(teacherID)
	if getRobboGroupsByTeacherIdErr != nil {
		err := getRobboGroupsByTeacherIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
	}, nil
}

// GetRobboGroupsByRobboUnitID is the resolver for the GetRobboGroupsByRobboUnitId field.
func (r *queryResolver) GetRobboGroupsByRobboUnitID(ctx context.Context, robboUnitID string) (models.RobboGroupResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := identityErr
		return &models.Error{Message: err.Error()}, err
	}
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	robboGroups, getRobboGroupsByRobboUnitIdErr := r.robboGroupDelegate.GetRobboGroupsByRobboUnitId(robboUnitID)
	if getRobboGroupsByRobboUnitIdErr != nil {
		err := getRobboGroupsByRobboUnitIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
	}, nil
}

// GetRobboGroupsByUnitAdminID is the resolver for the GetRobboGroupsByUnitAdminID field.
func (r *queryResolver) GetRobboGroupsByUnitAdminID(ctx context.Context, unitAdminID string) (models.RobboGroupResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := identityErr
		return &models.Error{Message: err.Error()}, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	robboGroups, _, getRobboGroupsByUnitAdminIdErr := r.robboGroupDelegate.GetRobboGroupsByUnitAdminId(unitAdminID, "0", "0")
	if getRobboGroupsByUnitAdminIdErr != nil {
		err := getRobboGroupsByUnitAdminIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
	}, nil
}

// GetAllRobboGroups is the resolver for the GetAllRobboGroups field.
func (r *queryResolver) GetAllRobboGroups(ctx context.Context) (models.RobboGroupResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := identityErr
		return &models.Error{Message: err.Error()}, err
	}
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}

	robboGroups, getAllRobboGroups := r.robboGroupDelegate.GetAllRobboGroups()
	if getAllRobboGroups != nil {
		err := getAllRobboGroups
		return &models.Error{Message: err.Error()}, err
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
	}, nil
}

// GetRobboGroupsByAccessToken is the resolver for the GetRobboGroupsByAccessToken field.
func (r *queryResolver) GetRobboGroupsByAccessToken(ctx context.Context) (models.RobboGroupResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	userId, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := identityErr
		return &models.Error{Message: err.Error()}, err
	}
	//todo: add admins
	allowedRoles := []models.Role{models.Teacher}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	robboGroups, getRobboGroupsByTeacherIdErr := r.robboGroupDelegate.GetRobboGroupsByTeacherId(userId)
	if getRobboGroupsByTeacherIdErr != nil {
		err := getRobboGroupsByTeacherIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
	}, nil
}

// SearchGroupsByName is the resolver for the SearchGroupsByName field.
func (r *queryResolver) SearchGroupsByName(ctx context.Context, name string) (models.RobboGroupResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := identityErr
		return &models.Error{Message: err.Error()}, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	robboGroups, searchRobboGroupByNameErr := r.robboGroupDelegate.SearchRobboGroupByName(name)
	if searchRobboGroupByNameErr != nil {
		err := searchRobboGroupByNameErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
	}, nil
}
