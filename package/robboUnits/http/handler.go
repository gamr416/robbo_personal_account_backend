package http

import (
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	"log"
	"net/http"
)

type Handler struct {
	authDelegate       auth.Delegate
	robboUnitsDelegate robboUnits.Delegate
}

func NewRobboUnitsHandler(
	authDelegate auth.Delegate,
	robboUnits robboUnits.Delegate,
) Handler {
	return Handler{
		authDelegate:       authDelegate,
		robboUnitsDelegate: robboUnits,
	}
}

func (h *Handler) InitRobboUnitsRoutes(router *gin.Engine) {
	robboUnits := router.Group("/robboUnits")
	{
		robboUnits.POST("/", h.CreateRobboUnit)
		robboUnits.GET("/:robboUnitId", h.GetRobboUnitById)
		robboUnits.GET("/unitAdmin", h.GetRobboUnitsByUnitAdminId)
		robboUnits.GET("/", h.GetAllRobboUnits)
		robboUnits.PUT("/", h.UpdateRobboUnit)
		robboUnits.DELETE("/:robboUnitId", h.DeleteRobboUnit)
	}
}

func (h *Handler) CreateRobboUnit(c *gin.Context) {
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

	robboUnitHttp := models.RobboUnitHTTP{}
	if err := c.BindJSON(&robboUnitHttp); err != nil {
		err = robboUnits.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	robboUnitId, err := h.robboUnitsDelegate.CreateRobboUnit(&robboUnitHttp)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"robboUnitId": robboUnitId,
	})
}

func (h *Handler) GetRobboUnitById(c *gin.Context) {
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
	robboUnitId := c.Param("robboUnitId")
	robboUnit, err := h.robboUnitsDelegate.GetRobboUnitById(robboUnitId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, robboUnit)
}

func (h *Handler) GetAllRobboUnits(c *gin.Context) {
	log.Println("Get All RobboUnits")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	robboUnits, err := h.robboUnitsDelegate.GetAllRobboUnit()
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, robboUnits)
}

func (h *Handler) GetRobboUnitsByUnitAdminId(c *gin.Context) {
	log.Println("Get RobboUnits By UnitAdminId")

	id, role, identityErr := h.authDelegate.UserIdentity(c)
	if identityErr != nil {
		log.Println(identityErr)
		ErrorHandling(identityErr, c)
		return
	}
	//TODO: add access to superAdmin
	allowedRoles := []models.Role{models.UnitAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	robboUnits, err := h.robboUnitsDelegate.GetRobboUnitsByUnitAdminId(id)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, robboUnits)
}

func (h *Handler) UpdateRobboUnit(c *gin.Context) {
	log.Println("Update RobboUnit")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	robboUnitHttp := models.RobboUnitHTTP{}

	if err := c.BindJSON(&robboUnitHttp); err != nil {
		err = robboUnits.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	err := h.robboUnitsDelegate.UpdateRobboUnit(&robboUnitHttp)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteRobboUnit(c *gin.Context) {
	log.Println("Delete RobboUnit")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	robboUnitId := c.Param("robboUnitId")
	err := h.robboUnitsDelegate.DeleteRobboUnit(robboUnitId)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.Status(http.StatusOK)
}

func ErrorHandling(err error, c *gin.Context) {
	switch err {
	case robboUnits.ErrBadRequest:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case robboUnits.ErrInternalServerLevel:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	case robboUnits.ErrBadRequestBody:
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
