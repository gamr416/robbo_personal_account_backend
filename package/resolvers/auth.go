package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// SingIn is the resolver for the SingIn field.
func (r *mutationResolver) SingIn(ctx context.Context, input models.SignInInput) (models.SignInResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	accessToken, refreshToken, err := r.authDelegate.SignIn(input.Email, input.Password, uint(input.UserRole))
	if err != nil {
		return &models.Error{Message: err.Error()}, err
	}
	setRefreshToken(refreshToken, ginContext)
	return &models.SingInResponse{
		AccessToken: accessToken,
	}, nil
}

// SingOut is the resolver for the SingOut field.
func (r *mutationResolver) SingOut(ctx context.Context) (*models.Error, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	setRefreshToken("", ginContext)
	return &models.Error{}, nil
}

// Refresh is the resolver for the Refresh field.
func (r *mutationResolver) Refresh(ctx context.Context) (models.SignInResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.Error{Message: "internal server error"}, err
	}
	refreshToken, err := getRefreshToken(ginContext)
	if err != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: err.Error(),
			Extensions: map[string]interface{}{
				"code": "401",
			},
		}
	}

	newAccessToken, err := r.authDelegate.RefreshToken(refreshToken)
	if err != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: err.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.SingInResponse{
		AccessToken: newAccessToken,
	}, nil
}

func getRefreshToken(c *gin.Context) (refreshToken string, err error) {
	refreshToken = c.Value("refresh_token").(string)
	if refreshToken == "" {
		return "", errors.New("error finding cookie")
	}
	return
}
func setRefreshToken(value string, c *gin.Context) {
	c.SetCookie(
		"refresh_token",
		value,
		60*60*24*7,
		"/",
		"0.0.0.0",
		false,
		false,
	)
}
