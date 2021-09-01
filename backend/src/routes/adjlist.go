package routes

import (
	"genre-tree/backend/src/adjlist"
	"genre-tree/backend/src/database"
	_ "genre-tree/backend/src/tree"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	/*
		// Validate input
		var input CreateBookInput
		if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
		}

		// Create book
		book := models.Book{Title: input.Title, Author: input.Author}
		models.DB.Create(&book)

		c.JSON(http.StatusOK, gin.H{"data": book})
	*/
}

func GetAdjlist(c *gin.Context) {
	var adjList []adjlist.AdjList
	database.DB.Find(&adjList)
	c.JSON(http.StatusOK, gin.H{"data": adjList})
}
