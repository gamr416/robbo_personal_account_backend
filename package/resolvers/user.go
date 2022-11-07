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
	err := r.usersDelegate.UpdateStudent(updateStudentInput)
	return updateStudentInput, err
}

// DeleteStudent is the resolver for the deleteStudent field.
func (r *mutationResolver) DeleteStudent(ctx context.Context, studentID string) (string, error) {
	id, _ := strconv.ParseUint(studentID, 10, 64)
	err := r.usersDelegate.DeleteStudent(uint(id))
	return studentID, err
}

// SetRobboGroupIDForStudent is the resolver for the setRobboGroupIdForStudent field.
func (r *mutationResolver) SetRobboGroupIDForStudent(ctx context.Context, studentID string, robboGroupID string, robboUnitID string) (string, error) {
	err := r.usersDelegate.AddStudentToRobboGroup(studentID, robboGroupID, robboUnitID)
	return "", err
}

// CreateTeacher is the resolver for the createTeacher field.
func (r *mutationResolver) CreateTeacher(ctx context.Context, input models.NewTeacher) (*models.TeacherHTTP, error) {
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

	teacherId, err := r.usersDelegate.CreateTeacher(&teacherInput)
	teacherInput.UserHTTP.ID = teacherId
	return &teacherInput, err
}

// UpdateTeacher is the resolver for the updateTeacher field.
func (r *mutationResolver) UpdateTeacher(ctx context.Context, input models.UpdateTeacherInput) (*models.TeacherHTTP, error) {
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
	err := r.usersDelegate.UpdateTeacher(updateTeacherInput)
	return updateTeacherInput, err
}

// DeleteTeacher is the resolver for the deleteTeacher field.
func (r *mutationResolver) DeleteTeacher(ctx context.Context, teacherID string) (string, error) {
	id, _ := strconv.ParseUint(teacherID, 10, 64)
	err := r.usersDelegate.DeleteTeacher(uint(id))
	return teacherID, err
}

// CreateParent is the resolver for the createParent field.
func (r *mutationResolver) CreateParent(ctx context.Context, input models.NewParent) (*models.ParentHTTP, error) {
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
	parentId, err := r.usersDelegate.CreateParent(&parentInput)
	parentInput.UserHTTP.ID = parentId
	return &parentInput, err
}

// AddChildToParent is the resolver for the addChildToParent field.
func (r *mutationResolver) AddChildToParent(ctx context.Context, parentID string, childID string) (string, error) {
	err := r.usersDelegate.CreateRelation(parentID, childID)
	return "", err
}

// UpdateParent is the resolver for the updateParent field.
func (r *mutationResolver) UpdateParent(ctx context.Context, input models.UpdateParentInput) (*models.ParentHTTP, error) {
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
	err := r.usersDelegate.UpdateParent(updateParentInput)
	return updateParentInput, err
}

// DeleteParent is the resolver for the deleteParent field.
func (r *mutationResolver) DeleteParent(ctx context.Context, parentID string) (string, error) {
	id, _ := strconv.ParseUint(parentID, 10, 64)
	err := r.usersDelegate.DeleteParent(uint(id))
	return parentID, err
}

// CreateUnitAdmin is the resolver for the createUnitAdmin field.
func (r *mutationResolver) CreateUnitAdmin(ctx context.Context, input models.NewUnitAdmin) (*models.UnitAdminHTTP, error) {
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
	unitAdminId, err := r.usersDelegate.CreateUnitAdmin(&unitAdminInput)
	unitAdminInput.UserHTTP.ID = unitAdminId
	return &unitAdminInput, err
}

// UpdateUnitAdmin is the resolver for the updateUnitAdmin field.
func (r *mutationResolver) UpdateUnitAdmin(ctx context.Context, input models.UpdateUnitAdminInput) (*models.UnitAdminHTTP, error) {
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
	err := r.usersDelegate.UpdateUnitAdmin(updateUnitAdminInput)
	return updateUnitAdminInput, err
}

// DeleteUnitAdmin is the resolver for the deleteUnitAdmin field.
func (r *mutationResolver) DeleteUnitAdmin(ctx context.Context, unitAdminID string) (string, error) {
	id, _ := strconv.ParseUint(unitAdminID, 10, 64)
	err := r.usersDelegate.DeleteUnitAdmin(uint(id))
	return unitAdminID, err
}

// SetNewUnitAdminForRobboUnit is the resolver for the setNewUnitAdminForRobboUnit field.
func (r *mutationResolver) SetNewUnitAdminForRobboUnit(ctx context.Context, unitAdminID string, robboUnitID string) (string, error) {
	err := r.usersDelegate.SetNewUnitAdminForRobboUnit(unitAdminID, robboUnitID)
	return "", err
}

// DeleteUnitAdminForRobboUnit is the resolver for the DeleteUnitAdminForRobboUnit field.
func (r *mutationResolver) DeleteUnitAdminForRobboUnit(ctx context.Context, unitAdminID string, robboUnitID string) (string, error) {
	err := r.usersDelegate.DeleteUnitAdminForRobboUnit(unitAdminID, robboUnitID)
	return "", err
}

// UpdateSuperAdmin is the resolver for the updateSuperAdmin field.
func (r *mutationResolver) UpdateSuperAdmin(ctx context.Context, input models.UpdateSuperAdminInput) (*models.SuperAdminHTTP, error) {
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
	err := r.usersDelegate.UpdateSuperAdmin(updateSuperAdminInput)
	return updateSuperAdminInput, err
}

// GetStudentsByParentID is the resolver for the GetStudentsByParentId field.
func (r *queryResolver) GetStudentsByParentID(ctx context.Context, parentID string) ([]*models.StudentHTTP, error) {
	return r.usersDelegate.GetStudentByParentId(parentID)
}

// GetStudentByID is the resolver for the GetStudentById field.
func (r *queryResolver) GetStudentByID(ctx context.Context, studentID string) (*models.StudentHTTP, error) {
	student, err := r.usersDelegate.GetStudentById(studentID)
	return student, err
}

// SearchStudentsByEmail is the resolver for the SearchStudentsByEmail field.
func (r *queryResolver) SearchStudentsByEmail(ctx context.Context, email string) ([]*models.StudentHTTP, error) {
	return r.usersDelegate.SearchStudentByEmail(email)
}

// GetAllTeachers is the resolver for the GetAllTeachers field.
func (r *queryResolver) GetAllTeachers(ctx context.Context) ([]*models.TeacherHTTP, error) {
	return r.usersDelegate.GetAllTeachers()
}

// GetTeacherByID is the resolver for the GetTeacherById field.
func (r *queryResolver) GetTeacherByID(ctx context.Context, teacherID string) (*models.TeacherHTTP, error) {
	teacher, err := r.usersDelegate.GetTeacherById(teacherID)
	return teacher, err
}

// GetAllParents is the resolver for the GetAllParents field.
func (r *queryResolver) GetAllParents(ctx context.Context) ([]*models.ParentHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}

	_, _, userIdentityErr := r.authDelegate.UserIdentity(ginContext)
	if userIdentityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}

	parents, err := r.usersDelegate.GetAllParent()
	return parents, err
}

// GetParentByID is the resolver for the GetParentById field.
func (r *queryResolver) GetParentByID(ctx context.Context, parentID string) (*models.ParentHTTP, error) {
	parent, err := r.usersDelegate.GetParentById(parentID)
	return parent, err
}

// GetAllUnitAdmins is the resolver for the GetAllUnitAdmins field.
func (r *queryResolver) GetAllUnitAdmins(ctx context.Context) ([]*models.UnitAdminHTTP, error) {
	return r.usersDelegate.GetAllUnitAdmins()
}

// GetUnitAdminsByRobboUnitID is the resolver for the GetUnitAdminsByRobboUnitId field.
func (r *queryResolver) GetUnitAdminsByRobboUnitID(ctx context.Context, robboUnitID string) ([]*models.UnitAdminHTTP, error) {
	return r.usersDelegate.GetUnitAdminByRobboUnitId(robboUnitID)
}

// GetUnitAdminByID is the resolver for the GetUnitAdminById field.
func (r *queryResolver) GetUnitAdminByID(ctx context.Context, unitAdminID string) (*models.UnitAdminHTTP, error) {
	unitAdmin, err := r.usersDelegate.GetUnitAdminById(unitAdminID)
	return &unitAdmin, err
}

// SearchUnitAdminsByEmail is the resolver for the SearchUnitAdminsByEmail field.
func (r *queryResolver) SearchUnitAdminsByEmail(ctx context.Context, email string) ([]*models.UnitAdminHTTP, error) {
	return r.usersDelegate.SearchUnitAdminByEmail(email)
}

// GetSuperAdminByID is the resolver for the GetSuperAdminById field.
func (r *queryResolver) GetSuperAdminByID(ctx context.Context, superAdminID string) (*models.SuperAdminHTTP, error) {
	superAdmin, err := r.usersDelegate.GetSuperAdminById(superAdminID)
	return &superAdmin, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{
		r,
	}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
