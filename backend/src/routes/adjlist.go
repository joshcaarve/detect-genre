package routes

import (
	"genre-tree/backend/src/adjlist"
	"genre-tree/backend/src/database"
	_ "genre-tree/backend/src/tree"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
func UpdateAdjListEntry(c *gin.Context) {
	// Get model if exist
	var entry adjlist.AdjList
	if err := database.DB.Where("id = ?", c.Param("id")).First(&entry).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
*/

func isLeafNode(entryID uint) bool {
	var adjList adjlist.AdjList
	if err := database.DB.Where("Parent = ?", entryID).First(&adjList).Error; err != nil {
		return true
	}
	return false
}

func deleteAllBelow(entryID uint) {
	if isLeafNode(entryID) {
		return
	}
	deleteAllBelowHelper(entryID)
}
func deleteAllBelowHelper(entryID uint) {
	if isLeafNode(entryID) {
		var adjList adjlist.AdjList
		database.DB.Where("ID = ?", entryID).First(&adjList)
		database.DB.Delete(&adjList)
	}
	for _, entry := range getChildrenID(entryID) {
		deleteAllBelowHelper(entry)
	}
}

func getChildrenID(entryID uint) []uint {
	var adjList []adjlist.AdjList
	if err := database.DB.Find(&adjList, "Parent = ?", entryID).Error; err != nil {

	}
	var idList []uint
	for _, entry := range adjList {
		idList = append(idList, entry.ID)
	}
	return idList
}

func GetChildrenRequest(c *gin.Context) {
	var adjList []adjlist.AdjList
	if err := database.DB.Find(&adjList, "Parent = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": adjList})
}

func DeleteAdjListEntry(c *gin.Context) {
	var adjList adjlist.AdjList
	if err := database.DB.Where("ID = ?", c.Param("id")).First(&adjList).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	deleteAllBelow(adjList.ID)
	database.DB.Delete(&adjList)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func IsLeafNodeRequest(c *gin.Context) {
	var adjList adjlist.AdjList
	if err := database.DB.Where("Parent = ?", c.Param("id")).First(&adjList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": false})
}

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
