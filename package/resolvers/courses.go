package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

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
func (r *queryResolver) GetCoursesByUser(ctx context.Context) (models.CourseResult, error) {
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
func (r *queryResolver) GetAllPublicCourses(ctx context.Context, pageNumber string) (models.CourseResult, error) {
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
func (r *queryResolver) GetEnrollments(ctx context.Context, username string) (models.EnrollmentResult, error) {
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
