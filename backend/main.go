package main

import (
	"genre-tree/backend/src/database"
	"genre-tree/backend/src/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDataBase()
	database.InitDB()
	r.GET("/api/list", routes.GetAdjList)
	r.GET("/api/list/leaf/:id", routes.IsLeafNodeRequest)
	r.GET("/api/list/children/:id", routes.GetChildrenRequest)
	r.POST("/api/list", routes.CreateAdjListEntry)
	r.DELETE("/api/list/:name", routes.DeleteAdjListEntry)
	r.GET("/api/tree", routes.GetTree)
	r.Run()
}
