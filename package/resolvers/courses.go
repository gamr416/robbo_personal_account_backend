package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// GetCourseContent is the resolver for the GetCourseContent field.
func (r *queryResolver) GetCourseContent(ctx context.Context, courseID string) (*models.CourseHTTP, error) {
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
	allowedRoles := []models.Role{models.Student, models.FreeListener, models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	courseHttp, getCourseContentErr := r.coursesDelegate.GetCourseContent(courseID)
	if getCourseContentErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return courseHttp, nil
}

// GetCoursesByUser is the resolver for the GetCoursesByUser field.
func (r *queryResolver) GetCoursesByUser(ctx context.Context) (*models.CoursesListHTTP, error) {
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
	allowedRoles := []models.Role{models.Student, models.Parent, models.FreeListener, models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	coursesListHttp, getCoursesByUserErr := r.coursesDelegate.GetCoursesByUser()
	if getCoursesByUserErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return coursesListHttp, nil
}

// GetAllPublicCourses is the resolver for the GetAllPublicCourses field.
func (r *queryResolver) GetAllPublicCourses(ctx context.Context, pageNumber string) (*models.CoursesListHTTP, error) {
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
	allowedRoles := []models.Role{models.Student, models.Parent, models.FreeListener, models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	coursesListHttp, getAllPublicCoursesErr := r.coursesDelegate.GetAllPublicCourses(pageNumber)
	if getAllPublicCoursesErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return coursesListHttp, nil
}

// GetEnrollments is the resolver for the GetEnrollments field.
func (r *queryResolver) GetEnrollments(ctx context.Context, username string) (*models.EnrollmentsListHTTP, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	enrollmentListHttp, GetEnrollmentsErr := r.coursesDelegate.GetEnrollments(username)
	if GetEnrollmentsErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return enrollmentListHttp, nil
}
