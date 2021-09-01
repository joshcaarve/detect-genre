package routes

import (
	"genre-tree/backend/src/adjlist"
	"genre-tree/backend/src/database"
	_ "genre-tree/backend/src/tree"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAdjListEntry(c *gin.Context) {

	var entry adjlist.AdjListEntry
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var parent adjlist.AdjList
	database.DB.Where("Name = ?", entry.ParentName).First(&parent)

	adjListEntry := adjlist.NewEntry(entry.Name, parent.ID)

	database.DB.Create(&adjListEntry)

	c.JSON(http.StatusOK, gin.H{"data": adjListEntry})
}

func GetAdjList(c *gin.Context) {
	var adjList []adjlist.AdjList
	database.DB.Find(&adjList)
	c.JSON(http.StatusOK, gin.H{"data": adjList})
}
