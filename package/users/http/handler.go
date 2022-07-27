package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	authDelegate  auth.Delegate
	usersDelegate users.Delegate
}

func NewUsersHandler(authDelegate auth.Delegate, usersDelegate users.Delegate) Handler {
	return Handler{
		authDelegate:  authDelegate,
		usersDelegate: usersDelegate,
	}
}

type getParentResponse struct {
	Parent models.ParentHTTP `json:"parent"`
}

type getTeacherResponse struct {
	Teacher models.TeacherHTTP `json:"teacher"`
}

type getUnitAdminResponse struct {
	UnitAdmin models.UnitAdminHTTP `json:"unitAdmin"`
}

type getStudentResponse struct {
	Student models.StudentHTTP `json:"student"`
}

type getSuperAdminResponse struct {
	SuperAdmin models.SuperAdminHTTP `json:"superAdmin"`
}

type getFreeListener struct {
	FreeListener models.FreeListenerHttp `json:"freeListener"`
}

type loginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) InitUsersRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/", h.GetUser)

		users.POST("/student", h.CreateStudent)
		users.DELETE("/student/:studentId", h.DeleteStudent)
		users.GET("/student/:studentId", h.GetStudentById)
		users.PUT("/student", h.UpdateStudent)

		users.POST("/teacher", h.CreateTeacher)
		users.DELETE("/teacher/:teacherId", h.DeleteTeacher)
		users.PUT("/teacher", h.UpdateTeacher)
		users.GET("/teacher/:teacherId", h.GetTeacherById)

		users.POST("/parent", h.CreateParent)
		users.GET("/parent/:parentId", h.GetParentById)
		users.PUT("/parent", h.UpdateParent)
		users.DELETE("/parent/:parentId", h.DeleteParent)

		users.POST("/freeListener", h.CreateFreeListener)
		users.DELETE("/freeListener/:freeListenerId", h.DeleteFreeListener)
		users.PUT("/freeListener", h.UpdateFreeListener)
		users.GET("/freeListener/:freeListenerId", h.GetFreeListenerById)

		users.POST("/unitAdmin", h.CreateUnitAdmin)
		users.DELETE("/unitAdmin/:unitAdminId", h.DeleteUnitAdmin)
		users.PUT("/unitAdmin", h.UpdateUnitAdmin)
		users.GET("/unitAdmin/:unitAdminId", h.GetUnitAdminByID)

		users.GET("/superAdmin/:superAdminId", h.GetSuperAdminById)
	}
}

func (h *Handler) GetUser(c *gin.Context) {
	fmt.Println("GetUser")
	id, role, err := h.userIdentity(c)
	userId, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	switch role {
	case models.Student:
		student, getStudentErr := h.usersDelegate.GetStudentById(uint(userId))
		if getStudentErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, student)
	case models.Teacher:
		teacher, getTeacherErr := h.usersDelegate.GetTeacherById(uint(userId))
		if getTeacherErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, teacher)
		return
	case models.Parent:
		parent, getParentErr := h.usersDelegate.GetParentById(uint(userId))
		if getParentErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, parent)
		return
	case models.FreeListener:
		freeListener, getFreeListenerErr := h.usersDelegate.GetFreeListenerById(uint(userId))
		if getFreeListenerErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, freeListener)
	case models.UnitAdmin:
		unitAdmin, getUnitAdminErr := h.usersDelegate.GetUnitAdminById(uint(userId))
		if getUnitAdminErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, unitAdmin)
		return
	case models.SuperAdmin:
		superAdmin, getSuperAdminErr := h.usersDelegate.GetSuperAdminById(uint(userId))
		if getSuperAdminErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, superAdmin)
		return
	}
}

func (h *Handler) GetStudentById(c *gin.Context) {
	fmt.Println("Get Student By Id")
	param := c.Param("studentId")
	studentId, _ := strconv.Atoi(param)

	student, err := h.usersDelegate.GetStudentById(uint(studentId))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getStudentResponse{
		student,
	})
}

func (h *Handler) CreateStudent(c *gin.Context) {
	fmt.Println("Create Student")

	studentHttp := &models.StudentHTTP{}

	if err := c.BindJSON(studentHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	studentId, err := h.usersDelegate.CreateStudent(studentHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student": studentId,
	})
}

func (h *Handler) UpdateStudent(c *gin.Context) {
	fmt.Println("Update Student")
	studentHTTP := models.StudentHTTP{}
	if err := c.BindJSON(studentHTTP); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := h.usersDelegate.UpdateStudent(&studentHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteStudent(c *gin.Context) {
	fmt.Println("Delete Student")

	studentId := c.Param("studentId")
	id, _ := strconv.Atoi(studentId)
	err := h.usersDelegate.DeleteStudent(uint(id))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) CreateTeacher(c *gin.Context) {
	fmt.Println("Create Teacher")

	teacherHttp := &models.TeacherHTTP{}

	if err := c.BindJSON(teacherHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	teacherId, err := h.usersDelegate.CreateTeacher(teacherHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teacher": teacherId,
	})

}

func (h *Handler) DeleteTeacher(c *gin.Context) {
	fmt.Println("Delete Teacher")

	teacherId := c.Param("teacherId")
	id, _ := strconv.Atoi(teacherId)
	err := h.usersDelegate.DeleteTeacher(uint(id))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetTeacherById(c *gin.Context) {
	fmt.Println("Get Teacher By Id")
	param := c.Param("teacherId")
	teacherId, _ := strconv.Atoi(param)

	teacher, err := h.usersDelegate.GetTeacherById(uint(teacherId))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getTeacherResponse{
		teacher,
	})
}

func (h *Handler) UpdateTeacher(c *gin.Context) {
	fmt.Println("Update Teacher")
	teacherHTTP := models.TeacherHTTP{}

	if err := c.BindJSON(teacherHTTP); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.usersDelegate.UpdateTeacher(&teacherHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetParentById(c *gin.Context) {
	fmt.Println("Get Parent By Id")
	id := c.Param("parentId")
	parentId, _ := strconv.Atoi(id)

	parent, err := h.usersDelegate.GetParentById(uint(parentId))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getParentResponse{
		parent,
	})
}

func (h *Handler) CreateParent(c *gin.Context) {
	fmt.Println("Create Parent")

	parentHttp := &models.ParentHTTP{}

	if err := c.BindJSON(parentHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	parentId, err := h.usersDelegate.CreateParent(parentHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"parent": parentId,
	})

}

func (h *Handler) DeleteParent(c *gin.Context) {
	fmt.Println("Delete Parent")

	parentId := c.Param("parentId")
	id, _ := strconv.Atoi(parentId)
	err := h.usersDelegate.DeleteParent(uint(id))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) UpdateParent(c *gin.Context) {
	fmt.Println("Update Parent")
	parentHTTP := models.ParentHTTP{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &parentHTTP)
	fmt.Println(parentHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = h.usersDelegate.UpdateParent(&parentHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetFreeListenerById(c *gin.Context) {
	fmt.Println("Get FreeListener By Id")
	id := c.Param("freeListenerId")
	freeListenerId, _ := strconv.Atoi(id)

	freeListener, err := h.usersDelegate.GetFreeListenerById(uint(freeListenerId))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getFreeListener{
		freeListener,
	})
}

func (h *Handler) CreateFreeListener(c *gin.Context) {
	fmt.Println("Create FreeListener")

	freeListenerHttp := &models.FreeListenerHttp{}

	if err := c.BindJSON(freeListenerHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	freeListenerId, err := h.usersDelegate.CreateFreeListener(freeListenerHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"freeListenerId": freeListenerId,
	})

}

func (h *Handler) DeleteFreeListener(c *gin.Context) {
	fmt.Println("Delete Free Listener")

	freeListenerId := c.Param("freeListenerId")
	id, _ := strconv.Atoi(freeListenerId)
	err := h.usersDelegate.DeleteFreeListener(uint(id))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) UpdateFreeListener(c *gin.Context) {
	fmt.Println("Update Free Listener")
	freeListenerHTTP := models.FreeListenerHttp{}

	if err := c.BindJSON(freeListenerHTTP); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.usersDelegate.UpdateFreeListener(&freeListenerHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetUnitAdminByID(c *gin.Context) {
	fmt.Println("Get Unit Admin By ID")
	id := c.Param("unitAdminId")
	unitAdminId, _ := strconv.Atoi(id)

	unitAdmin, err := h.usersDelegate.GetUnitAdminById(uint(unitAdminId))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getUnitAdminResponse{
		unitAdmin,
	})
}

func (h *Handler) UpdateUnitAdmin(c *gin.Context) {
	fmt.Println("Update Unit Admin")
	unitAdminHTTP := models.UnitAdminHTTP{}

	if err := c.BindJSON(unitAdminHTTP); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.usersDelegate.UpdateUnitAdmin(&unitAdminHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CreateUnitAdmin(c *gin.Context) {
	fmt.Println("Create Unit Admin")

	unitAdminHttp := &models.UnitAdminHTTP{}
	if err := c.BindJSON(unitAdminHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	unitAdminId, err := h.usersDelegate.CreateUnitAdmin(unitAdminHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"unitAdmin": unitAdminId,
	})
}

func (h *Handler) DeleteUnitAdmin(c *gin.Context) {
	fmt.Println("Delete Unit Admin")

	unitAdminId := c.Param("unitAdminId")
	id, _ := strconv.Atoi(unitAdminId)
	err := h.usersDelegate.DeleteUnitAdmin(uint(id))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) GetSuperAdminById(c *gin.Context) {
	fmt.Println("Get Super Admin By Id")
	param := c.Param("superAdminId")
	superAdminId, _ := strconv.Atoi(param)

	superAdmin, err := h.usersDelegate.GetSuperAdminById(uint(superAdminId))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getSuperAdminResponse{
		superAdmin,
	})
}
