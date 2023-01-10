package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateCourseRelationGroup is the resolver for the CreateCourseRelationGroup field.
func (r *mutationResolver) CreateCourseRelationGroup(ctx context.Context, input models.NewCourseRelationGroup) (models.CourseRelationResult, error) {
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

	courseRelation := &models.CourseRelationHTTP{
		CourseID: input.CourseID,
		ObjectID: input.RobboGroupID,
	}
	newCourseRelation, createCourseRelationGroupErr := r.coursesDelegate.CreateCourseRelationGroup(courseRelation)
	if createCourseRelationGroupErr != nil {
		err := createCourseRelationGroupErr
		return &models.Error{Message: err.Error()}, err
	}
	return newCourseRelation, nil
}

// CreateCourseRelationUnit is the resolver for the CreateCourseRelationUnit field.
func (r *mutationResolver) CreateCourseRelationUnit(ctx context.Context, input models.NewCourseRelationUnit) (models.CourseRelationResult, error) {
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

	courseRelation := &models.CourseRelationHTTP{
		CourseID: input.CourseID,
		ObjectID: input.RobboUnitID,
	}
	newCourseRelation, createCourseRelationUnitErr := r.coursesDelegate.CreateCourseRelationUnit(courseRelation)
	if createCourseRelationUnitErr != nil {
		err := createCourseRelationUnitErr
		return &models.Error{Message: err.Error()}, err
	}
	return newCourseRelation, nil
}

// CreateCourseRelationRole is the resolver for the CreateCourseRelationRole field.
func (r *mutationResolver) CreateCourseRelationRole(ctx context.Context, input models.NewCourseRelationRole) (models.CourseRelationResult, error) {
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

	courseRelation := &models.CourseRelationHTTP{
		CourseID: input.CourseID,
		ObjectID: input.RoleID,
	}
	newCourseRelation, createCourseRelationRoleErr := r.coursesDelegate.CreateCourseRelationRole(courseRelation)
	if createCourseRelationRoleErr != nil {
		err := createCourseRelationRoleErr
		return &models.Error{Message: err.Error()}, err
	}
	return newCourseRelation, nil
}

// DeleteCourseRelationByID is the resolver for the DeleteCourseRelationById field.
func (r *mutationResolver) DeleteCourseRelationByID(ctx context.Context, courseRelationID string) (*models.DeletedCourseRelation, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.DeletedCourseRelation{CourseRelationID: ""}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := identityErr
		return &models.DeletedCourseRelation{CourseRelationID: ""}, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.DeletedCourseRelation{CourseRelationID: ""}, err
	}

	_, deleteCourseRelationByIdErr := r.coursesDelegate.DeleteCourseRelationById(courseRelationID)
	if deleteCourseRelationByIdErr != nil {
		err := errors.New("baq request")
		return &models.DeletedCourseRelation{CourseRelationID: ""}, err
	}
	return &models.DeletedCourseRelation{CourseRelationID: courseRelationID}, nil
}

// GetCourseRelationsByCourseID is the resolver for the GetCourseRelationsByCourseId field.
func (r *queryResolver) GetCourseRelationsByCourseID(ctx context.Context, courseID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsByCourseIdErr := r.coursesDelegate.GetCourseRelationsByCourseId(courseID)
	if getCourseRelationsByCourseIdErr != nil {
		err := getCourseRelationsByCourseIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsByRobboUnitID is the resolver for the GetCourseRelationsByRobboUnitId field.
func (r *queryResolver) GetCourseRelationsByRobboUnitID(ctx context.Context, robboUnitID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsByRobboUnitIdErr := r.coursesDelegate.GetCourseRelationsByRobboUnitId(robboUnitID)
	if getCourseRelationsByRobboUnitIdErr != nil {
		err := getCourseRelationsByRobboUnitIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsByRobboGroupID is the resolver for the GetCourseRelationsByRobboGroupId field.
func (r *queryResolver) GetCourseRelationsByRobboGroupID(ctx context.Context, robboGroupID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsByRobboGroupIdErr := r.coursesDelegate.GetCourseRelationsByRobboGroupId(robboGroupID)
	if getCourseRelationsByRobboGroupIdErr != nil {
		err := getCourseRelationsByRobboGroupIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsByRoleID is the resolver for the GetCourseRelationsByRoleId field.
func (r *queryResolver) GetCourseRelationsByRoleID(ctx context.Context, roleID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsByRoleIdErr := r.coursesDelegate.GetCourseRelationsByRoleId(roleID)
	if getCourseRelationsByRoleIdErr != nil {
		err := getCourseRelationsByRoleIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsUnits is the resolver for the GetCourseRelationsUnits field.
func (r *queryResolver) GetCourseRelationsUnits(ctx context.Context) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsUnitsErr := r.coursesDelegate.GetCourseRelationsUnits()
	if getCourseRelationsUnitsErr != nil {
		err := getCourseRelationsUnitsErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsGroups is the resolver for the GetCourseRelationsGroups field.
func (r *queryResolver) GetCourseRelationsGroups(ctx context.Context) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsGroupsErr := r.coursesDelegate.GetCourseRelationsGroups()
	if getCourseRelationsGroupsErr != nil {
		err := getCourseRelationsGroupsErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsRoles is the resolver for the GetCourseRelationsRoles field.
func (r *queryResolver) GetCourseRelationsRoles(ctx context.Context) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsRolesErr := r.coursesDelegate.GetCourseRelationsRoles()
	if getCourseRelationsRolesErr != nil {
		err := getCourseRelationsRolesErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseContent is the resolver for the GetCourseContent field.
func (r *queryResolver) GetCourseContent(ctx context.Context, courseID string) (models.CourseResult, error) {
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
	allowedRoles := []models.Role{models.Student, models.FreeListener, models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	courseHttp, getCourseContentErr := r.coursesDelegate.GetCourseContent(courseID)
	if getCourseContentErr != nil {
		err := getCourseContentErr
		return &models.Error{Message: err.Error()}, err
	}
	return courseHttp, nil
}

// GetCoursesByUser is the resolver for the GetCoursesByUser field.
func (r *queryResolver) GetCoursesByUser(ctx context.Context) (models.CoursesResult, error) {
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
	allowedRoles := []models.Role{models.Student, models.Parent, models.FreeListener, models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	courses, getCoursesByUserErr := r.coursesDelegate.GetCoursesByUser()
	if getCoursesByUserErr != nil {
		err := getCoursesByUserErr
		return &models.Error{Message: err.Error()}, err
	}
	return courses, nil
}

// GetAllPublicCourses is the resolver for the GetAllPublicCourses field.
func (r *queryResolver) GetAllPublicCourses(ctx context.Context, pageNumber string) (models.CoursesResult, error) {
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
	allowedRoles := []models.Role{models.Student, models.Parent, models.FreeListener, models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	courses, getAllPublicCoursesErr := r.coursesDelegate.GetAllPublicCourses(pageNumber)
	if getAllPublicCoursesErr != nil {
		err := getAllPublicCoursesErr
		return &models.Error{Message: err.Error()}, err
	}
	return courses, nil
}

// GetEnrollments is the resolver for the GetEnrollments field.
func (r *queryResolver) GetEnrollments(ctx context.Context, username string) (models.EnrollmentsResult, error) {
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
	enrollments, getEnrollmentsErr := r.coursesDelegate.GetEnrollments(username)
	if getEnrollmentsErr != nil {
		err := getEnrollmentsErr
		return &models.Error{Message: err.Error()}, err
	}
	return enrollments, nil
}
