package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// GetRobboGroupByID is the resolver for the GetRobboGroupById field.
func (r *queryResolver) GetRobboGroupByID(ctx context.Context, id string) (*models.RobboGroupHTTP, error) {
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
	robboGroupsHttp, err := r.robboGroupDelegate.GetRobboGroupById(id)
	return &robboGroupsHttp, err
}

// GetRobboGroupsByTeacherID is the resolver for the GetRobboGroupsByTeacherId field.
func (r *queryResolver) GetRobboGroupsByTeacherID(ctx context.Context, teacherID string) ([]*models.RobboGroupHTTP, error) {
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
	robboGroupsHttp, err := r.robboGroupDelegate.GetRobboGroupsByTeacherId(teacherID)
	return robboGroupsHttp, err
}

// GetRobboGroupsByRobboUnitID is the resolver for the GetRobboGroupsByRobboUnitId field.
func (r *queryResolver) GetRobboGroupsByRobboUnitID(ctx context.Context, robboUnitID string) ([]*models.RobboGroupHTTP, error) {
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
	robboGroupsHttp, err := r.robboGroupDelegate.GetRobboGroupsByRobboUnitId(robboUnitID)
	return robboGroupsHttp, err
}

// SearchGroupsByName is the resolver for the SearchGroupsByName field.
func (r *queryResolver) SearchGroupsByName(ctx context.Context, name string) ([]*models.RobboGroupHTTP, error) {
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
	robboGroupsHttp, err := r.robboGroupDelegate.SearchRobboGroupByName(name)
	return robboGroupsHttp, err
}
