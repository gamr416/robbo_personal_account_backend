package resolvers

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authDelegate  auth.Delegate
	usersDelegate users.Delegate
}

type MutationResolver struct{ *Resolver }
type QueryResolver struct{ *Resolver }

func NewUsersResolver(
	authDelegate auth.Delegate,
	usersDelegate users.Delegate,
) Resolver {
	return Resolver{
		authDelegate:  authDelegate,
		usersDelegate: usersDelegate,
	}
}
