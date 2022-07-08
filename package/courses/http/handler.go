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
		course.GET("/getCoursesByUser", h.GetCoursesByUser)
		course.GET("/getAllPublicCourses/:pageNumber", h.GetAllPublicCourses)
		course.GET("/getEnrollments/:username", h.GetEnrollments)
		//	course.PUT("/updateCourse/:courseId", h.UpdateCourse)*/
		course.DELETE("/deleteCourse/:courseId", h.DeleteCourse)
	}
}

func (h *Handler) CreateCourse(c *gin.Context) {
	fmt.Println("Create Course")

	courseId := c.Param("courseId")

	courseHTTP := models.CourseHTTP{}

	courseId, err := h.coursesDelegate.CreateCourse(&courseHTTP, courseId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, createCourseResponse{
		courseId,
	})

}

func (h *Handler) GetCourseContent(c *gin.Context) {
	fmt.Println("Get Course Content")
	courseId := c.Param("courseId")
	body, err := h.coursesDelegate.GetCourseContent(courseId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	courseHt := &models.CourseHTTP{}
	err = json.Unmarshal(body, courseHt)

	if err != nil {
		log.Println(err)
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, courseHt)
}
func (h *Handler) GetCoursesByUser(c *gin.Context) {
	fmt.Println("Get Courses For User")
	body, err := h.coursesDelegate.GetCoursesByUser()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
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

func (h *Handler) GetAllPublicCourses(c *gin.Context) {
	fmt.Println("Get All Public Courses")
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
	log.Println(body)
	respBody, err := json.Marshal(body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, respBody)
}

func (h *Handler) GetEnrollments(c *gin.Context) {
	fmt.Println("Get Enrollments")
	username := c.Param("username")

	body, err := h.coursesDelegate.GetEnrollments(username)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
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

func (h *Handler) DeleteCourse(c *gin.Context) {
	fmt.Println("Delete Course")

	courseId := c.Param("courseId")

	err := h.coursesDelegate.DeleteCourse(courseId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
