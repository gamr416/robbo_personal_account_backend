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
		robboGroup.DELETE("/:robboGroupId", h.DeleteRobboGroup)
		//robboGroup.POST("/robboGroupId", h.GetRobboGroupsByRobboUnitId)
		robboGroup.POST("/setTeacher", h.SetTeacherForRobboGroup)
		robboGroup.DELETE("/deleteTeacher", h.DeleteTeacherForRobboGroup)
	}
}

func (h *Handler) CreateRobboGroup(c *gin.Context) {
	log.Println("Create Robbo Unit")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	robboUnitId := c.Param("robboUnitId")
	robboGroupHttp := models.RobboGroupHTTP{}
	if err := c.BindJSON(&robboGroupHttp); err != nil {
		err = robboGroup.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	fmt.Println(robboGroupHttp)

	robboGroupHttp.RobboUnitID = robboUnitId
	robboGroupId, err := h.robboGroupDelegate.CreateRobboGroup(&robboGroupHttp)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"robboGroupId": robboGroupId,
	})
}

func (h *Handler) GetRobboGroupById(c *gin.Context) {
	log.Println("Get RobboUnit By Id")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	robboGroupId := c.Param("robboGroupId")

	robboGroup, err := h.robboGroupDelegate.GetRobboGroupById(robboGroupId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, robboGroup)
}

func (h *Handler) GetRobboGroupsByRobboUnitId(c *gin.Context) {
	log.Println("Get all robboUnits")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	robboUnitId := c.Param("robboUnitId")

	robboGroups, err := h.robboGroupDelegate.GetRobboGroupsByRobboUnitId(robboUnitId)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, robboGroups)
}

func (h *Handler) DeleteRobboGroup(c *gin.Context) {
	log.Println("Delete RobboGroup")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	robboGroupId := c.Param("robboGroupId")
	err := h.robboGroupDelegate.DeleteRobboGroup(robboGroupId)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.Status(http.StatusOK)
}

type SetTeacherForRobboGroupInput struct {
	TeacherId    string `json:"teacherId"`
	RobboGroupId string `json:"robboGroupId"`
}

func (h *Handler) SetTeacherForRobboGroup(c *gin.Context) {
	log.Println("Set Teacher For RobboGroup")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	setTeacherForRobboGroupInput := new(SetTeacherForRobboGroupInput)

	if err := c.BindJSON(setTeacherForRobboGroupInput); err != nil {
		err = robboGroup.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	setTeacherForRobboGroupErr := h.robboGroupDelegate.SetTeacherForRobboGroup(setTeacherForRobboGroupInput.TeacherId, setTeacherForRobboGroupInput.RobboGroupId)

	if setTeacherForRobboGroupErr != nil {
		log.Println(setTeacherForRobboGroupErr)
		ErrorHandling(setTeacherForRobboGroupErr, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteTeacherForRobboGroup(c *gin.Context) {
	log.Println("Delete Teacher For RobboGroup")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	deleteTeacherForRobboGroupInput := new(SetTeacherForRobboGroupInput)

	if err := c.BindJSON(deleteTeacherForRobboGroupInput); err != nil {
		err = robboGroup.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	deleteTeacherForRobboGroupErr := h.robboGroupDelegate.DeleteTeacherForRobboGroup(deleteTeacherForRobboGroupInput.TeacherId, deleteTeacherForRobboGroupInput.RobboGroupId)


	if deleteTeacherForRobboGroupErr != nil {
		log.Println(deleteTeacherForRobboGroupErr)
		ErrorHandling(deleteTeacherForRobboGroupErr, c)
		return
	}

	c.Status(http.StatusOK)
}

func ErrorHandling(err error, c *gin.Context) {
	switch err {
	case robboGroup.ErrBadRequest:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case robboGroup.ErrInternalServerLevel:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	case robboGroup.ErrBadRequestBody:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case auth.ErrInvalidAccessToken:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	case auth.ErrTokenNotFound:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	case auth.ErrNotAccess:
		c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
