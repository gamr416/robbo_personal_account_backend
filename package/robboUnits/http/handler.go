package http

import (
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
)

type Handler struct {
	authDelegate       auth.Delegate
	robboUnitsDelegate robboUnits.Delegate
}

func NewRobboUnitsHandler(authDelegate auth.Delegate, robboUnits robboUnits.Delegate) Handler {
	return Handler{
		authDelegate:       authDelegate,
		robboUnitsDelegate: robboUnits,
	}
}

func (h *Handler) InitProjectRoutes(router *gin.Engine) {
	robboUnits := router.Group("/robboUnits")
	{
		robboUnits.POST("/", h.CreateRobboUnit)
		robboUnits.GET("/:robboUnitId", h.GetRobboUnitById)
		robboUnits.GET("/", h.GetAllRobboUnits)
		robboUnits.PUT("/", h.UpdateRobboUnit)
		robboUnits.DELETE("/:robboUnitId", h.DeleteRobboUnit)
	}
}

func (h *Handler) CreateRobboUnit(c *gin.Context) {

}

func (h *Handler) GetRobboUnitById(c *gin.Context) {

}

func (h *Handler) GetAllRobboUnits(c *gin.Context) {

}

func (h *Handler) UpdateRobboUnit(c *gin.Context) {

}

func (h *Handler) DeleteRobboUnit(c *gin.Context) {

}