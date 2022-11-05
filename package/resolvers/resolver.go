package resolvers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authDelegate       auth.Delegate
	usersDelegate      users.Delegate
	robboGroupDelegate robboGroup.Delegate
	robboUnitsDelegate robboUnits.Delegate
}

type MutationResolver struct{ *Resolver }
type QueryResolver struct{ *Resolver }

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func NewUsersResolver(
	authDelegate auth.Delegate,
	usersDelegate users.Delegate,
	robboGroupDelegate robboGroup.Delegate,
	robboUnitsDelegate robboUnits.Delegate,
) Resolver {
	return Resolver{
		authDelegate:       authDelegate,
		usersDelegate:      usersDelegate,
		robboGroupDelegate: robboGroupDelegate,
		robboUnitsDelegate: robboUnitsDelegate,
	}
}
