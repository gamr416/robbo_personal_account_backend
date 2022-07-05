package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"net/http"
	"strconv"
)

type Handler struct {
	authDelegate    auth.Delegate
	coursesDelegate courses.Delegate
}

func NewCoursesHandler(authDelegate auth.Delegate, coursesDelegate courses.Delegate) Handler {
	return Handler{
		authDelegate:    authDelegate,
		coursesDelegate: coursesDelegate,
	}
}

type createCourseResponse struct {
	courseId string `json:"courseId"`
}

func (h *Handler) InitCourseRoutes(router *gin.Engine) {
	course := router.Group("/course")
	{
		course.POST("/createCourse/:courseId", h.CreateCourse)
		course.GET("/getCourseContent/:courseId", h.GetCourseContent)
		course.GET("/getCoursesByUser/:user", h.GetCoursesByUser)
		course.GET("/getAllPublicCourses/:pageNumber", h.GetAllPublicCourses)
		course.PUT("/updateCourse/:courseId", h.UpdateCourse)
		//course.GET("/deleteCourse/:courseId", h.DeleteCourse)
	}
}

func (h *Handler) CreateCourse(c *gin.Context) {
	fmt.Println("Create Course")

	courseId := c.Param("courseId")

	courseHTTP := models.CourseHTTP{}

	courseId, err := h.coursesDelegate.CreateCourse(&courseHTTP, courseId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, createCourseResponse{
		courseId,
	})

}

func (h *Handler) GetAllPublicCourses(c *gin.Context) {
	pageNumber, err := strconv.Atoi(c.Param("pageNumber"))
	body, err := h.coursesDelegate.GetAllPublicCourses(pageNumber)
	s2, _ := json.Marshal(body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, s2)
}

func (h *Handler) GetCourseContent(c *gin.Context) {
	fmt.Println("Get Course Content")
	courseId := c.Param("courseId")
	body, err := h.coursesDelegate.GetCourseContent(courseId)
	s2, _ := json.Marshal(body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, s2)
}

func (h *Handler) GetCoursesByUser(c *gin.Context) {
	fmt.Println("Get Courses By User")
	username := c.Param("user")
	body, err := h.coursesDelegate.GetCoursesByUser(username)
	s2, _ := json.Marshal(body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, s2)
}

func (h *Handler) UpdateCourse(c *gin.Context) {

}

func (h *Handler) DeleteCourse(c *gin.Context) {
	fmt.Println("Delete Course")

	courseId := c.Param("courseId")

	courseHTTP := models.CourseHTTP{}
	courseHTTP.CourseID = courseId

	err := h.coursesDelegate.DeleteCourse(&courseHTTP)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
