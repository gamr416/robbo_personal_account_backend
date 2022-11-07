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

func NewRobboGroupHandler(
	authDelegate auth.Delegate,
	robboGroup robboGroup.Delegate,
) Handler {
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
		//robboGroup.POST("/robboGroupId", h.GetRobboGroupsByRobboUnitId)
		robboGroup.POST("/setTeacher", h.SetTeacherForRobboGroup)
		robboGroup.DELETE("/deleteTeacher", h.DeleteTeacherForRobboGroup)
	}
}

func (h *Handler) CreateRobboGroup(c *gin.Context) {
	fmt.Println("Create Robbo Unit")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	robboUnitId := c.Param("robboUnitId")
	robboGroupHttp := models.RobboGroupHTTP{}
	if err := c.BindJSON(&robboGroupHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(robboGroupHttp)

	robboGroupHttp.RobboUnitID = robboUnitId
	robboGroupId, err := h.robboGroupDelegate.CreateRobboGroup(&robboGroupHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"robboGroupId": robboGroupId,
	})
}

func (h *Handler) GetRobboGroupById(c *gin.Context) {
	fmt.Println("Get RobboUnit By Id")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
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
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
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
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	robboGroupId := c.Param("robboGroupId")
	err := h.robboGroupDelegate.DeleteRobboGroup(robboGroupId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

type SetTeacherForRobboGroupInput struct {
	TeacherId    string `json:"teacherId"`
	RobboGroupId string `json:"robboGroupId"`
}

func (h *Handler) SetTeacherForRobboGroup(c *gin.Context) {
	fmt.Println("SetTeacherForRobboGroup")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	setTeacherForRobboGroupInput := new(SetTeacherForRobboGroupInput)

	if err := c.BindJSON(setTeacherForRobboGroupInput); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	setTeacherForRobboGroupErr := h.robboGroupDelegate.SetTeacherForRobboGroup(setTeacherForRobboGroupInput.TeacherId, setTeacherForRobboGroupInput.RobboGroupId)

	if setTeacherForRobboGroupErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteTeacherForRobboGroup(c *gin.Context) {
	fmt.Println("DeleteTeacherForRobboGroup")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	deleteTeacherForRobboGroupInput := new(SetTeacherForRobboGroupInput)

	if err := c.BindJSON(deleteTeacherForRobboGroupInput); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	deleteTeacherForRobboGroupErr := h.robboGroupDelegate.DeleteTeacherForRobboGroup(deleteTeacherForRobboGroupInput.TeacherId, deleteTeacherForRobboGroupInput.RobboGroupId)

	if deleteTeacherForRobboGroupErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
