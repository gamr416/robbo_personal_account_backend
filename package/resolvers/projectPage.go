package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// GetProjectPageByID is the resolver for the GetProjectPageById field.
func (r *queryResolver) GetProjectPageByID(ctx context.Context, projectPageID string) (*models.ProjectPageHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, _, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	projectPageHttp, getProjectPageByIDErr := r.projectPageDelegate.GetProjectPageById(projectPageID)
	if getProjectPageByIDErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return &projectPageHttp, nil
}

// GetAllProjectPagesByUserID is the resolver for the GetAllProjectPagesByUserID field.
func (r *queryResolver) GetAllProjectPagesByUserID(ctx context.Context, userID string) ([]*models.ProjectPageHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, _, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	projectPageListHttp, getAllProjectPagesErr := r.projectPageDelegate.GetAllProjectPagesByUserId(userID)
	if getAllProjectPagesErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return projectPageListHttp, nil
}

// GetAllProjectPagesByAccessToken is the resolver for the GetAllProjectPagesByAccessToken field.
func (r *queryResolver) GetAllProjectPagesByAccessToken(ctx context.Context) ([]*models.ProjectPageHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	userId, _, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	projectPageListHttp, getAllProjectPagesErr := r.projectPageDelegate.GetAllProjectPagesByUserId(userId)
	if getAllProjectPagesErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return projectPageListHttp, nil
}
