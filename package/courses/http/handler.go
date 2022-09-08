package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	authDelegate    auth.Delegate
	coursesDelegate courses.Delegate
}

func NewCoursesHandler(
	authDelegate auth.Delegate,
	coursesDelegate courses.Delegate,
) Handler {
	return Handler{
		authDelegate:    authDelegate,
		coursesDelegate: coursesDelegate,
	}
}

type testCourseResponse struct {
	CourseId string `json:"courseId"`
}

type getCoursesListResponse struct {
	Results    []models.CourseHTTP `json:"results"`
	Pagination struct {
		Next     string      `json:"next"`
		Previous interface{} `json:"previous"`
		Count    int         `json:"count"`
		NumPages int         `json:"num_pages"`
	} `json:"pagination"`
}

type getEnrollmentsResponse struct {
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Created  time.Time `json:"created"`
		Mode     string    `json:"mode"`
		IsActive bool      `json:"is_active"`
		User     string    `json:"user"`
		CourseID string    `json:"course_id"`
	} `json:"results"`
}

func (h *Handler) InitCourseRoutes(router *gin.Engine) {
	course := router.Group("/course")
	{
		course.POST("/createCourse/:courseId", h.CreateCourse)
		course.GET("/getCourseContent/:courseId", h.GetCourseContent)
		course.GET("/getCoursesByUser", h.GetCoursesByUser)
		course.GET("/getAllPublicCourses/:pageNumber", h.GetAllPublicCourses)
		course.GET("/getEnrollments/:username", h.GetEnrollments)
		course.PUT("/updateCourse", h.UpdateCourse)
		course.DELETE("/deleteCourse/:courseId", h.DeleteCourse)
	}
}

func (h *Handler) UpdateCourse(c *gin.Context) {
	fmt.Println("Update Course")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	courseHTTP := models.CourseHTTP{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &courseHTTP)
	fmt.Println(courseHTTP)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = h.coursesDelegate.UpdateCourse(&courseHTTP)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CreateCourse(c *gin.Context) {
	fmt.Println("Create Course")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	courseId := c.Param("courseId")
	courseHTTP := models.CourseHTTP{}
	courseId, err := h.coursesDelegate.CreateCourse(&courseHTTP, courseId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, testCourseResponse{
		courseId,
	})
}

func (h *Handler) GetCourseContent(c *gin.Context) {
	fmt.Println("Get Course Content")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	courseId := c.Param("courseId")
	body, err := h.coursesDelegate.GetCourseContent(courseId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	courseHTTP := &models.CourseHTTP{}
	log.Println(courseHTTP)
	err = json.Unmarshal(body, courseHTTP)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, courseHTTP)
}

func (h *Handler) GetCoursesByUser(c *gin.Context) {
	fmt.Println("Get courses For User")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	body, err := h.coursesDelegate.GetCoursesByUser()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var coursesHTTP getCoursesListResponse
	err = json.Unmarshal(body, &coursesHTTP)
	log.Println(coursesHTTP)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, coursesHTTP)
}

func (h *Handler) GetAllPublicCourses(c *gin.Context) {
	fmt.Println("Get All Public courses")
	pN := c.Param("pageNumber")
	pageNumber, err := strconv.Atoi(pN)
	if err != nil {
		log.Println("Nit number in url")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	body, err := h.coursesDelegate.GetAllPublicCourses(pageNumber)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var coursesHTTP getCoursesListResponse
	err = json.Unmarshal(body, &coursesHTTP)
	log.Println(coursesHTTP)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, coursesHTTP)
}

func (h *Handler) GetEnrollments(c *gin.Context) {
	fmt.Println("Get Enrollments")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	username := c.Param("username")

	body, err := h.coursesDelegate.GetEnrollments(username)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var enrollmentsHTTP getEnrollmentsResponse
	err = json.Unmarshal(body, &enrollmentsHTTP)
	log.Println(enrollmentsHTTP)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, enrollmentsHTTP)
}

func (h *Handler) DeleteCourse(c *gin.Context) {
	fmt.Println("Delete Course")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	courseId := c.Param("courseId")
	err := h.coursesDelegate.DeleteCourse(courseId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
