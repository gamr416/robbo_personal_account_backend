package auth

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrInvalidTypeClaims  = errors.New("token claims are not of type *StandardClaims")
	ErrTokenNotFound      = errors.New("token not found")
)
