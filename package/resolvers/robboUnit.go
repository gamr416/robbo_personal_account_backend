package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateRobboUnit is the resolver for the CreateRobboUnit field.
func (r *mutationResolver) CreateRobboUnit(ctx context.Context, input models.NewRobboUnit) (*models.RobboUnitHTTP, error) {
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

	robboUnitHttp := models.RobboUnitHTTP{
		Name: input.Name,
		City: input.City,
	}

	newRobboUnit, createRobboUnitErr := r.robboUnitsDelegate.CreateRobboUnit(&robboUnitHttp)
	if createRobboUnitErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return &newRobboUnit, nil
}

// UpdateRobboUnit is the resolver for the UpdateRobboUnit field.
func (r *mutationResolver) UpdateRobboUnit(ctx context.Context, input models.UpdateRobboUnit) (*models.RobboUnitHTTP, error) {
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

	updateRobboUnitInput := &models.RobboUnitHTTP{
		ID:   input.ID,
		Name: input.Name,
		City: input.City,
	}

	robboUnitUpdated, updateRobboUnitErr := r.robboUnitsDelegate.UpdateRobboUnit(updateRobboUnitInput)
	if updateRobboUnitErr != nil {
		err := errors.New("baq request")
		return nil, err
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
func (r *queryResolver) GetRobboUnitByID(ctx context.Context, id string) (*models.RobboUnitHTTP, error) {
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

	robboUnitsHttp, getRobboUnitIdErr := r.robboUnitsDelegate.GetRobboUnitById(id)
	if getRobboUnitIdErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return &robboUnitsHttp, nil
}

// GetAllRobboUnits is the resolver for the GetAllRobboUnits field.
func (r *queryResolver) GetAllRobboUnits(ctx context.Context) ([]*models.RobboUnitHTTP, error) {
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
	robboUnitsHttp, err := r.robboUnitsDelegate.GetAllRobboUnit()
	return robboUnitsHttp, err
}

// GetRobboUnitsByUnitAdminID is the resolver for the GetRobboUnitsByUnitAdminId field.
func (r *queryResolver) GetRobboUnitsByUnitAdminID(ctx context.Context) ([]*models.RobboUnitHTTP, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	robboUnitsHttp, getRobboUnitsByUnitAdminIdErr := r.robboUnitsDelegate.GetRobboUnitsByUnitAdminId(userId)
	if getRobboUnitsByUnitAdminIdErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return robboUnitsHttp, nil
}
