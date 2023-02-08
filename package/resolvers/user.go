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
		err := updateSuperAdminErr
		return &models.Error{Message: err.Error()}, err
	}
	return superAdminUpdated, nil
}

// GetPairsStudentParentsByAccessToken is the resolver for the GetPairsStudentParentsByAccessToken field.
func (r *queryResolver) GetPairsStudentParentsByAccessToken(ctx context.Context) (models.PairsStudentParentsResult, error) {
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
	allowedRoles := []models.Role{models.Teacher}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	pairsStudentParents, getPairsStudentParentsByAccessTokenErr := r.usersDelegate.GetPairsStudentParentsByTeacherId(userId)
	if getPairsStudentParentsByAccessTokenErr != nil {
		err := getPairsStudentParentsByAccessTokenErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.StudentParentsHTTPList{
		PairsStudentParents: pairsStudentParents,
	}, nil
}

// GetIndividualStudentsByTeacherID is the resolver for the GetIndividualStudentsByTeacherId field.
func (r *queryResolver) GetIndividualStudentsByTeacherID(ctx context.Context, teacherID string) (models.StudentsResult, error) {
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
	allowedRoles := []models.Role{models.Teacher, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	students, getIndividualStudentsByTeacherIdErr := r.usersDelegate.GetIndividualStudentsByTeacherId(teacherID)
	if getIndividualStudentsByTeacherIdErr != nil {
		err := getIndividualStudentsByTeacherIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.StudentHTTPList{
		Students: students,
	}, nil
}

// GetStudentsByParentID is the resolver for the GetStudentsByParentId field.
func (r *queryResolver) GetStudentsByParentID(ctx context.Context, parentID string) (models.StudentsResult, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin, models.Parent}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	students, getStudentByParentIdErr := r.usersDelegate.GetStudentByParentId(parentID)
	if getStudentByParentIdErr != nil {
		err := getStudentByParentIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.StudentHTTPList{
		Students: students,
	}, nil
}

// GetStudentByID is the resolver for the GetStudentById field.
func (r *queryResolver) GetStudentByID(ctx context.Context, studentID string) (models.StudentResult, error) {
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
	allowedRoles := []models.Role{models.Parent, models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	student, getStudentByIdErr := r.usersDelegate.GetStudentById(studentID)
	if getStudentByIdErr != nil {
		err := getStudentByIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return student, nil
}

// GetStudentsByRobboGroup is the resolver for the GetStudentsByRobboGroup field.
func (r *queryResolver) GetStudentsByRobboGroup(ctx context.Context, robboGroupID string) (models.StudentsResult, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin, models.Teacher}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	students, getStudentByRobboGroupErr := r.usersDelegate.GetStudentsByRobboGroupId(robboGroupID)
	if getStudentByRobboGroupErr != nil {
		err := getStudentByRobboGroupErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.StudentHTTPList{
		Students: students,
	}, nil
}

// GetStudentsByRobboUnitID is the resolver for the GetStudentsByRobboUnitId field.
func (r *queryResolver) GetStudentsByRobboUnitID(ctx context.Context, robboUnitID string) (models.StudentsResult, error) {
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
	students, getStudentsByRobboUnitIdErr := r.usersDelegate.GetStudentsByRobboUnitId(robboUnitID)
	if getStudentsByRobboUnitIdErr != nil {
		err := getStudentsByRobboUnitIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.StudentHTTPList{
		Students: students,
	}, nil
}

// SearchStudentsByEmail is the resolver for the SearchStudentsByEmail field.
func (r *queryResolver) SearchStudentsByEmail(ctx context.Context, email string, page string, pageSize string) (models.StudentsResult, error) {
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
	students, countRows, searchStudentByEmailErr := r.usersDelegate.SearchStudentByEmail(email, page, pageSize)
	if searchStudentByEmailErr != nil {
		err := searchStudentByEmailErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.StudentHTTPList{
		Students:  students,
		CountRows: countRows,
	}, nil
}

// GetAllTeachers is the resolver for the GetAllTeachers field.
func (r *queryResolver) GetAllTeachers(ctx context.Context, page string, pageSize string) (models.TeachersResult, error) {
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
	teachers, countRows, getAllTeachersErr := r.usersDelegate.GetAllTeachers(page, pageSize)
	if getAllTeachersErr != nil {
		err := getAllTeachersErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.TeacherHTTPList{
		Teachers:  teachers,
		CountRows: countRows,
	}, nil
}

// GetTeacherByID is the resolver for the GetTeacherById field.
func (r *queryResolver) GetTeacherByID(ctx context.Context, teacherID string) (models.TeacherResult, error) {
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
	teacher, getTeacherByIdErr := r.usersDelegate.GetTeacherById(teacherID)
	if getTeacherByIdErr != nil {
		err := getTeacherByIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return teacher, nil
}

// GetTeachersByRobboGroupID is the resolver for the GetTeachersByRobboGroupId field.
func (r *queryResolver) GetTeachersByRobboGroupID(ctx context.Context, robboGroupID string) (models.TeachersResult, error) {
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
	allowedRoles := []models.Role{models.SuperAdmin, models.UnitAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	teachers, getTeacherByRobboGroupIdErr := r.usersDelegate.GetTeacherByRobboGroupId(robboGroupID)
	if getTeacherByRobboGroupIdErr != nil {
		err := getTeacherByRobboGroupIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.TeacherHTTPList{
		Teachers: teachers,
	}, nil
}

// SearchTeachersByEmail is the resolver for the SearchTeachersByEmail field.
func (r *queryResolver) SearchTeachersByEmail(ctx context.Context, email string, page string, pageSize string) (models.TeachersResult, error) {
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
	teachers, countRows, searchTeachersByEmailErr := r.usersDelegate.SearchTeacherByEmail(email, page, pageSize)
	if searchTeachersByEmailErr != nil {
		err := searchTeachersByEmailErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.TeacherHTTPList{
		Teachers:  teachers,
		CountRows: countRows,
	}, nil
}

// GetAllParents is the resolver for the GetAllParents field.
func (r *queryResolver) GetAllParents(ctx context.Context, page string, pageSize string) (models.ParentsResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles)
	if accessErr != nil {
		err := accessErr

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


// SearchUnitAdminsByEmail is the resolver for the SearchUnitAdminsByEmail field.
func (r *queryResolver) SearchUnitAdminsByEmail(ctx context.Context, email string, page string, pageSize string) (models.UnitAdminsResult, error) {
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
	unitAdmins, countRows, searchUnitAdminByEmailErr := r.usersDelegate.SearchUnitAdminByEmail(email, page, pageSize)
	if searchUnitAdminByEmailErr != nil {
		err := searchUnitAdminByEmailErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.UnitAdminHTTPList{
		UnitAdmins: unitAdmins,
		CountRows:  countRows,
	}, nil
}

// GetUser is the resolver for the GetUser field.
func (r *queryResolver) GetUser(ctx context.Context, peekUserID *string, peekUserRole *int) (models.GetUserResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "internal server error",
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
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
