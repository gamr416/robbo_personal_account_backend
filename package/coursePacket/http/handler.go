package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/coursePacket"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Handler struct {
	authDelegate         auth.Delegate
	coursePacketDelegate coursePacket.Delegate
}

func NewCoursePacketHandler(
	authDelegate auth.Delegate,
	coursePacketDelegate coursePacket.Delegate,
) Handler {
	return Handler{
		authDelegate:         authDelegate,
		coursePacketDelegate: coursePacketDelegate,
	}
}

type testCourseResponse struct {
	CourseId string `json:"courseId"`
}

func (h *Handler) InitCoursePacketRoutes(router *gin.Engine) {
	course := router.Group("/coursePacket")
	{
		course.POST("/createCoursePacket/:coursePacketId", h.CreateCoursePacket)
		course.GET("/getCoursePacket/:coursePacketId", h.GetCoursePacketById)
		course.GET("/getAllCoursePackets/", h.GetAllCoursePackets)
		course.PUT("/updateCoursePacket", h.UpdateCoursePacket)
		course.DELETE("/deleteCoursePacket/:coursePacketId", h.DeleteCoursePacket)
	}
}

func (h *Handler) UpdateCoursePacket(c *gin.Context) {
	fmt.Println("Update Course Packet")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	coursePacketHTTP := models.CoursePacketHTTP{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &coursePacketHTTP)
	fmt.Println(coursePacketHTTP)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = h.coursePacketDelegate.UpdateCoursePacket(&coursePacketHTTP)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CreateCoursePacket(c *gin.Context) {
	fmt.Println("Create Course Packet")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	courseId := c.Param("coursePacketId")
	coursePacketHTTP := models.CoursePacketHTTP{}
	courseId, err := h.coursePacketDelegate.CreateCoursePacket(&coursePacketHTTP, courseId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, testCourseResponse{
		courseId,
	})
}

func (h *Handler) GetCoursePacketById(c *gin.Context) {
	fmt.Println("Get CoursePacket By Id")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	coursePacketId := c.Param("coursePacketId")

	coursePacket, err := h.coursePacketDelegate.GetCoursePacketById(coursePacketId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, coursePacket)
}

func (h *Handler) GetAllCoursePackets(c *gin.Context) {
	fmt.Println("Get all CoursePackets")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	coursePackets, err := h.coursePacketDelegate.GetAllCoursePackets()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, coursePackets)
}

func (h *Handler) DeleteCoursePacket(c *gin.Context) {
	fmt.Println("Delete Course Packet")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	courseId := c.Param("coursePacketId")
	err := h.coursePacketDelegate.DeleteCoursePacket(courseId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
