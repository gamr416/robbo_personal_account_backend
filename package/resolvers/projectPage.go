package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateProjectPage is the resolver for the CreateProjectPage field.
func (r *mutationResolver) CreateProjectPage(ctx context.Context) (string, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return "", err
	}
	userID, _, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return "", err
	}
	projectId, createProjectPageErr := r.projectPageDelegate.CreateProjectPage(userID)
	if createProjectPageErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return projectId, nil
}

// UpdateProjectPage is the resolver for the UpdateProjectPage field.
func (r *mutationResolver) UpdateProjectPage(ctx context.Context, input models.UpdateProjectPage) (*models.ProjectPageHTTP, error) {
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
	updateProjectPageInput := &models.ProjectPageHTTP{
		ProjectID:   input.ProjectID,
		Instruction: input.Instruction,
		Notes:       input.Notes,
		Preview:     input.Preview,
		LinkScratch: input.LinkScratch,
		Title:       input.Title,
		IsShared:    input.IsShared,
	}

	updateProjectPageErr := r.projectPageDelegate.UpdateProjectPage(updateProjectPageInput)
	if updateProjectPageErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return updateProjectPageInput, nil
}

// DeleteProjectPage is the resolver for the DeleteProjectPage field.
func (r *mutationResolver) DeleteProjectPage(ctx context.Context, projectID string) (string, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return "", err
	}
	_, _, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return "", err
	}
	deleteProjectPageErr := r.projectPageDelegate.DeleteProjectPage(projectID)
	if deleteProjectPageErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return projectID, nil
}

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

// GetAllProjectPageByUserID is the resolver for the GetAllProjectPageByUserID field.
func (r *queryResolver) GetAllProjectPageByUserID(ctx context.Context) ([]*models.ProjectPageHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	userID, _, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
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
