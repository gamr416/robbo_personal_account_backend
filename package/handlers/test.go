package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *RequestHandler) test(c *gin.Context) {
	fmt.Println("Test")
	c.Status(http.StatusOK)
}
