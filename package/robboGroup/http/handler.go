package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup"
	"log"
	"net/http"
)

type Handler struct {
	authDelegate       auth.Delegate
	robboGroupDelegate robboGroup.Delegate
}

func NewRobboGroupHandler(authDelegate auth.Delegate, robboGroup robboGroup.Delegate) Handler {
	return Handler{
		authDelegate:       authDelegate,
		robboGroupDelegate: robboGroup,
	}
}

func (h *Handler) InitRobboGroupRoutes(router *gin.Engine) {
	robboGroup := router.Group("/robboUnits/:robboUnitId/robboGroup")
	{
		robboGroup.POST("/", h.CreateRobboGroup)
		robboGroup.GET("/:robboGroupId", h.GetRobboGroupById)
		robboGroup.GET("/", h.GetRobboGroupsByRobboUnitId)
		robboGroup.DELETE("/:robboGroupId", h.DeleteRobboUnit)
	}
}

func (h *Handler) CreateRobboGroup(c *gin.Context) {
	fmt.Println("Create Robbo Unit")

	robboGroupHttp := models.RobboGroupHttp{}
	if err := c.BindJSON(&robboGroupHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	robboUnitId, err := h.robboGroupDelegate.CreateRobboGroup(&robboGroupHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"robboGroupId": robboUnitId,
	})
}

func (h *Handler) GetRobboGroupById(c *gin.Context) {
	fmt.Println("Get RobboUnit By Id")
	robboGroupId := c.Param("robboGroupId")

	robboGroup, err := h.robboGroupDelegate.GetRobboGroupById(robboGroupId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, robboGroup)
}

func (h *Handler) GetRobboGroupsByRobboUnitId(c *gin.Context) {
	fmt.Println("Get all robboUnits")

	robboUnitId := c.Param("robboUnitId")

	robboGroups, err := h.robboGroupDelegate.GetRobboGroupsByRobboUnitId(robboUnitId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, robboGroups)
}

func (h *Handler) DeleteRobboUnit(c *gin.Context) {
	fmt.Println("Delete RobboUnit")

	robboGroupId := c.Param("robboGroupId")
	err := h.robboGroupDelegate.DeleteRobboGroup(robboGroupId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
