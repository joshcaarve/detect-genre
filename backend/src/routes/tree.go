package routes

import (
	"genre-tree/backend/src/adjlist"
	"genre-tree/backend/src/database"
	"genre-tree/backend/src/tree"
	"net/http"

	"github.com/gin-gonic/gin"
)

// this function though
func AdjlistToTree() *tree.Tree {
	var root []adjlist.AdjList
	database.DB.Find(&root)

	m := make(map[uint]*tree.Tree)
	for _, entry := range root {
		treeNode := tree.NewTree(entry.Name)
		m[entry.ID] = treeNode
		if value, ok := m[entry.Parent]; ok {
			value.AddChild(treeNode)
		}
	}
	return m[1] // root has ID of 1 and PID of 0 so m[1] gives the root of the tree
}

func GetTree(c *gin.Context) {
	root := AdjlistToTree()
	c.JSON(http.StatusOK, gin.H{"data": root})
}
