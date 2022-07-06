package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"log"
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
		course.GET("/createCourse/:courseId", h.CreateCourse)
		course.GET("/getCourseContent/:courseId", h.GetCourseContent)
		course.GET("/getCoursesForUser", h.GetCoursesForUser)
		course.GET("/getAllPublicCourses/:pageNumber", h.GetAllPublicCourses)
		course.PUT("/updateCourse/:courseId", h.UpdateCourse)
		//course.GET("/deleteCourse/:courseId", h.DeleteCourse)
	}
}

func (h *Handler) CreateCourse(c *gin.Context) {
	fmt.Println("Create Course")

	courseId := c.Param("courseId")

	courseHTTP := models.CourseHTTP{}

	courseId, statusCode, err := h.coursesDelegate.CreateCourse(&courseHTTP, courseId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(statusCode)
		return
	}

	c.JSON(http.StatusOK, createCourseResponse{
		courseId,
	})

}

func (h *Handler) GetAllPublicCourses(c *gin.Context) {
	fmt.Println("Get all public courses")
	pageNumber, err := strconv.Atoi(c.Param("pageNumber"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	body, statusCode, err := h.coursesDelegate.GetAllPublicCourses(pageNumber)
	fmt.Println(body)
	//log.Println(body)
	if err != nil {
		fmt.Println("check")
		c.AbortWithStatus(statusCode)
		return
	}

	respBody, err := json.Marshal(body)
	if err != nil {
		c.AbortWithStatus(statusCode)
		return
	}
	c.JSON(http.StatusOK, respBody)
}

func (h *Handler) GetCourseContent(c *gin.Context) {
	fmt.Println("Get Course Content")
	courseId := c.Param("courseId")
	body, statusCode, err := h.coursesDelegate.GetCourseContent(courseId)
	if err != nil {
		c.AbortWithStatus(statusCode)
		return
	}
	log.Println(body)
	respBody, err := json.Marshal(body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err != nil {
		c.AbortWithStatus(statusCode)
		return
	}
	c.JSON(http.StatusOK, respBody)
}

func (h *Handler) GetCoursesForUser(c *gin.Context) {
	fmt.Println("Get Courses For User")
	body, statusCode, err := h.coursesDelegate.GetCoursesForUser()
	if err != nil {
		c.AbortWithStatus(statusCode)
		return
	}
	log.Println(body)
	respBody, err := json.Marshal(body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, respBody)
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
