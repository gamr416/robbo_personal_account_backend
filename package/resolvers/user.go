package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/skinnykaen/robbo_student_personal_account.git/graph/generated"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateStudent is the resolver for the createStudent field.
func (r *mutationResolver) CreateStudent(ctx context.Context, input models.NewStudent) (models.StudentResult, error) {
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
	studentInput := models.StudentHTTP{
		UserHTTP: &models.UserHTTP{
			Email:      input.Email,
			Password:   input.Password,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
			Role:       0,
		},
	}

	newStudent, createStudentErr := r.usersDelegate.CreateStudent(&studentInput, input.ParentID)
	if createStudentErr != nil {
		err := createStudentErr
		return &models.Error{Message: err.Error()}, err
	}
	return newStudent, nil
}

// UpdateStudent is the resolver for the updateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, input models.UpdateStudentInput) (models.StudentResult, error) {
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
	updateStudentInput := &models.StudentHTTP{
		UserHTTP: &models.UserHTTP{
			ID:         input.StudentHTTP.UserHTTP.ID,
			Email:      input.StudentHTTP.UserHTTP.Email,
			Firstname:  input.StudentHTTP.UserHTTP.Firstname,
			Lastname:   input.StudentHTTP.UserHTTP.Lastname,
			Middlename: input.StudentHTTP.UserHTTP.Middlename,
			Nickname:   input.StudentHTTP.UserHTTP.Nickname,
		},
	}
	studentUpdated, updateStudentErr := r.usersDelegate.UpdateStudent(updateStudentInput)
	if updateStudentErr != nil {
		err := updateStudentErr
		return &models.Error{Message: err.Error()}, err
	}
	return studentUpdated, nil
}

// DeleteStudent is the resolver for the deleteStudent field.
func (r *mutationResolver) DeleteStudent(ctx context.Context, studentID string) (*models.DeletedStudent, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.DeletedStudent{StudentID: ""}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return &models.DeletedStudent{StudentID: ""}, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.DeletedStudent{StudentID: ""}, err
	}

	deleteStudentErr := r.usersDelegate.DeleteStudent(studentID)
	if deleteStudentErr != nil {
		err := errors.New("baq request")
		return &models.DeletedStudent{StudentID: ""}, err
	}
	return &models.DeletedStudent{StudentID: studentID}, nil
}

// SetRobboGroupIDForStudent is the resolver for the setRobboGroupIdForStudent field.
func (r *mutationResolver) SetRobboGroupIDForStudent(ctx context.Context, studentID string, robboGroupID string, robboUnitID string) (string, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return "", err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return "", err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return "", err
	}

	addStudentToRobboGroupErr := r.usersDelegate.AddStudentToRobboGroup(studentID, robboGroupID, robboUnitID)
	if addStudentToRobboGroupErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return "", nil
}

// CreateStudentTeacherRelation is the resolver for the createStudentTeacherRelation field.
func (r *mutationResolver) CreateStudentTeacherRelation(ctx context.Context, studentID string, teacherID string) (models.StudentResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return &models.Error{Message: "status unauthorized"}, err
	}
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.Error{Message: "no access"}, err
	}
	student, createStudentTeacherRelationErr := r.usersDelegate.CreateStudentTeacherRelation(studentID, teacherID)
	if createStudentTeacherRelationErr != nil {
		err := errors.New("baq request")
		return &models.Error{Message: "baq request"}, err
	}
	return student, nil
}

// DeleteStudentTeacherRelation is the resolver for the deleteStudentTeacherRelation field.
func (r *mutationResolver) DeleteStudentTeacherRelation(ctx context.Context, studentID string, teacherID string) (models.StudentResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return &models.Error{Message: "status unauthorized"}, err
	}
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.Error{Message: "no access"}, err
	}
	student, createStudentTeacherRelationErr := r.usersDelegate.DeleteStudentTeacherRelation(studentID, teacherID)
	if createStudentTeacherRelationErr != nil {
		err := errors.New("baq request")
		return &models.Error{Message: "baq request"}, err
	}
	return student, nil
}

// CreateTeacher is the resolver for the createTeacher field.
func (r *mutationResolver) CreateTeacher(ctx context.Context, input models.NewTeacher) (models.TeacherResult, error) {
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
	teacherInput := models.TeacherHTTP{
		UserHTTP: &models.UserHTTP{
			Email:      input.Email,
			Password:   input.Password,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
			Role:       1,
		},
	}

	newTeacher, createTeacherErr := r.usersDelegate.CreateTeacher(&teacherInput)
	if createTeacherErr != nil {
		err := createTeacherErr
		return &models.Error{Message: err.Error()}, err
	}

	return &newTeacher, nil
}

// UpdateTeacher is the resolver for the updateTeacher field.
func (r *mutationResolver) UpdateTeacher(ctx context.Context, input models.UpdateTeacherInput) (models.TeacherResult, error) {
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
	updateTeacherInput := &models.TeacherHTTP{
		UserHTTP: &models.UserHTTP{
			ID:         input.TeacherHTTP.UserHTTP.ID,
			Email:      input.TeacherHTTP.UserHTTP.Email,
			Firstname:  input.TeacherHTTP.UserHTTP.Firstname,
			Lastname:   input.TeacherHTTP.UserHTTP.Lastname,
			Middlename: input.TeacherHTTP.UserHTTP.Middlename,
			Nickname:   input.TeacherHTTP.UserHTTP.Nickname,
		},
	}
	teacherUpdated, updateTeacherErr := r.usersDelegate.UpdateTeacher(updateTeacherInput)
	if updateTeacherErr != nil {
		err := updateTeacherErr
		return &models.Error{Message: err.Error()}, err
	}
	return &teacherUpdated, nil
}

// DeleteTeacher is the resolver for the deleteTeacher field.
func (r *mutationResolver) DeleteTeacher(ctx context.Context, teacherID string) (*models.DeletedTeacher, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.DeletedTeacher{TeacherID: ""}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return &models.DeletedTeacher{TeacherID: ""}, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.DeletedTeacher{TeacherID: ""}, err
	}
	deleteTeacherErr := r.usersDelegate.DeleteTeacher(teacherID)
	if deleteTeacherErr != nil {
		err := errors.New("baq request")
		return &models.DeletedTeacher{TeacherID: ""}, err
	}
	return &models.DeletedTeacher{TeacherID: teacherID}, nil
}

// CreateParent is the resolver for the createParent field.
func (r *mutationResolver) CreateParent(ctx context.Context, input models.NewParent) (models.ParentResult, error) {
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
	parentInput := models.ParentHTTP{
		UserHTTP: &models.UserHTTP{
			Email:      input.Email,
			Password:   input.Password,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
			Role:       2,
		},
		Children: nil,
	}
	newParent, createParentErr := r.usersDelegate.CreateParent(&parentInput)
	if createParentErr != nil {
		err := createParentErr
		return &models.Error{Message: err.Error()}, err
	}
	return newParent, nil
}

// AddChildToParent is the resolver for the addChildToParent field.
func (r *mutationResolver) AddChildToParent(ctx context.Context, parentID string, childID string) (string, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return "", err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return "", err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return "", err
	}
	createRelationErr := r.usersDelegate.CreateRelation(parentID, childID)
	if createRelationErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return "", nil
}

// UpdateParent is the resolver for the updateParent field.
func (r *mutationResolver) UpdateParent(ctx context.Context, input models.UpdateParentInput) (models.ParentResult, error) {
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
	updateParentInput := &models.ParentHTTP{
		UserHTTP: &models.UserHTTP{
			ID:         input.ParentHTTP.UserHTTP.ID,
			Email:      input.ParentHTTP.UserHTTP.Email,
			Firstname:  input.ParentHTTP.UserHTTP.Firstname,
			Lastname:   input.ParentHTTP.UserHTTP.Lastname,
			Middlename: input.ParentHTTP.UserHTTP.Middlename,
			Nickname:   input.ParentHTTP.UserHTTP.Nickname,
		},
	}
	parentUpdated, updateParentErr := r.usersDelegate.UpdateParent(updateParentInput)
	if updateParentErr != nil {
		err := updateParentErr
		return &models.Error{Message: err.Error()}, err
	}
	return parentUpdated, nil
}

// DeleteParent is the resolver for the deleteParent field.
func (r *mutationResolver) DeleteParent(ctx context.Context, parentID string) (*models.DeletedParent, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.DeletedParent{ParentID: ""}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return &models.DeletedParent{ParentID: ""}, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.DeletedParent{ParentID: ""}, err
	}
	deleteParentErr := r.usersDelegate.DeleteParent(parentID)
	if deleteParentErr != nil {
		err := errors.New("baq request")
		return &models.DeletedParent{ParentID: ""}, err
	}
	return &models.DeletedParent{ParentID: parentID}, nil
}

// CreateUnitAdmin is the resolver for the createUnitAdmin field.
func (r *mutationResolver) CreateUnitAdmin(ctx context.Context, input models.NewUnitAdmin) (models.UnitAdminResult, error) {
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
	unitAdminInput := models.UnitAdminHTTP{
		UserHTTP: &models.UserHTTP{
			Email:      input.Email,
			Password:   input.Password,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
			Role:       4,
		},
	}
	newUnitAdmin, createUnitAdminErr := r.usersDelegate.CreateUnitAdmin(&unitAdminInput)
	if createUnitAdminErr != nil {
		err := createUnitAdminErr
		return &models.Error{Message: err.Error()}, err
	}
	return newUnitAdmin, nil
}

// UpdateUnitAdmin is the resolver for the updateUnitAdmin field.
func (r *mutationResolver) UpdateUnitAdmin(ctx context.Context, input models.UpdateUnitAdminInput) (models.UnitAdminResult, error) {
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
	updateUnitAdminInput := &models.UnitAdminHTTP{
		UserHTTP: &models.UserHTTP{
			ID:         input.UnitAdminHTTP.UserHTTP.ID,
			Email:      input.UnitAdminHTTP.UserHTTP.Email,
			Firstname:  input.UnitAdminHTTP.UserHTTP.Firstname,
			Lastname:   input.UnitAdminHTTP.UserHTTP.Lastname,
			Middlename: input.UnitAdminHTTP.UserHTTP.Middlename,
			Nickname:   input.UnitAdminHTTP.UserHTTP.Nickname,
		},
	}
	unitAdminUpdated, updateUnitAdminErr := r.usersDelegate.UpdateUnitAdmin(updateUnitAdminInput)
	if updateUnitAdminErr != nil {
		err := updateUnitAdminErr
		return &models.Error{Message: err.Error()}, err
	}
	return unitAdminUpdated, nil
}

// DeleteUnitAdmin is the resolver for the deleteUnitAdmin field.
func (r *mutationResolver) DeleteUnitAdmin(ctx context.Context, unitAdminID string) (*models.DeletedUnitAdmin, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.DeletedUnitAdmin{UnitAdminID: ""}, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return &models.DeletedUnitAdmin{UnitAdminID: ""}, err
	}
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.DeletedUnitAdmin{UnitAdminID: ""}, err
	}
	deleteUnitAdminErr := r.usersDelegate.DeleteUnitAdmin(unitAdminID)
	if deleteUnitAdminErr != nil {
		err := errors.New("baq request")
		return &models.DeletedUnitAdmin{UnitAdminID: ""}, err
	}
	return &models.DeletedUnitAdmin{UnitAdminID: unitAdminID}, nil
}

// SetNewUnitAdminForRobboUnit is the resolver for the setNewUnitAdminForRobboUnit field.
func (r *mutationResolver) SetNewUnitAdminForRobboUnit(ctx context.Context, unitAdminID string, robboUnitID string) (string, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return "", err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return "", err
	}
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return "", err
	}
	setNewUnitAdminForRobboUnitErr := r.usersDelegate.SetNewUnitAdminForRobboUnit(unitAdminID, robboUnitID)
	if setNewUnitAdminForRobboUnitErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return "", nil
}

// DeleteUnitAdminForRobboUnit is the resolver for the DeleteUnitAdminForRobboUnit field.
func (r *mutationResolver) DeleteUnitAdminForRobboUnit(ctx context.Context, unitAdminID string, robboUnitID string) (string, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return "", err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return "", err
	}
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return "", err
	}
	deleteUnitAdminForRobboUnitErr := r.usersDelegate.DeleteUnitAdminForRobboUnit(unitAdminID, robboUnitID)
	if deleteUnitAdminForRobboUnitErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return "", nil
}

// UpdateSuperAdmin is the resolver for the updateSuperAdmin field.
func (r *mutationResolver) UpdateSuperAdmin(ctx context.Context, input models.UpdateSuperAdminInput) (models.SuperAdminResult, error) {
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
	updateSuperAdminInput := &models.SuperAdminHTTP{
		UserHTTP: &models.UserHTTP{
			ID:         input.SuperAdminHTTP.UserHTTP.ID,
			Email:      input.SuperAdminHTTP.UserHTTP.Email,
			Firstname:  input.SuperAdminHTTP.UserHTTP.Firstname,
			Lastname:   input.SuperAdminHTTP.UserHTTP.Lastname,
			Middlename: input.SuperAdminHTTP.UserHTTP.Middlename,
			Nickname:   input.SuperAdminHTTP.UserHTTP.Nickname,
		},
	}
	superAdminUpdated, updateSuperAdminErr := r.usersDelegate.UpdateSuperAdmin(updateSuperAdminInput)
	if updateSuperAdminErr != nil {
		err := updateSuperAdminErr
		return &models.Error{Message: err.Error()}, err
	}
	return superAdminUpdated, nil
}

// GetStudentsByParentID is the resolver for the GetStudentsByParentId field.
func (r *queryResolver) GetStudentsByParentID(ctx context.Context, parentID string) (models.StudentResult, error) {
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
func (r *queryResolver) GetStudentsByRobboGroup(ctx context.Context, robboGroupID string) (models.StudentResult, error) {
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
func (r *queryResolver) GetStudentsByRobboUnitID(ctx context.Context, robboUnitID string) (models.StudentResult, error) {
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

// GetStudentsByTeacherID is the resolver for the GetStudentsByTeacherId field.
func (r *queryResolver) GetStudentsByTeacherID(ctx context.Context, teacherID string) ([]*models.StudentHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	students, err := r.usersDelegate.GetStudentsByTeacherId(teacherID)
	return students, err
}

// SearchStudentsByEmail is the resolver for the SearchStudentsByEmail field.
func (r *queryResolver) SearchStudentsByEmail(ctx context.Context, email string, parentID string) (models.StudentResult, error) {
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
	students, searchStudentByEmailErr := r.usersDelegate.SearchStudentByEmail(email, parentID)
	if searchStudentByEmailErr != nil {
		err := searchStudentByEmailErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.StudentHTTPList{
		Students: students,
	}, nil
}

// GetAllTeachers is the resolver for the GetAllTeachers field.
func (r *queryResolver) GetAllTeachers(ctx context.Context) (models.TeacherResult, error) {
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
	teachers, getAllTeachersErr := r.usersDelegate.GetAllTeachers()
	if getAllTeachersErr != nil {
		err := getAllTeachersErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.TeacherHTTPList{
		Teachers: teachers,
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
func (r *queryResolver) GetTeachersByRobboGroupID(ctx context.Context, robboGroupID string) (models.TeacherResult, error) {
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

// GetTeachersByStudentID is the resolver for the GetTeachersByStudentId field.
func (r *queryResolver) GetTeachersByStudentID(ctx context.Context, studentID string) ([]*models.TeacherHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	teachers, err := r.usersDelegate.GetTeachersByStudentId(studentID)
	return teachers, err
}

// GetAllParents is the resolver for the GetAllParents field.
func (r *queryResolver) GetAllParents(ctx context.Context) (models.ParentResult, error) {
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
	parents, getAllParentsErr := r.usersDelegate.GetAllParent()
	if getAllParentsErr != nil {
		err := getAllParentsErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.ParentHTTPList{
		Parents: parents,
	}, nil
}

// GetParentByID is the resolver for the GetParentById field.
func (r *queryResolver) GetParentByID(ctx context.Context, parentID string) (models.ParentResult, error) {
	ginContext, ginContextErr := GinContextFromContext(ctx)
	if ginContextErr != nil {
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
	parent, getParentByIdErr := r.usersDelegate.GetParentById(parentID)
	if getParentByIdErr != nil {
		err := getParentByIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return parent, nil
}

// GetAllUnitAdmins is the resolver for the GetAllUnitAdmins field.
func (r *queryResolver) GetAllUnitAdmins(ctx context.Context) (models.UnitAdminResult, error) {
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
	unitAdmins, getAllUnitAdminsErr := r.usersDelegate.GetAllUnitAdmins()
	if getAllUnitAdminsErr != nil {
		err := getAllUnitAdminsErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.UnitAdminHTTPList{
		UnitAdmins: unitAdmins,
	}, nil
}

// GetUnitAdminsByRobboUnitID is the resolver for the GetUnitAdminsByRobboUnitId field.
func (r *queryResolver) GetUnitAdminsByRobboUnitID(ctx context.Context, robboUnitID string) (models.UnitAdminResult, error) {
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

	unitAdmins, getUnitAdminByRobboUnitIdErr := r.usersDelegate.GetUnitAdminByRobboUnitId(robboUnitID)
	if getUnitAdminByRobboUnitIdErr != nil {
		err := getUnitAdminByRobboUnitIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.UnitAdminHTTPList{
		UnitAdmins: unitAdmins,
	}, nil
}

// GetUnitAdminByID is the resolver for the GetUnitAdminById field.
func (r *queryResolver) GetUnitAdminByID(ctx context.Context, unitAdminID string) (models.UnitAdminResult, error) {
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
	unitAdmin, getUnitAdminByIdErr := r.usersDelegate.GetUnitAdminById(unitAdminID)
	if getUnitAdminByIdErr != nil {
		err := getUnitAdminByIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &unitAdmin, nil
}

// SearchUnitAdminsByEmail is the resolver for the SearchUnitAdminsByEmail field.
func (r *queryResolver) SearchUnitAdminsByEmail(ctx context.Context, email string, robboUnitID string) (models.UnitAdminResult, error) {
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
	unitAdmins, searchUnitAdminByEmailErr := r.usersDelegate.SearchUnitAdminByEmail(email, robboUnitID)
	if searchUnitAdminByEmailErr != nil {
		err := searchUnitAdminByEmailErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.UnitAdminHTTPList{
		UnitAdmins: unitAdmins,
	}, nil
}

// GetSuperAdminByID is the resolver for the GetSuperAdminById field.
func (r *queryResolver) GetSuperAdminByID(ctx context.Context, superAdminID string) (models.SuperAdminResult, error) {
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
	superAdmin, getSuperAdminByIdErr := r.usersDelegate.GetSuperAdminById(superAdminID)
	if getSuperAdminByIdErr != nil {
		err := getSuperAdminByIdErr
		return &models.Error{Message: err.Error()}, err
	}
	return &superAdmin, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
