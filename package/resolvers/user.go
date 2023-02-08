package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/skinnykaen/robbo_student_personal_account.git/graph/generated"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// UpdateSuperAdmin is the resolver for the updateSuperAdmin field.
func (r *mutationResolver) UpdateSuperAdmin(ctx context.Context, input models.UpdateProfileInput) (models.SuperAdminResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}
	updateSuperAdminInput := &models.SuperAdminHTTP{
		UserHTTP: &models.UserHTTP{
			ID:         input.ID,
			Email:      input.Email,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
			Role:       5,
		},
	}
	superAdminUpdated, updateSuperAdminErr := r.usersDelegate.UpdateSuperAdmin(updateSuperAdminInput)
	if updateSuperAdminErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: updateSuperAdminErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return superAdminUpdated, nil
}

// GetSuperAdminByID is the resolver for the GetSuperAdminById field.
func (r *queryResolver) GetSuperAdminByID(ctx context.Context, superAdminID string) (models.SuperAdminResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	superAdmin, getSuperAdminByIdErr := r.usersDelegate.GetSuperAdminById(superAdminID)
	return nil, &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: getSuperAdminByIdErr.Error(),
		Extensions: map[string]interface{}{
			"code": "500",
		},
	}
	return &superAdmin, nil
}

// GetUser is the resolver for the GetUser field.
func (r *queryResolver) GetUser(ctx context.Context, peekUserID *string, peekUserRole *int) (models.GetUserResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	var userId string
	var userRole models.Role
	if utils.UseString(peekUserID) == "" || peekUserID == nil {
		userId = ginContext.Value("user_id").(string)
		userRole = ginContext.Value("user_role").(models.Role)
	} else {
		userId = *peekUserID
		userRole = models.Role(*peekUserRole)
	}
	switch userRole {
	case models.Student:
		student, getStudentErr := r.usersDelegate.GetStudentById(userId)
		if getStudentErr != nil {
			return nil, getStudentErr
		}
		return student, nil
	case models.Teacher:
		teacher, getTeacherErr := r.usersDelegate.GetTeacherById(userId)
		if getTeacherErr != nil {
			return nil, getTeacherErr
		}
		return teacher, nil
	case models.Parent:
		parent, getParentErr := r.usersDelegate.GetParentById(userId)
		if getParentErr != nil {
			return nil, getParentErr
		}
		return parent, nil
	case models.UnitAdmin:
		unitAdmin, getUnitAdminErr := r.usersDelegate.GetUnitAdminById(userId)
		if getUnitAdminErr != nil {
			return nil, getUnitAdminErr
		}
		return unitAdmin, nil
	case models.SuperAdmin:
		superAdmin, getSuperAdminErr := r.usersDelegate.GetSuperAdminById(userId)
		if getSuperAdminErr != nil {
			return nil, getSuperAdminErr
		}
		return superAdmin, nil
	default:
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "internal server error",
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
