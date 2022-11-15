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
		course.POST("/:coursePacketId", h.CreateCoursePacket)
		course.GET("/:coursePacketId", h.GetCoursePacketById)
		course.GET("/", h.GetAllCoursePackets)
		course.PUT("/", h.UpdateCoursePacket)
		course.DELETE("/:coursePacketId", h.DeleteCoursePacket)
	}
}

func (h *Handler) UpdateCoursePacket(c *gin.Context) {
	log.Println("Update Course Packet")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
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
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CreateCoursePacket(c *gin.Context) {
	log.Println("Create Course Packet")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
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
		log.Println(userIdentityErr)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
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
	log.Println("Get all CoursePackets")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	coursePackets, err := h.coursePacketDelegate.GetAllCoursePackets()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, coursePackets)
}

func (h *Handler) DeleteCoursePacket(c *gin.Context) {
	log.Println("Delete Course Packet")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	courseId := c.Param("coursePacketId")
	err := h.coursePacketDelegate.DeleteCoursePacket(courseId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
