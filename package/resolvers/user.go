package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/skinnykaen/robbo_student_personal_account.git/graph/generated"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateStudent is the resolver for the createStudent field.
func (r *mutationResolver) CreateStudent(ctx context.Context, input models.NewStudent) (*models.StudentHTTP, error) {
	panic(fmt.Errorf("not implemented: CreateStudent - createStudent"))
}

// CreateParent is the resolver for the createParent field.
func (r *mutationResolver) CreateParent(ctx context.Context, input models.NewParent) (*models.ParentHTTP, error) {
	panic(fmt.Errorf("not implemented: CreateParent - createParent"))
}

// Students is the resolver for the Students field.
func (r *queryResolver) Students(ctx context.Context, limit *int, offset *int) ([]*models.StudentHTTP, error) {
	panic(fmt.Errorf("not implemented: Students - Students"))
}

// GetAllParents is the resolver for the GetAllParents field.
func (r *queryResolver) GetAllParents(ctx context.Context) ([]*models.ParentHTTP, error) {
	panic(fmt.Errorf("not implemented: GetAllParents - GetAllParents"))
}

// GetParentByID is the resolver for the GetParentById field.
func (r *queryResolver) GetParentByID(ctx context.Context, parentID string) (*models.ParentHTTP, error) {
	panic(fmt.Errorf("not implemented: GetParentByID - GetParentById"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
