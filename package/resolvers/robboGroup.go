package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// GetRobboGroupByID is the resolver for the GetRobboGroupById field.
func (r *queryResolver) GetRobboGroupByID(ctx context.Context, id string) (*models.RobboGroupHTTP, error) {
	robboGroupsHttp, err := r.robboGroupDelegate.GetRobboGroupById(id)
	return &robboGroupsHttp, err
}

// GetRobboGroupsByTeacherID is the resolver for the GetRobboGroupsByTeacherId field.
func (r *queryResolver) GetRobboGroupsByTeacherID(ctx context.Context, teacherID string) ([]*models.RobboGroupHTTP, error) {
	robboGroupsHttp, err := r.robboGroupDelegate.GetRobboGroupsByTeacherId(teacherID)
	return robboGroupsHttp, err
}

// SearchGroupsByName is the resolver for the SearchGroupsByName field.
func (r *queryResolver) SearchGroupsByName(ctx context.Context, name string) ([]*models.RobboGroupHTTP, error) {
	robboGroupsHttp, err := r.robboGroupDelegate.SearchRobboGroupByName(name)
	return robboGroupsHttp, err
}
