package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
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

func (h *Handler) InitUsersRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/", h.GetUser)

		users.POST("/student", h.CreateStudent)
		users.DELETE("/student/:studentId", h.DeleteStudent)
		users.GET("/student/:studentId", h.GetStudentById)
		users.GET("/student/search/:studentEmail", h.SearchStudentByEmail)
		users.GET("students/:parentId", h.GetStudentByParentId)
		users.PUT("/student", h.UpdateStudent)

		users.POST("/teacher", h.CreateTeacher)
		users.GET("/teachers", h.GetAllTeachers)
		users.DELETE("/teacher/:teacherId", h.DeleteTeacher)
		users.PUT("/teacher", h.UpdateTeacher)
		users.GET("/teacher/:teacherId", h.GetTeacherById)

		users.POST("/parent", h.CreateParent)
		users.GET("/parent/:parentId", h.GetParentById)
		users.GET("/parents", h.GetAllParent)
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
		users.GET("/unitAdmins", h.GetAllUnitAdmins)

		users.GET("/superAdmin/:superAdminId", h.GetSuperAdminById)
		users.PUT("/superAdmin", h.UpdateSuperAdmin)
		users.DELETE("/superAdmin", h.DeleteSuperAdmin)

		users.POST("/relation", h.CreateRelation)
	}
}

func (h *Handler) GetUser(c *gin.Context) {
	fmt.Println("GetUser")
	userId, role, err := h.userIdentity(c)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	switch role {
	case models.Student:
		student, getStudentErr := h.usersDelegate.GetStudentById(userId)
		if getStudentErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, student)
	case models.Teacher:
		teacher, getTeacherErr := h.usersDelegate.GetTeacherById(userId)
		if getTeacherErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, teacher)
		return
	case models.Parent:
		parent, getParentErr := h.usersDelegate.GetParentById(userId)
		if getParentErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, parent)
		return
	case models.FreeListener:
		freeListener, getFreeListenerErr := h.usersDelegate.GetFreeListenerById(userId)
		if getFreeListenerErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, freeListener)
	case models.UnitAdmin:
		unitAdmin, getUnitAdminErr := h.usersDelegate.GetUnitAdminById(userId)
		if getUnitAdminErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, unitAdmin)
		return
	case models.SuperAdmin:
		superAdmin, getSuperAdminErr := h.usersDelegate.GetSuperAdminById(userId)
		if getSuperAdminErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, superAdmin)
		return
	}
}

func (h *Handler) SearchStudentByEmail(c *gin.Context) {
	fmt.Println("GetStudentByEmail")
	studentEmail := c.Param("studentEmail")

	students, err := h.usersDelegate.SearchStudentByEmail(studentEmail)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *Handler) GetStudentById(c *gin.Context) {
	fmt.Println("Get Student By Id")
	studentId := c.Param("studentId")

	student, err := h.usersDelegate.GetStudentById(studentId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getStudentResponse{
		student,
	})
}

func (h *Handler) GetStudentByParentId(c *gin.Context) {
	fmt.Println("Get Student By Parent Id")
	parentId := c.Param("parentId")

	students, err := h.usersDelegate.GetStudentByParentId(parentId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, students)
}

type createStudentInput struct {
	Student  *models.StudentHTTP `json:"student"`
	ParentId string              `json:"parentId"`
}

func (h *Handler) CreateStudent(c *gin.Context) {
	fmt.Println("Create Student")

	_, role, userIdentityErr := h.userIdentity(c)
	if role != models.SuperAdmin || userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	input := new(createStudentInput)

	if err := c.BindJSON(input); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(input.ParentId)

	studentId, err := h.usersDelegate.CreateStudent(input.Student, input.ParentId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"studentId": studentId,
	})
}

type updateStudentInput struct {
	Student *models.StudentHTTP `json:"student"`
}

func (h *Handler) UpdateStudent(c *gin.Context) {
	fmt.Println("Update Student")

	userHttp := models.UserHttp{}

	if err := c.BindJSON(&userHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	studentHttp := &models.StudentHTTP{
		UserHttp: userHttp,
	}

	err := h.usersDelegate.UpdateStudent(studentHttp)
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

	userHttp := models.UserHttp{
		Role: uint(models.Student),
	}

	if err := c.BindJSON(&userHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	teacherHttp := &models.TeacherHTTP{
		UserHttp: userHttp,
	}

	teacherId, err := h.usersDelegate.CreateTeacher(teacherHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teacherId": teacherId,
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
	teacherId := c.Param("teacherId")

	teacher, err := h.usersDelegate.GetTeacherById(teacherId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getTeacherResponse{
		teacher,
	})
}

func (h *Handler) GetAllTeachers(c *gin.Context) {
	fmt.Println("Get All Teachers")
	teachers, err := h.usersDelegate.GetAllTeachers()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, teachers)
}

type updateTeacherInput struct {
	Teacher *models.TeacherHTTP `json:"teacher"`
}

func (h *Handler) UpdateTeacher(c *gin.Context) {
	fmt.Println("Update Teacher")
	userHttp := models.UserHttp{}

	if err := c.BindJSON(&userHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	teacherHttp := &models.TeacherHTTP{
		UserHttp: userHttp,
	}

	err := h.usersDelegate.UpdateTeacher(teacherHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetParentById(c *gin.Context) {
	fmt.Println("Get Parent By Id")
	parentId := c.Param("parentId")

	parent, err := h.usersDelegate.GetParentById(parentId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getParentResponse{
		parent,
	})
}

func (h *Handler) GetAllParent(c *gin.Context) {
	fmt.Println("Get All Parents")
	parents, err := h.usersDelegate.GetAllParent()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, parents)
}

func (h *Handler) CreateParent(c *gin.Context) {
	fmt.Println("Create Parent")
	_, role, userIdentityErr := h.userIdentity(c)
	if role != models.SuperAdmin || userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userHttp := models.UserHttp{
		Role: uint(models.Parent),
	}

	if err := c.BindJSON(&userHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	parentHttp := &models.ParentHTTP{
		UserHttp: userHttp,
	}

	parentId, err := h.usersDelegate.CreateParent(parentHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"parentId": parentId,
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

type updateParentInput struct {
	Parent *models.ParentHTTP `json:"parent"`
}

func (h *Handler) UpdateParent(c *gin.Context) {
	fmt.Println("Update Parent")
	userHttp := models.UserHttp{}

	if err := c.BindJSON(&userHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	parentHttp := &models.ParentHTTP{
		UserHttp: userHttp,
	}

	err := h.usersDelegate.UpdateParent(parentHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetFreeListenerById(c *gin.Context) {
	fmt.Println("Get FreeListener By Id")
	freeListenerId := c.Param("freeListenerId")

	freeListener, err := h.usersDelegate.GetFreeListenerById(freeListenerId)

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

	userHttp := models.UserHttp{
		Role: uint(models.FreeListener),
	}

	if err := c.BindJSON(&userHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	freeListener := &models.FreeListenerHttp{
		UserHttp: userHttp,
	}

	freeListenerId, err := h.usersDelegate.CreateFreeListener(freeListener)
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

type updateFreeListenerInput struct {
	FreeListener *models.FreeListenerHttp `json:"freeListener"`
}

func (h *Handler) UpdateFreeListener(c *gin.Context) {
	fmt.Println("Update Free Listener")
	userHttp := models.UserHttp{}

	if err := c.BindJSON(&userHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	freeListenerHttp := &models.FreeListenerHttp{
		UserHttp: userHttp,
	}

	err := h.usersDelegate.UpdateFreeListener(freeListenerHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetUnitAdminByID(c *gin.Context) {
	fmt.Println("Get Unit Admin By ID")
	unitAdminId := c.Param("unitAdminId")

	unitAdmin, err := h.usersDelegate.GetUnitAdminById(unitAdminId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getUnitAdminResponse{
		unitAdmin,
	})
}

func (h *Handler) GetAllUnitAdmins(c *gin.Context) {
	fmt.Println("Get All UnitAdmins")
	unitAdmins, err := h.usersDelegate.GetAllUnitAdmins()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, unitAdmins)
}

type updateUnitAdminInput struct {
	UnitAdmin *models.UnitAdminHTTP `json:"unitAdmin"`
}

func (h *Handler) UpdateUnitAdmin(c *gin.Context) {
	fmt.Println("Update Unit Admin")
	userHttp := models.UserHttp{}

	if err := c.BindJSON(&userHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	unitAdminHttp := &models.UnitAdminHTTP{
		UserHttp: userHttp,
	}

	err := h.usersDelegate.UpdateUnitAdmin(unitAdminHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CreateUnitAdmin(c *gin.Context) {
	fmt.Println("Create Unit Admin")

	userHttp := models.UserHttp{
		Role: uint(models.UnitAdmin),
	}
	if err := c.BindJSON(&userHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	unitAdminHttp := &models.UnitAdminHTTP{
		UserHttp: userHttp,
	}

	unitAdminId, err := h.usersDelegate.CreateUnitAdmin(unitAdminHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"unitAdminId": unitAdminId,
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
	superAdminId := c.Param("superAdminId")

	superAdmin, err := h.usersDelegate.GetSuperAdminById(superAdminId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getSuperAdminResponse{
		superAdmin,
	})
}

type updateSuperAdminInput struct {
	SuperAdmin *models.SuperAdminHTTP `json:"superAdmin"`
}

func (h *Handler) UpdateSuperAdmin(c *gin.Context) {
	fmt.Println("Update Super Admin")
	userHttp := models.UserHttp{}

	if err := c.BindJSON(&userHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	superAdminHTTP := &models.SuperAdminHTTP{
		UserHttp: userHttp,
	}

	err := h.usersDelegate.UpdateSuperAdmin(superAdminHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteSuperAdmin(c *gin.Context) {
	fmt.Println("Delete Super Admin")
	superAdminId := c.Param("AdminId")
	id, _ := strconv.Atoi(superAdminId)
	err := h.usersDelegate.DeleteSuperAdmin(uint(id))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

type createRelation struct {
	ParentId string `json:"parentId"`
	ChildId  string `json:"childId"`
}

func (h *Handler) CreateRelation(c *gin.Context) {
	fmt.Println("CreateRelation")

	createRelationInput := new(createRelation)

	if err := c.BindJSON(createRelationInput); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	createRelationErr := h.usersDelegate.CreateRelation(createRelationInput.ParentId, createRelationInput.ChildId)

	if createRelationErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
