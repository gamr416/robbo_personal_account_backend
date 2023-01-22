package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateCourseRelationGroup is the resolver for the CreateCourseRelationGroup field.
func (r *mutationResolver) CreateAccessCourseRelationRobboGroup(ctx context.Context, input models.NewAccessCourseRelationRobboGroup) (models.CourseRelationResult, error) {
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
	newCourseRelation, createAccessCourseRelationRobboGroupErr := r.coursesDelegate.CreateAccessCourseRelationRobboGroup(courseRelation)
	if createAccessCourseRelationRobboGroupErr != nil {
		err := createAccessCourseRelationRobboGroupErr
		return &models.Error{Message: err.Error()}, err
	}
	return newCourseRelation, nil
}

// CreateCourseRelationUnit is the resolver for the CreateCourseRelationUnit field.
func (r *mutationResolver) CreateAccessCourseRelationRobboUnit(ctx context.Context, input models.NewAccessCourseRelationRobboUnit) (models.CourseRelationResult, error) {
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
		return &models.Error{Message: err.Error()}, err
	}

	courseRelation := &models.CourseRelationHTTP{
		CourseID: input.CourseID,
		ObjectID: input.RobboUnitID,
	}
	newCourseRelation, createAccessCourseRelationRobboUnitErr := r.coursesDelegate.CreateAccessCourseRelationRobboUnit(courseRelation)
	if createAccessCourseRelationRobboUnitErr != nil {
		err := createAccessCourseRelationRobboUnitErr
		return &models.Error{Message: err.Error()}, err
	}
	return newCourseRelation, nil
}

// CreateCourseRelationStudent is the resolver for the CreateCourseRelationStudent field.
func (r *mutationResolver) CreateAccessCourseRelationStudent(ctx context.Context, input models.NewAccessCourseRelationStudent) (models.CourseRelationResult, error) {
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
	newCourseRelation, createAccessCourseRelationStudentErr := r.coursesDelegate.CreateAccessCourseRelationStudent(courseRelation)
	if createAccessCourseRelationStudentErr != nil {
		err := createAccessCourseRelationStudentErr
		return &models.Error{Message: err.Error()}, err
	}
	return newCourseRelation, nil
}

// CreateCourseRelationTeacher is the resolver for the CreateCourseRelationTeacher field.
func (r *mutationResolver) CreateAccessCourseRelationTeacher(ctx context.Context, input models.NewAccessCourseRelationTeacher) (models.CourseRelationResult, error) {
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
	newCourseRelation, createAccessCourseRelationTeacherErr := r.coursesDelegate.CreateAccessCourseRelationTeacher(courseRelation)
	if createAccessCourseRelationTeacherErr != nil {
		err := createAccessCourseRelationTeacherErr
		return &models.Error{Message: err.Error()}, err
	}
	return newCourseRelation, nil
}

// CreateCourseRelationUnitAdmin is the resolver for the CreateCourseRelationUnitAdmin field.
func (r *mutationResolver) CreateAccessCourseRelationUnitAdmin(ctx context.Context, input models.NewAccessCourseRelationUnitAdmin) (models.CourseRelationResult, error) {
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
	newCourseRelation, createAccessCourseRelationUnitAdminErr := r.coursesDelegate.CreateAccessCourseRelationUnitAdmin(courseRelation)
	if createAccessCourseRelationUnitAdminErr != nil {
		err := createAccessCourseRelationUnitAdminErr
		return &models.Error{Message: err.Error()}, err
	}
	return newCourseRelation, nil
}

// DeleteCourseRelationByID is the resolver for the DeleteCourseRelationById field.
func (r *mutationResolver) DeleteAccessCourseRelationByID(ctx context.Context, courseRelationID string) (*models.DeletedCourseRelation, error) {
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

	_, deleteCourseRelationByIdErr := r.coursesDelegate.DeleteAccessCourseRelationById(courseRelationID)
	if deleteCourseRelationByIdErr != nil {
		err := errors.New("baq request")
		return &models.DeletedCourseRelation{CourseRelationID: ""}, err
	}
	return &models.DeletedCourseRelation{CourseRelationID: courseRelationID}, nil
}

// GetStudentsAdmittedToTheCourse is the resolver for the GetStudentsAdmittedToTheCourse field.
func (r *queryResolver) GetStudentsAdmittedToTheCourse(ctx context.Context, courseID string, page *string, pageSize *string) (models.StudentsResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	//userId := ginContext.Value("user_id").(string)
	userRole := ginContext.Value("user_role").(models.Role)

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	students, getStudentsAdmittedToTheCourseErr := r.coursesDelegate.GetStudentsAdmittedToTheCourse(courseID, page, pageSize)
	if getStudentsAdmittedToTheCourseErr != nil {
		err := getStudentsAdmittedToTheCourseErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.StudentHTTPList{Students: students}, nil
}

// GetUnitAdminsAdmittedToTheCourse is the resolver for the GetUnitAdminsAdmittedToTheCourse field.
func (r *queryResolver) GetUnitAdminsAdmittedToTheCourse(ctx context.Context, courseID string, page *string, pageSize *string) (models.UnitAdminsResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	//userId := ginContext.Value("user_id").(string)
	userRole := ginContext.Value("user_role").(models.Role)

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	unitAdmins, getUnitAdminsAdmittedToTheCourseErr := r.coursesDelegate.GetUnitAdminsAdmittedToTheCourse(courseID, page, pageSize)
	if getUnitAdminsAdmittedToTheCourseErr != nil {
		err := getUnitAdminsAdmittedToTheCourseErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.UnitAdminHTTPList{UnitAdmins: unitAdmins}, nil
}

// GetTeachersAdmittedToTheCourse is the resolver for the GetTeachersAdmittedToTheCourse field.
func (r *queryResolver) GetTeachersAdmittedToTheCourse(ctx context.Context, courseID string, page *string, pageSize *string) (models.TeachersResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	//userId := ginContext.Value("user_id").(string)
	userRole := ginContext.Value("user_role").(models.Role)

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	teachers, getTeachersAdmittedToTheCourseErr := r.coursesDelegate.GetTeachersAdmittedToTheCourse(courseID, page, pageSize)
	if getTeachersAdmittedToTheCourseErr != nil {
		err := getTeachersAdmittedToTheCourseErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.TeacherHTTPList{Teachers: teachers}, nil
}

// GetRobboGroupsAdmittedToTheCourse is the resolver for the GetRobboGroupsAdmittedToTheCourse field.
func (r *queryResolver) GetRobboGroupsAdmittedToTheCourse(ctx context.Context, courseID string, page *string, pageSize *string) (models.RobboGroupsResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	//userId := ginContext.Value("user_id").(string)
	userRole := ginContext.Value("user_role").(models.Role)

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	robboGroups, getRobboGroupsAdmittedToTheCourseErr := r.coursesDelegate.GetRobboGroupsAdmittedToTheCourse(courseID, page, pageSize)
	if getRobboGroupsAdmittedToTheCourseErr != nil {
		err := getRobboGroupsAdmittedToTheCourseErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.RobboGroupHTTPList{RobboGroups: robboGroups}, nil
}

// GetRobboUnitsAdmittedToTheCourse is the resolver for the GetRobboUnitsAdmittedToTheCourse field.
func (r *queryResolver) GetRobboUnitsAdmittedToTheCourse(ctx context.Context, courseID string, page *string, pageSize *string) (models.RobboUnitsResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	//userId := ginContext.Value("user_id").(string)
	userRole := ginContext.Value("user_role").(models.Role)

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	robboUnits, getRobboUnitsAdmittedToTheCourseErr := r.coursesDelegate.GetRobboUnitsAdmittedToTheCourse(courseID, page, pageSize)
	if getRobboUnitsAdmittedToTheCourseErr != nil {
		err := getRobboUnitsAdmittedToTheCourseErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.RobboUnitHTTPList{RobboUnits: robboUnits}, nil
}

// GetCourseRelationsByCourseID is the resolver for the GetCourseRelationsByCourseId field.
func (r *queryResolver) GetAccessCourseRelationsByCourseID(ctx context.Context, courseID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getAccessCourseRelationsByCourseIdErr := r.coursesDelegate.GetAccessCourseRelationsByCourseId(courseID)
	if getAccessCourseRelationsByCourseIdErr != nil {
		err := getAccessCourseRelationsByCourseIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsByRobboUnitID is the resolver for the GetCourseRelationsByRobboUnitId field.
func (r *queryResolver) GetAccessCourseRelationsByRobboUnitID(ctx context.Context, robboUnitID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getAccessCourseRelationsByRobboUnitIdErr := r.coursesDelegate.GetAccessCourseRelationsByRobboUnitId(robboUnitID)
	if getAccessCourseRelationsByRobboUnitIdErr != nil {
		err := getAccessCourseRelationsByRobboUnitIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsByRobboGroupID is the resolver for the GetCourseRelationsByRobboGroupId field.
func (r *queryResolver) GetAccessCourseRelationsByRobboGroupID(ctx context.Context, robboGroupID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getAccessCourseRelationsByRobboGroupIdErr := r.coursesDelegate.GetAccessCourseRelationsByRobboGroupId(robboGroupID)
	if getAccessCourseRelationsByRobboGroupIdErr != nil {
		err := getAccessCourseRelationsByRobboGroupIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsByStudentID is the resolver for the GetCourseRelationsByStudentId field.
func (r *queryResolver) GetAccessCourseRelationsByStudentID(ctx context.Context, studentID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getAccessCourseRelationsByStudentIdErr := r.coursesDelegate.GetAccessCourseRelationsByStudentId(studentID)
	if getAccessCourseRelationsByStudentIdErr != nil {
		err := getAccessCourseRelationsByStudentIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsByTeacherID is the resolver for the GetCourseRelationsByTeacherId field.
func (r *queryResolver) GetAccessCourseRelationsByTeacherID(ctx context.Context, teacherID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getAccessCourseRelationsByTeacherIdErr := r.coursesDelegate.GetAccessCourseRelationsByTeacherId(teacherID)
	if getAccessCourseRelationsByTeacherIdErr != nil {
		err := getAccessCourseRelationsByTeacherIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsByUnitAdminID is the resolver for the GetCourseRelationsByUnitAdminId field.
func (r *queryResolver) GetAccessCourseRelationsByUnitAdminID(ctx context.Context, unitAdminID string) (models.CourseRelationsResult, error) {
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
	courseRelations, getAccessCourseRelationsByUnitAdminIdErr := r.coursesDelegate.GetAccessCourseRelationsByUnitAdminId(unitAdminID)
	if getAccessCourseRelationsByUnitAdminIdErr != nil {
		err := getAccessCourseRelationsByUnitAdminIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsUnits is the resolver for the GetCourseRelationsUnits field.
func (r *queryResolver) GetAccessCourseRelationsRobboUnits(ctx context.Context) (models.CourseRelationsResult, error) {
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
	courseRelations, getAccessCourseRelationsRobboUnitsErr := r.coursesDelegate.GetAccessCourseRelationsRobboUnits()
	if getAccessCourseRelationsRobboUnitsErr != nil {
		err := getAccessCourseRelationsRobboUnitsErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsGroups is the resolver for the GetCourseRelationsGroups field.
func (r *queryResolver) GetAccessCourseRelationsRobboGroups(ctx context.Context) (models.CourseRelationsResult, error) {
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
	courseRelations, getAccessCourseRelationsRobboGroupsErr := r.coursesDelegate.GetAccessCourseRelationsRobboGroups()
	if getAccessCourseRelationsRobboGroupsErr != nil {
		err := getAccessCourseRelationsRobboGroupsErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsStudents is the resolver for the GetCourseRelationsStudents field.
func (r *queryResolver) GetAccessCourseRelationsStudents(ctx context.Context) (models.CourseRelationsResult, error) {
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
	courseRelations, getAccessCourseRelationsStudentsErr := r.coursesDelegate.GetAccessCourseRelationsStudents()
	if getAccessCourseRelationsStudentsErr != nil {
		err := getAccessCourseRelationsStudentsErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsTeachers is the resolver for the GetCourseRelationsTeachers field.
func (r *queryResolver) GetAccessCourseRelationsTeachers(ctx context.Context) (models.CourseRelationsResult, error) {
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
	courseRelations, getAccessCourseRelationsTeachersErr := r.coursesDelegate.GetAccessCourseRelationsTeachers()
	if getAccessCourseRelationsTeachersErr != nil {
		err := getAccessCourseRelationsTeachersErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.CourseRelationHTTPList{
		courseRelations,
	}, nil
}

// GetCourseRelationsUnitAdmins is the resolver for the GetCourseRelationsUnitAdmins field.
func (r *queryResolver) GetAccessCourseRelationsUnitAdmins(ctx context.Context) (models.CourseRelationsResult, error) {
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
	courseRelations, getAccessCourseRelationsUnitAdminsErr := r.coursesDelegate.GetAccessCourseRelationsUnitAdmins()
	if getAccessCourseRelationsUnitAdminsErr != nil {
		err := getAccessCourseRelationsUnitAdminsErr
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
	userId := ginContext.Value("user_id").(string)
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{
		models.Student,
		models.FreeListener,
		models.Teacher,
		models.UnitAdmin,
		models.SuperAdmin,
	}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	courses, getCoursesByUserErr := r.coursesDelegate.GetCoursesByUser(userId, userRole)
	if getCoursesByUserErr != nil {
		err := getCoursesByUserErr
		return &models.Error{Message: err.Error()}, err
	}
	return courses, nil
}

// GetCoursesByRobboUnitID is the resolver for the GetCoursesByRobboUnitId field.
func (r *queryResolver) GetCoursesByRobboUnitID(ctx context.Context, robboUnitID string, page *string, pageSize *string) (models.CoursesResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{
		models.UnitAdmin,
		models.SuperAdmin,
	}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	courses, getCoursesByUserErr := r.coursesDelegate.GetCoursesByRobboUnitId(robboUnitID, *page, *pageSize)
	if getCoursesByUserErr != nil {
		err := getCoursesByUserErr
		return &models.Error{Message: err.Error()}, err
	}
	return courses, nil
}

// GetCoursesByRobboGroupID is the resolver for the GetCoursesByRobboGroupId field.
func (r *queryResolver) GetCoursesByRobboGroupID(ctx context.Context, robboGroupID string, page *string, pageSize *string) (models.CoursesResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{
		models.Teacher,
		models.UnitAdmin,
		models.SuperAdmin,
	}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles)
	if accessErr != nil {
		err := accessErr
		return &models.Error{Message: err.Error()}, err
	}
	courses, getCoursesByUserErr := r.coursesDelegate.GetCoursesByRobboGroupId(robboGroupID, *page, *pageSize)
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
