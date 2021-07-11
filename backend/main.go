package main

import (
	"detect-genre/backend/src/routes"
	_ "detect-genre/backend/src/routes"
	_ "net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {

	//root.PrettyPrint(os.Stdout, " - ")

	r := gin.Default()

	r.GET("/api/books", routes.FindBooks)
	r.GET("/api/tree", routes.GetTree)
	r.Run()

}
