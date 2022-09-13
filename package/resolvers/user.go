package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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

// GetAllParents is the resolver for the GetAllParents field.
func (r *queryResolver) GetAllParents(ctx context.Context) ([]*models.ParentHTTP, error) {
	parents, err := r.usersDelegate.GetAllParent()
	return parents, err
}

// GetParentByID is the resolver for the GetParentById field.
func (r *queryResolver) GetParentByID(ctx context.Context, parentID string) (*models.ParentHTTP, error) {
	parent, err := r.usersDelegate.GetParentById(parentID)
	return parent, err
}

func (r *queryResolver) GetAllTeachers(ctx context.Context) ([]*models.TeacherHTTP, error) {
	return r.usersDelegate.GetAllTeachers()
}

func (r *queryResolver) GetTeacherByID(ctx context.Context, teacherID string) (*models.TeacherHTTP, error) {
	teacher, err := r.usersDelegate.GetTeacherById(teacherID)
	return teacher, err
}

func (r *queryResolver) GetAllUnitAdmins(ctx context.Context) ([]*models.UnitAdminHTTP, error) {
	return r.usersDelegate.GetAllUnitAdmins()
}

func (r *queryResolver) GetUnitAdminsByRobboUnitID(ctx context.Context, robboUnitID string) ([]*models.UnitAdminHTTP, error) {
	return r.usersDelegate.GetUnitAdminByRobboUnitId(robboUnitID)
}

func (r *queryResolver) GetUnitAdminByID(ctx context.Context, unitAdminID string) (*models.UnitAdminHTTP, error) {
	unitAdmin, err := r.usersDelegate.GetUnitAdminById(unitAdminID)
	return &unitAdmin, err
}

func (r *queryResolver) SearchUnitAdminsByEmail(ctx context.Context, email string) ([]*models.UnitAdminHTTP, error) {
	return r.usersDelegate.SearchUnitAdminByEmail(email)
}

func (r *queryResolver) GetSuperAdminByID(ctx context.Context, superAdminID string) (*models.SuperAdminHTTP, error) {
	superAdmin, err := r.usersDelegate.GetSuperAdminById(superAdminID)
	return &superAdmin, err
}

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
