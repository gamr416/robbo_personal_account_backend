package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateRobboUnit is the resolver for the CreateRobboUnit field.
func (r *mutationResolver) CreateRobboUnit(ctx context.Context, input models.NewRobboUnit) (models.RobboUnitResult, error) {
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

	robboUnitHttp := models.RobboUnitHTTP{
		Name: input.Name,
		City: input.City,
	}

	newRobboUnit, createRobboUnitErr := r.robboUnitsDelegate.CreateRobboUnit(&robboUnitHttp)
	if createRobboUnitErr != nil {
		err := createRobboUnitErr
		return &models.Error{Message: err.Error()}, err
	}
	return &newRobboUnit, nil
}

// UpdateRobboUnit is the resolver for the UpdateRobboUnit field.
func (r *mutationResolver) UpdateRobboUnit(ctx context.Context, input models.UpdateRobboUnit) (models.RobboUnitResult, error) {
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

	updateRobboUnitInput := &models.RobboUnitHTTP{
		ID:   input.ID,
		Name: input.Name,
		City: input.City,
	}

	robboUnitUpdated, updateRobboUnitErr := r.robboUnitsDelegate.UpdateRobboUnit(updateRobboUnitInput)
	if updateRobboUnitErr != nil {
		err := updateRobboUnitErr
		return &models.Error{Message: err.Error()}, err
	}
	return &robboUnitUpdated, nil
}

// DeleteRobboUnit is the resolver for the DeleteRobboUnit field.
func (r *mutationResolver) DeleteRobboUnit(ctx context.Context, robboUnitID string) (*models.DeletedRobboUnit, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.DeletedRobboUnit{RobboUnitID: ""}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return &models.DeletedRobboUnit{RobboUnitID: ""}, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.DeletedRobboUnit{RobboUnitID: ""}, err
	}

	deleteRobboUnitErr := r.robboUnitsDelegate.DeleteRobboUnit(robboUnitID)
	if deleteRobboUnitErr != nil {
		err := errors.New("baq request")
		return &models.DeletedRobboUnit{RobboUnitID: ""}, err
	}
	return &models.DeletedRobboUnit{RobboUnitID: robboUnitID}, nil
}

// GetRobboUnitByID is the resolver for the GetRobboUnitById field.
func (r *queryResolver) GetRobboUnitByID(ctx context.Context, id string) (models.RobboUnitResult, error) {
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

	robboUnit, getRobboUnitIdErr := r.robboUnitsDelegate.GetRobboUnitById(id)
	if getRobboUnitIdErr != nil {
		err := getRobboUnitIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &robboUnit, nil
}

// GetAllRobboUnits is the resolver for the GetAllRobboUnits field.
func (r *queryResolver) GetAllRobboUnits(ctx context.Context, page string, pageSize string) (models.RobboUnitsResult, error) {
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
	robboUnits, countRows, getAllRobboUnitErr := r.robboUnitsDelegate.GetAllRobboUnit(page, pageSize)
	if getAllRobboUnitErr != nil {
		err := getAllRobboUnitErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.RobboUnitHTTPList{
		RobboUnits: robboUnits,
		CountRows:  countRows,
	}, nil
}

// GetRobboUnitsByUnitAdminID is the resolver for the GetRobboUnitsByUnitAdminId field.
func (r *queryResolver) GetRobboUnitsByUnitAdminID(ctx context.Context, unitAdminID string) (models.RobboUnitsResult, error) {
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
	robboUnits, countRows, getRobboUnitsByUnitAdminIdErr := r.robboUnitsDelegate.GetRobboUnitsByUnitAdminId(unitAdminID, "0", "0")
	if getRobboUnitsByUnitAdminIdErr != nil {
		err := getRobboUnitsByUnitAdminIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.RobboUnitHTTPList{
		RobboUnits: robboUnits,
		CountRows:  countRows,
	}, nil
}

// GetRobboUnitsByAccessToken is the resolver for the GetRobboUnitsByAccessToken field.
func (r *queryResolver) GetRobboUnitsByAccessToken(ctx context.Context, page string, pageSize string) (models.RobboUnitsResult, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	robboUnits, countRows, getRobboUnitsByUnitAdminIdErr := r.robboUnitsDelegate.GetRobboUnitsByUnitAdminId(userId, page, pageSize)
	if getRobboUnitsByUnitAdminIdErr != nil {
		err := getRobboUnitsByUnitAdminIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.RobboUnitHTTPList{
		RobboUnits: robboUnits,
		CountRows:  countRows,
	}, nil
}

// SearchRobboUnitsByName is the resolver for the SearchRobboUnitsByName field.
func (r *queryResolver) SearchRobboUnitsByName(ctx context.Context, name string) (models.RobboUnitsResult, error) {
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
	robboUnits, searchTeachersByEmailErr := r.robboUnitsDelegate.SearchRobboUnitsByName(name)
	if searchTeachersByEmailErr != nil {
		err := searchTeachersByEmailErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.RobboUnitHTTPList{
		RobboUnits: robboUnits,
	}, nil
}
