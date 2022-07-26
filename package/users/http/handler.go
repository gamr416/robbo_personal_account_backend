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

type getUnitAdminByIdResponse struct {
	UnitAdmin models.UnitAdminHTTP `json:"unit_admin"`
}

type getTeacherByIdResponse struct {
	Teacher models.TeacherHTTP `json:"teacher"`
}

type getParentByIdResponse struct {
	Parent models.ParentHTTP `json:"parent"`
}

type getStudentByIdResponse struct {
	Student models.StudentHTTP `json:"student"`
}

type getParentResponse struct {
	Parent models.ParentHTTP `json:"parent"`
}

type getTeacherResponse struct {
	Teacher models.TeacherHTTP `json:"teacher"`
}

type getUnitAdminResponse struct {
	UnitAdmin models.UnitAdminHTTP `json:"unit_admin"`
}

type getStudentResponse struct {
	Student models.StudentHTTP `json:"student"`
}

type getSuperAdminResponse struct {
	SuperAdmin models.SuperAdminHTTP `json:"super_admin"`
}

type getSuperAdminByIdResponse struct {
	SuperAdmin models.SuperAdminHTTP `json:"super_admin"`
}

type loginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) InitUsersRoutes(router *gin.Engine) {
	users := router.Group("/user")
	{
		//users.POST("/create/unitAdmin", h.CreateUnitAdmin)
		users.DELETE("/delete/unitAdmin/:unitAdminId", h.DeleteUnitAdmin)
		users.PUT("/update/unitAdmin", h.UpdateUnitAdmin)
		users.GET("/get/unitAdmin/:unitAdminId", h.GetUnitAdminByID)
		//users.POST("/create/teacher", h.CreateTeacher)
		users.DELETE("/delete/teacher/:teacherId", h.DeleteTeacher)
		users.PUT("/update/teacher", h.UpdateTeacher)
		users.GET("/get/teacher/:teacherId", h.GetTeacherById)
		//users.POST("/create/parent", h.CreateParent)
		users.GET("/get/parent/:parentId", h.GetParentById)
		users.PUT("/update/parent", h.UpdateParent)
		users.DELETE("/delete/parent/:parentId", h.DeleteParent)
		//users.POST("/login/parent", h.GetParent)
		//users.POST("/login/unitAdmin", h.GetUnitAdmin)
		//users.POST("/login/superAdmin", h.GetSuperAdmin)
		users.GET("/get/superAdmin/:superAdminId", h.GetSuperAdminById)
		//users.POST("/login/teacher", h.GetTeacher)
		//users.POST("/create/student", h.CreateStudent)
		users.DELETE("/delete/student/:studentId", h.DeleteStudent)
		users.GET("/get/student/:studentId", h.GetStudentById)
		users.PUT("/update/student", h.UpdateStudent)
		//users.GET("/login/student", h.GetStudent)
	}
}

func (h *Handler) UpdateStudent(c *gin.Context) {
	fmt.Println("Update Student")
	studentHTTP := models.StudentHTTP{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &studentHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = h.usersDelegate.UpdateStudent(&studentHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
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

	c.JSON(http.StatusOK, getStudentByIdResponse{
		student,
	})
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

	c.JSON(http.StatusOK, getSuperAdminByIdResponse{
		superAdmin,
	})
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

	c.JSON(http.StatusOK, getTeacherByIdResponse{
		teacher,
	})
}

func (h *Handler) GetUnitAdmin(c *gin.Context) {
	fmt.Println("Get Unit Admin")
	loginUserHTTP := loginUser{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &loginUserHTTP)
	fmt.Println(loginUserHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	unitAdmin, err := h.usersDelegate.GetUnitAdmin(loginUserHTTP.Email, loginUserHTTP.Password)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getUnitAdminResponse{
		unitAdmin,
	})
}

func (h *Handler) GetStudent(c *gin.Context) {
	fmt.Println("Get Student")
	loginUserHTTP := loginUser{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &loginUserHTTP)
	fmt.Println(loginUserHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	student, err := h.usersDelegate.GetStudent(loginUserHTTP.Email, loginUserHTTP.Password)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getStudentResponse{
		student,
	})
}

func (h *Handler) GetSuperAdmin(c *gin.Context) {
	fmt.Println("Get Super Admin")
	loginUserHTTP := loginUser{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &loginUserHTTP)
	fmt.Println(loginUserHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	superAdmin, err := h.usersDelegate.GetSuperAdmin(loginUserHTTP.Email, loginUserHTTP.Password)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getSuperAdminResponse{
		superAdmin,
	})
}

func (h *Handler) GetTeacher(c *gin.Context) {
	fmt.Println("Get Teacher")
	loginUserHTTP := loginUser{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &loginUserHTTP)
	fmt.Println(loginUserHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	teacher, err := h.usersDelegate.GetTeacher(loginUserHTTP.Email, loginUserHTTP.Password)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getTeacherResponse{
		teacher,
	})
}

func (h *Handler) GetParent(c *gin.Context) {
	fmt.Println("Get Parent")
	loginUserHTTP := loginUser{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &loginUserHTTP)
	fmt.Println(loginUserHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	parent, err := h.usersDelegate.GetParent(loginUserHTTP.Email, loginUserHTTP.Password)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getParentResponse{
		parent,
	})
}

func (h *Handler) UpdateTeacher(c *gin.Context) {
	fmt.Println("Update Teacher")
	teacherHTTP := models.TeacherHTTP{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &teacherHTTP)
	fmt.Println(teacherHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = h.usersDelegate.UpdateTeacher(&teacherHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
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

	c.JSON(http.StatusOK, getUnitAdminByIdResponse{
		unitAdmin,
	})
}

func (h *Handler) UpdateUnitAdmin(c *gin.Context) {
	fmt.Println("Update Unit Admin")
	unitAdminHTTP := models.UnitAdminHTTP{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &unitAdminHTTP)
	fmt.Println(unitAdminHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = h.usersDelegate.UpdateUnitAdmin(&unitAdminHTTP)
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

	c.JSON(http.StatusOK, getParentByIdResponse{
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
