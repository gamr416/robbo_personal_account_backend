package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/cohorts"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/middleware"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	middleware.Middleware
	authDelegate    auth.Delegate
	cohortsDelegate cohorts.Delegate
}

func NewCohortsHandler(authDelegate auth.Delegate, cohortsDelegate cohorts.Delegate) Handler {
	return Handler{
		authDelegate:    authDelegate,
		cohortsDelegate: cohortsDelegate,
	}
}

type testCohortResponse struct {
	CohortID string `json:"cohortId"`
}

func (h *Handler) InitCohortRoutes(router *gin.Engine) {
	cohort := router.Group("/cohort")
	{
		cohort.POST("/createCohort/:courseId", h.CreateCohort)
		cohort.POST("/addStudent/:username/:courseId/:cohortId", h.AddStudent)
	}
}

func (h *Handler) CreateCohort(c *gin.Context) {
	fmt.Println("Create Cohort")
	_, _, userIdentityErr := h.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	createCohortResponse := models.CreateCohortHTTP{}
	courseId := c.Param("courseId")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &createCohortResponse)
	fmt.Println(createCohortResponse)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	cohortHTTP := models.CohortHTTP{}

	cohortId, err := h.cohortsDelegate.CreateCohort(&cohortHTTP, &createCohortResponse, courseId)

	fmt.Println(cohortHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, testCohortResponse{
		cohortId,
	})
}

func (h *Handler) AddStudent(c *gin.Context) {
	fmt.Println("Add Student")
	_, _, userIdentityErr := h.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	tempCohortId := c.Param("cohortId")
	cohortId, _ := strconv.Atoi(tempCohortId)
	courseId := c.Param("courseId")
	username := c.Param("username")
	err := h.cohortsDelegate.AddStudent(username, courseId, cohortId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}
