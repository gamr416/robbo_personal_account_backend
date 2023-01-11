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

// CreateCourseRelationStudent is the resolver for the CreateCourseRelationStudent field.
func (r *mutationResolver) CreateCourseRelationStudent(ctx context.Context, input models.NewCourseRelationStudent) (models.CourseRelationResult, error) {
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
		ObjectID: input.StudentID,
	}
	newCourseRelation, createCourseRelationStudentErr := r.coursesDelegate.CreateCourseRelationStudent(courseRelation)
	if createCourseRelationStudentErr != nil {
		err := createCourseRelationStudentErr
		return &models.Error{Message: err.Error()}, err
	}
	return newCourseRelation, nil
}

// CreateCourseRelationTeacher is the resolver for the CreateCourseRelationTeacher field.
func (r *mutationResolver) CreateCourseRelationTeacher(ctx context.Context, input models.NewCourseRelationTeacher) (models.CourseRelationResult, error) {
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
		ObjectID: input.TeacherID,
	}
	newCourseRelation, createCourseRelationTeacherErr := r.coursesDelegate.CreateCourseRelationTeacher(courseRelation)
	if createCourseRelationTeacherErr != nil {
		err := createCourseRelationTeacherErr
		return &models.Error{Message: err.Error()}, err
	}
	return newCourseRelation, nil
}

// CreateCourseRelationUnitAdmin is the resolver for the CreateCourseRelationUnitAdmin field.
func (r *mutationResolver) CreateCourseRelationUnitAdmin(ctx context.Context, input models.NewCourseRelationUnitAdmin) (models.CourseRelationResult, error) {
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
		ObjectID: input.UnitAdminID,
	}
	newCourseRelation, createCourseRelationUnitAdminErr := r.coursesDelegate.CreateCourseRelationUnitAdmin(courseRelation)
	if createCourseRelationUnitAdminErr != nil {
		err := createCourseRelationUnitAdminErr
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

// GetCourseRelationsByStudentID is the resolver for the GetCourseRelationsByStudentId field.
func (r *queryResolver) GetCourseRelationsByStudentID(ctx context.Context, studentID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsByStudentIdErr := r.coursesDelegate.GetCourseRelationsByStudentId(studentID)
	if getCourseRelationsByStudentIdErr != nil {
		err := getCourseRelationsByStudentIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsByTeacherID is the resolver for the GetCourseRelationsByTeacherId field.
func (r *queryResolver) GetCourseRelationsByTeacherID(ctx context.Context, teacherID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsByTeacherIdErr := r.coursesDelegate.GetCourseRelationsByTeacherId(teacherID)
	if getCourseRelationsByTeacherIdErr != nil {
		err := getCourseRelationsByTeacherIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsByUnitAdminID is the resolver for the GetCourseRelationsByUnitAdminId field.
func (r *queryResolver) GetCourseRelationsByUnitAdminID(ctx context.Context, unitAdminID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsByUnitAdminIdErr := r.coursesDelegate.GetCourseRelationsByUnitAdminId(unitAdminID)
	if getCourseRelationsByUnitAdminIdErr != nil {
		err := getCourseRelationsByUnitAdminIdErr
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

// GetCourseRelationsStudents is the resolver for the GetCourseRelationsStudents field.
func (r *queryResolver) GetCourseRelationsStudents(ctx context.Context) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsStudentsErr := r.coursesDelegate.GetCourseRelationsStudents()
	if getCourseRelationsStudentsErr != nil {
		err := getCourseRelationsStudentsErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsTeachers is the resolver for the GetCourseRelationsTeachers field.
func (r *queryResolver) GetCourseRelationsTeachers(ctx context.Context) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsTeachersErr := r.coursesDelegate.GetCourseRelationsTeachers()
	if getCourseRelationsTeachersErr != nil {
		err := getCourseRelationsTeachersErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsUnitAdmins is the resolver for the GetCourseRelationsUnitAdmins field.
func (r *queryResolver) GetCourseRelationsUnitAdmins(ctx context.Context) (models.CourseRelationsResult, error) {
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
	courseRelations, getCourseRelationsUnitAdminsErr := r.coursesDelegate.GetCourseRelationsUnitAdmins()
	if getCourseRelationsUnitAdminsErr != nil {
		err := getCourseRelationsUnitAdminsErr
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
