package routes

import (
	"detect-genre/backend/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindBooks /books
func FindBooks(c *gin.Context) {
	var books []models.Book
	c.JSON(http.StatusOK, gin.H{"data": books})
}
