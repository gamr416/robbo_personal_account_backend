package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateProjectPage is the resolver for the CreateProjectPage field.
func (r *mutationResolver) CreateProjectPage(ctx context.Context) (models.ProjectPageResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userId := ginContext.Value("user_id").(string)
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	newProjectPage, createProjectPageErr := r.projectPageDelegate.CreateProjectPage(userId)
	if createProjectPageErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: createProjectPageErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &newProjectPage, nil
}

// UpdateProjectPage is the resolver for the UpdateProjectPage field.
func (r *mutationResolver) UpdateProjectPage(ctx context.Context, input models.UpdateProjectPage) (models.ProjectPageResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	updateProjectPageInput := &models.ProjectPageHTTP{
		ProjectID:     input.ProjectID,
		ProjectPageID: input.ProjectPageID,
		Instruction:   input.Instruction,
		Notes:         input.Notes,
		Title:         input.Title,
		IsShared:      input.IsShared,
	}

	updateProjectPage, updateProjectPageErr := r.projectPageDelegate.UpdateProjectPage(updateProjectPageInput)
	if updateProjectPageErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: updateProjectPageErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return updateProjectPage, nil
}

// DeleteProjectPage is the resolver for the DeleteProjectPage field.
func (r *mutationResolver) DeleteProjectPage(ctx context.Context, projectID string) (*models.DeletedProjectPage, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	deleteProjectPageErr := r.projectPageDelegate.DeleteProjectPage(projectID)
	if deleteProjectPageErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: deleteProjectPageErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.DeletedProjectPage{ProjectPageID: projectID}, nil
}

// GetProjectPageByID is the resolver for the GetProjectPageById field.
func (r *queryResolver) GetProjectPageByID(ctx context.Context, projectPageID string) (models.ProjectPageResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	projectPageHttp, getProjectPageByIdErr := r.projectPageDelegate.GetProjectPageById(projectPageID)
	if getProjectPageByIdErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getProjectPageByIdErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &projectPageHttp, nil
}

// GetAllProjectPagesByUserID is the resolver for the GetAllProjectPagesByUserID field.
func (r *queryResolver) GetAllProjectPagesByUserID(ctx context.Context, userID string) (models.ProjectPageResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	projectPages, getAllProjectPagesErr := r.projectPageDelegate.GetAllProjectPagesByUserId(userID)
	if getAllProjectPagesErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getAllProjectPagesErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.ProjectPageHTTPList{
		ProjectPages: projectPages,
	}, nil
}

// GetAllProjectPagesByAccessToken is the resolver for the GetAllProjectPagesByAccessToken field.
func (r *queryResolver) GetAllProjectPagesByAccessToken(ctx context.Context) (models.ProjectPageResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userId := ginContext.Value("user_id").(string)
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.Student, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	projectPages, getAllProjectPagesErr := r.projectPageDelegate.GetAllProjectPagesByUserId(userId)
	if getAllProjectPagesErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getAllProjectPagesErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.ProjectPageHTTPList{
		ProjectPages: projectPages,
	}, nil
}
