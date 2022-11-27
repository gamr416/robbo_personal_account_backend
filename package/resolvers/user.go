package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"strconv"

	"github.com/skinnykaen/robbo_student_personal_account.git/graph/generated"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateStudent is the resolver for the createStudent field.
func (r *mutationResolver) CreateStudent(ctx context.Context, input models.NewStudent) (*models.StudentHTTP, error) {
	ginContext, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if role < models.UnitAdmin || identityErr != nil {
		return nil, identityErr
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

	studentId, err := r.usersDelegate.CreateStudent(&studentInput, input.ParentID)
	studentInput.UserHTTP.ID = studentId
	return &studentInput, err
}

// UpdateStudent is the resolver for the updateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, input models.UpdateStudentInput) (*models.StudentHTTP, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
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
	updateStudentErr := r.usersDelegate.UpdateStudent(updateStudentInput)
	if updateStudentErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return updateStudentInput, nil
}

// DeleteStudent is the resolver for the deleteStudent field.
func (r *mutationResolver) DeleteStudent(ctx context.Context, studentID string) (string, error) {
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
	id, _ := strconv.ParseUint(studentID, 10, 64)
	deleteStudentErr := r.usersDelegate.DeleteStudent(uint(id))
	if deleteStudentErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return studentID, nil
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

// CreateTeacher is the resolver for the createTeacher field.
func (r *mutationResolver) CreateTeacher(ctx context.Context, input models.NewTeacher) (*models.TeacherHTTP, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
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

	teacherId, createTeacherErr := r.usersDelegate.CreateTeacher(&teacherInput)
	if createTeacherErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	teacherInput.UserHTTP.ID = teacherId
	return &teacherInput, nil
}

// UpdateTeacher is the resolver for the updateTeacher field.
func (r *mutationResolver) UpdateTeacher(ctx context.Context, input models.UpdateTeacherInput) (*models.TeacherHTTP, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
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
	updateTeacherErr := r.usersDelegate.UpdateTeacher(updateTeacherInput)
	if updateTeacherErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return updateTeacherInput, nil
}

// DeleteTeacher is the resolver for the deleteTeacher field.
func (r *mutationResolver) DeleteTeacher(ctx context.Context, teacherID string) (string, error) {
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
	id, _ := strconv.ParseUint(teacherID, 10, 64)
	deleteTeacherErr := r.usersDelegate.DeleteTeacher(uint(id))
	if deleteTeacherErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return teacherID, nil
}

// CreateParent is the resolver for the createParent field.
func (r *mutationResolver) CreateParent(ctx context.Context, input models.NewParent) (*models.ParentHTTP, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
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
	parentId, createParentErr := r.usersDelegate.CreateParent(&parentInput)
	if createParentErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	parentInput.UserHTTP.ID = parentId
	return &parentInput, nil
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
func (r *mutationResolver) UpdateParent(ctx context.Context, input models.UpdateParentInput) (*models.ParentHTTP, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
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
	updateParentErr := r.usersDelegate.UpdateParent(updateParentInput)
	if updateParentErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return updateParentInput, nil
}

// DeleteParent is the resolver for the deleteParent field.
func (r *mutationResolver) DeleteParent(ctx context.Context, parentID string) (string, error) {
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
	id, _ := strconv.ParseUint(parentID, 10, 64)
	deleteParentErr := r.usersDelegate.DeleteParent(uint(id))
	if deleteParentErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return parentID, nil
}

// CreateUnitAdmin is the resolver for the createUnitAdmin field.
func (r *mutationResolver) CreateUnitAdmin(ctx context.Context, input models.NewUnitAdmin) (*models.UnitAdminHTTP, error) {
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
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
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
	unitAdminId, createUnitAdminErr := r.usersDelegate.CreateUnitAdmin(&unitAdminInput)
	if createUnitAdminErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	unitAdminInput.UserHTTP.ID = unitAdminId
	return &unitAdminInput, nil
}

// UpdateUnitAdmin is the resolver for the updateUnitAdmin field.
func (r *mutationResolver) UpdateUnitAdmin(ctx context.Context, input models.UpdateUnitAdminInput) (*models.UnitAdminHTTP, error) {
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
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
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
	updateUnitAdminErr := r.usersDelegate.UpdateUnitAdmin(updateUnitAdminInput)
	if updateUnitAdminErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return updateUnitAdminInput, nil
}

// DeleteUnitAdmin is the resolver for the deleteUnitAdmin field.
func (r *mutationResolver) DeleteUnitAdmin(ctx context.Context, unitAdminID string) (string, error) {
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
	id, _ := strconv.ParseUint(unitAdminID, 10, 64)
	deleteUnitAdminErr := r.usersDelegate.DeleteUnitAdmin(uint(id))
	if deleteUnitAdminErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return unitAdminID, nil
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
func (r *mutationResolver) UpdateSuperAdmin(ctx context.Context, input models.UpdateSuperAdminInput) (*models.SuperAdminHTTP, error) {
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
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
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
	updateSuperAdminErr := r.usersDelegate.UpdateSuperAdmin(updateSuperAdminInput)
	if updateSuperAdminErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return updateSuperAdminInput, nil
}

// GetStudentsByParentID is the resolver for the GetStudentsByParentId field.
func (r *queryResolver) GetStudentsByParentID(ctx context.Context, parentID string) ([]*models.StudentHTTP, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin, models.Parent}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	students, err := r.usersDelegate.GetStudentByParentId(parentID)
	return students, err
}

// GetStudentByID is the resolver for the GetStudentById field.
func (r *queryResolver) GetStudentByID(ctx context.Context, studentID string) (*models.StudentHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, _, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	student, err := r.usersDelegate.GetStudentById(studentID)
	return student, err
}

// GetStudentsByRobboGroup is the resolver for the GetStudentsByRobboGroup field.
func (r *queryResolver) GetStudentsByRobboGroup(ctx context.Context, robboGroupID string) ([]*models.StudentHTTP, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin, models.Teacher}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	students, err := r.usersDelegate.GetStudentsByRobboGroupId(robboGroupID)
	return students, err
}

// SearchStudentsByEmail is the resolver for the SearchStudentsByEmail field.
func (r *queryResolver) SearchStudentsByEmail(ctx context.Context, email string, parentID string) ([]*models.StudentHTTP, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	students, err := r.usersDelegate.SearchStudentByEmail(email, parentID)
	return students, err
}

// GetAllTeachers is the resolver for the GetAllTeachers field.
func (r *queryResolver) GetAllTeachers(ctx context.Context) ([]*models.TeacherHTTP, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	teachers, err := r.usersDelegate.GetAllTeachers()
	return teachers, err
}

// GetTeacherByID is the resolver for the GetTeacherById field.
func (r *queryResolver) GetTeacherByID(ctx context.Context, teacherID string) (*models.TeacherHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, _, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	teacher, err := r.usersDelegate.GetTeacherById(teacherID)
	return teacher, err
}

// GetTeachersByRobboGroupID is the resolver for the GetTeachersByRobboGroupId field.
func (r *queryResolver) GetTeachersByRobboGroupID(ctx context.Context, robboGroupID string) ([]*models.TeacherHTTP, error) {
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
	allowedRoles := []models.Role{models.SuperAdmin, models.UnitAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	teachers, err := r.usersDelegate.GetTeacherByRobboGroupId(robboGroupID)
	return teachers, err
}

// GetAllParents is the resolver for the GetAllParents field.
func (r *queryResolver) GetAllParents(ctx context.Context) ([]*models.ParentHTTP, error) {
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
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	parents, err := r.usersDelegate.GetAllParent()
	return parents, err
}

// GetParentByID is the resolver for the GetParentById field.
func (r *queryResolver) GetParentByID(ctx context.Context, parentID string) (*models.ParentHTTP, error) {
	ginContext, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	_, _, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		return nil, identityErr
	}

	parent, err := r.usersDelegate.GetParentById(parentID)
	return parent, err
}

// GetAllUnitAdmins is the resolver for the GetAllUnitAdmins field.
func (r *queryResolver) GetAllUnitAdmins(ctx context.Context) ([]*models.UnitAdminHTTP, error) {
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
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	return r.usersDelegate.GetAllUnitAdmins()
}

// GetUnitAdminsByRobboUnitID is the resolver for the GetUnitAdminsByRobboUnitId field.
func (r *queryResolver) GetUnitAdminsByRobboUnitID(ctx context.Context, robboUnitID string) ([]*models.UnitAdminHTTP, error) {
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
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	return r.usersDelegate.GetUnitAdminByRobboUnitId(robboUnitID)
}

// GetUnitAdminByID is the resolver for the GetUnitAdminById field.
func (r *queryResolver) GetUnitAdminByID(ctx context.Context, unitAdminID string) (*models.UnitAdminHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, _, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	unitAdmin, err := r.usersDelegate.GetUnitAdminById(unitAdminID)
	return &unitAdmin, err
}

// SearchUnitAdminsByEmail is the resolver for the SearchUnitAdminsByEmail field.
func (r *queryResolver) SearchUnitAdminsByEmail(ctx context.Context, email string, robboUnitID string) ([]*models.UnitAdminHTTP, error) {
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
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	return r.usersDelegate.SearchUnitAdminByEmail(email, robboUnitID)
}

// GetSuperAdminByID is the resolver for the GetSuperAdminById field.
func (r *queryResolver) GetSuperAdminByID(ctx context.Context, superAdminID string) (*models.SuperAdminHTTP, error) {
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
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	superAdmin, err := r.usersDelegate.GetSuperAdminById(superAdminID)
	return &superAdmin, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
