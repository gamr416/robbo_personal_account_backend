package http

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth/delegate/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUp(t *testing.T) {
	r := gin.Default()
	delegate := new(mock.AuthDelegateMock)
	handler := new(Handler)
	handler.delegate = delegate
	handler.InitAuthRoutes(r)

	signUpBody := &signInput{
		Email:    "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	delegate.On("SignUp", signUpBody.Email, signUpBody.Password).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/sign-up", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestSignIn(t *testing.T) {
	r := gin.Default()
	delegate := new(mock.AuthDelegateMock)
	handler := new(Handler)
	handler.delegate = delegate
	handler.InitAuthRoutes(r)

	signUpBody := &signInput{
		Email:    "testemail",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	delegate.On("SignIn", signUpBody.Email, signUpBody.Password).Return("jwt", nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/sign-in", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"token\":\"jwt\"}", w.Body.String())
}
