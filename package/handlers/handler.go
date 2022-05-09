package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/delegate"
)

type RequestHandler struct {
	//fx.In
	delegate.AuthDelegate
}

func NewHandler(authDelegate delegate.AuthDelegate) RequestHandler {
	return RequestHandler{
		AuthDelegate: authDelegate,
	}
}

func (h *RequestHandler) InitRoutes() *gin.Engine {
	router := gin.New()
	test := router.Group("/test")
	{
		test.GET("/test", h.test)
	}
	return router
}
