package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello /hello
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello"})
}
