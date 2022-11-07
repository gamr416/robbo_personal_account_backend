package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// GetRobboUnitByID is the resolver for the GetRobboUnitById field.
func (r *queryResolver) GetRobboUnitByID(ctx context.Context, id string) (*models.RobboUnitHTTP, error) {
	ginContext, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(err)
	userId, userRole, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		return nil, identityErr
	}
	fmt.Println(userId)
	fmt.Println(userRole)
	fmt.Println(identityErr)
	robboUnitsHttp, err := r.robboUnitsDelegate.GetRobboUnitById(id)
	return &robboUnitsHttp, err
}

// GetAllRobboUnits is the resolver for the GetAllRobboUnits field.
func (r *queryResolver) GetAllRobboUnits(ctx context.Context) ([]*models.RobboUnitHTTP, error) {
	ginContext, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(err)
	userId, userRole, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		return nil, identityErr
	}
	fmt.Println(userId)
	fmt.Println(userRole)
	fmt.Println(identityErr)
	robboUnitsHttp, err := r.robboUnitsDelegate.GetAllRobboUnit()
	return robboUnitsHttp, err
}

// GetRobboUnitsByUnitAdminID is the resolver for the GetRobboUnitsByUnitAdminId field.
func (r *queryResolver) GetRobboUnitsByUnitAdminID(ctx context.Context, unitAdminID string) ([]*models.RobboUnitHTTP, error) {
	ginContext, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(err)
	userId, userRole, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		return nil, identityErr
	}
	fmt.Println(userId)
	fmt.Println(userRole)
	fmt.Println(identityErr)
	robboUnitsHttp, err := r.robboUnitsDelegate.GetRobboUnitsByUnitAdminId(unitAdminID)
	return robboUnitsHttp, err
}
