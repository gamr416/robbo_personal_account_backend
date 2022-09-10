package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/skinnykaen/robbo_student_personal_account.git/graph/generated"
	"github.com/skinnykaen/robbo_student_personal_account.git/graph/model"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateStudent is the resolver for the createStudent field.
func (r *mutationResolver) CreateStudent(ctx context.Context, input model.NewStudent) (*model.Student, error) {
	panic(fmt.Errorf("not implemented: CreateStudent - createStudent"))
}

// CreateParent is the resolver for the createParent field.
func (r *mutationResolver) CreateParent(ctx context.Context, input model.NewParent) (*model.Parent, error) {
	userHttp := models.UserHttp{
		Email:      input.Email,
		Password:   input.Password,
		Lastname:   input.Lastname,
		Firstname:  input.Firstname,
		Nickname:   input.Nickname,
		Middlename: input.Middlename,
		Role:       uint(models.Parent),
	}
	parentHttp := &models.ParentHTTP{
		UserHttp: userHttp,
	}

	parentId, err := r.usersDelegate.CreateParent(parentHttp)
	parent := &model.Parent{
		User: &model.User{
			ID:         parentId,
			Email:      input.Email,
			Password:   input.Password,
			Lastname:   input.Lastname,
			Firstname:  input.Firstname,
			Nickname:   input.Nickname,
			Middlename: input.Middlename,
		},
	}
	return parent, err
}

// Students is the resolver for the Students field.
func (r *queryResolver) Students(ctx context.Context, limit *int, offset *int) ([]*model.Student, error) {
	panic(fmt.Errorf("not implemented: Students - Students"))
}

// GetAllParents is the resolver for the GetAllParents field.
func (r *queryResolver) GetAllParents(ctx context.Context) ([]*model.Parent, error) {
	panic(fmt.Errorf("not implemented: GetAllParents - GetAllParents"))
}

// GetParentByID is the resolver for the GetParentById field.
func (r *queryResolver) GetParentByID(ctx context.Context, parentID string) (*model.Parent, error) {
	panic(fmt.Errorf("not implemented: GetParentByID - GetParentById"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
