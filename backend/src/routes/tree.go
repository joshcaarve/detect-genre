package routes

import (
	"detect-genre/backend/src/tree"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTree /return json of tree
func GetTree(c *gin.Context) {
	root := tree.NewTree("Electronic")
	dnB := tree.NewTree("DnB")
	dnB.AddChild("Deep")
	dnB.AddChild("Jungle")

	house := tree.NewTree("House")
	house.AddChild("Deep")
	house.AddChild("BigRoom")

	dubstep := tree.NewTree("Dubstep")

	root.AddChild(dnB)
	root.AddChild(house)
	root.AddChild(dubstep)
	c.JSON(http.StatusOK, gin.H{"data": root})
}
