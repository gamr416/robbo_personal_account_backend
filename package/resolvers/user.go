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
func (r *mutationResolver) UpdateStudent(ctx context.Context, input models.UpdateProfileInput) (models.StudentResult, error) {
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
			ID:         input.ID,
			Email:      input.Email,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
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
func (r *mutationResolver) SetRobboGroupIDForStudent(ctx context.Context, studentID string, robboGroupID string, robboUnitID string) (*models.Error, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.Error{Message: "no access"}, err
	}

	addStudentToRobboGroupErr := r.usersDelegate.AddStudentToRobboGroup(studentID, robboGroupID, robboUnitID)
	if addStudentToRobboGroupErr != nil {
		return &models.Error{Message: addStudentToRobboGroupErr.Error()}, addStudentToRobboGroupErr
	}
	return &models.Error{}, nil
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
func (r *mutationResolver) UpdateTeacher(ctx context.Context, input models.UpdateProfileInput) (models.TeacherResult, error) {
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
			ID:         input.ID,
			Email:      input.Email,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
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
func (r *mutationResolver) AddChildToParent(ctx context.Context, parentID string, childID string) (models.StudentsResult, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.Error{Message: "no access"}, err
	}
	students, createRelationErr := r.usersDelegate.CreateStudentParentRelation(parentID, childID)
	if createRelationErr != nil {
		err := errors.New("baq request")
		return &models.Error{Message: "baq request"}, err
	}
	return &models.StudentHTTPList{Students: students}, nil
}

// UpdateParent is the resolver for the updateParent field.
func (r *mutationResolver) UpdateParent(ctx context.Context, input models.UpdateProfileInput) (models.ParentResult, error) {
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
			ID:         input.ID,
			Email:      input.Email,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
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
func (r *mutationResolver) UpdateUnitAdmin(ctx context.Context, input models.UpdateProfileInput) (models.UnitAdminResult, error) {
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
			ID:         input.ID,
			Email:      input.Email,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
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
func (r *mutationResolver) SetNewUnitAdminForRobboUnit(ctx context.Context, unitAdminID string, robboUnitID string) (*models.Error, error) {
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
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.Error{Message: "no access"}, err
	}
	setNewUnitAdminForRobboUnitErr := r.usersDelegate.SetNewUnitAdminForRobboUnit(unitAdminID, robboUnitID)
	if setNewUnitAdminForRobboUnitErr != nil {
		err := errors.New("baq request")
		return &models.Error{Message: "baq request"}, err
	}
	return &models.Error{Message: ""}, nil
}

// DeleteUnitAdminForRobboUnit is the resolver for the DeleteUnitAdminForRobboUnit field.
func (r *mutationResolver) DeleteUnitAdminForRobboUnit(ctx context.Context, unitAdminID string, robboUnitID string) (*models.Error, error) {
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
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return &models.Error{Message: "no access"}, err
	}
	deleteUnitAdminForRobboUnitErr := r.usersDelegate.DeleteUnitAdminForRobboUnit(unitAdminID, robboUnitID)
	if deleteUnitAdminForRobboUnitErr != nil {
		return &models.Error{Message: deleteUnitAdminForRobboUnitErr.Error()}, deleteUnitAdminForRobboUnitErr
	}
	return &models.Error{}, nil
}

// UpdateSuperAdmin is the resolver for the updateSuperAdmin field.
func (r *mutationResolver) UpdateSuperAdmin(ctx context.Context, input models.UpdateProfileInput) (models.SuperAdminResult, error) {
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
			ID:         input.ID,
			Email:      input.Email,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
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
func (r *queryResolver) SearchStudentsByEmail(ctx context.Context, email string, parentID string) (models.StudentsResult, error) {
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
func (r *queryResolver) SearchTeachersByEmail(ctx context.Context, email string) (models.TeachersResult, error) {
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
	teachers, searchTeachersByEmailErr := r.usersDelegate.SearchTeacherByEmail(email)
	if searchTeachersByEmailErr != nil {
		err := searchTeachersByEmailErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.TeacherHTTPList{
		Teachers: teachers,
	}, nil
}

// GetAllParents is the resolver for the GetAllParents field.
func (r *queryResolver) GetAllParents(ctx context.Context, page string, pageSize string) (models.ParentsResult, error) {
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
	parents, countRows, getAllParentsErr := r.usersDelegate.GetAllParent(page, pageSize)
	if getAllParentsErr != nil {
		err := getAllParentsErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.ParentHTTPList{
		Parents:   parents,
		CountRows: countRows,
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
func (r *queryResolver) GetAllUnitAdmins(ctx context.Context, page string, pageSize string) (models.UnitAdminsResult, error) {
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
	unitAdmins, countRows, getAllUnitAdminsErr := r.usersDelegate.GetAllUnitAdmins(page, pageSize)
	if getAllUnitAdminsErr != nil {
		err := getAllUnitAdminsErr
		return &models.Error{Message: err.Error()}, err
	}
	return &models.UnitAdminHTTPList{
		UnitAdmins: unitAdmins,
		CountRows:  countRows,
	}, nil
}

// GetUnitAdminsByRobboUnitID is the resolver for the GetUnitAdminsByRobboUnitId field.
func (r *queryResolver) GetUnitAdminsByRobboUnitID(ctx context.Context, robboUnitID string) (models.UnitAdminsResult, error) {
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
func (r *queryResolver) SearchUnitAdminsByEmail(ctx context.Context, email string, robboUnitID string) (models.UnitAdminsResult, error) {
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
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
