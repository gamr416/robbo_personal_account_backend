package http

import (
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
)

type Handler struct {
	authDelegate        auth.Delegate
	projectsDelegate    projects.Delegate
	projectPageDelegate projectPage.Delegate
}

func NewProjectPageHandler(authDelegate auth.Delegate, projectsDelegate projects.Delegate, projectPageDelegate projectPage.Delegate) Handler {
	return Handler{
		authDelegate:        authDelegate,
		projectsDelegate:    projectsDelegate,
		projectPageDelegate: projectPageDelegate,
	}
}

func (h *Handler) InitProjectRoutes(router *gin.Engine) {
	project := router.Group("/projectPage")
	{
		project.POST("/", h.CreateProjectPage)
		project.GET("/:projectPageId", h.GetProjectPage)
		project.PUT("/:projectPageId", h.UpdateProjectPage)
		project.DELETE("/", h.DeleteProjectPage)
	}
}

type testResponse struct {
	Id string `json:"id"`
}

type testRequest struct {
	Body string `json:"body"`
}

func (h *Handler) CreateProjectPage(c *gin.Context) {
	/**fmt.Println("Create Project")
	jsonDataBytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	projectHTTP := models.ProjectHTTP{}
	projectHTTP.Json = string(jsonDataBytes)

	projectId, err := h.projectsDelegate.CreateProject(&projectHTTP)
	fmt.Println(projectId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, testResponse{
		Id: projectId,
	})**/
}

func (h *Handler) GetProjectPage(c *gin.Context) {

}

func (h *Handler) UpdateProjectPage(c *gin.Context) {
	/**fmt.Println("Update Project")
	jsonDataBytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	projectId := c.Param("projectId")

	projectHTTP := models.ProjectHTTP{}
	projectHTTP.ID = projectId
	projectHTTP.Json = string(jsonDataBytes)

	err = h.projectsDelegate.UpdateProject(&projectHTTP)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, testResponse{
		Id: "1",
	})**/
}

func (h *Handler) DeleteProjectPage(c *gin.Context) {

}
