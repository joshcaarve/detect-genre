package main

import (
	"detect-genre/backend/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/api/books", routes.FindBooks)
	r.GET("/api/tree", routes.GetTree)
	r.POST("api/file", routes.SaveFileHandler)
	r.Run()

}
