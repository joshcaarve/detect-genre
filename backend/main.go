package main

import (
	"genre-tree/backend/src/database"
	"genre-tree/backend/src/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {
	/*
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
	*/
	/*
		root := adjlist.NewEntry("Electronic", 0)
		house := root.AddChild("House")
		dnB := root.AddChild("DnB")
		deep := dnB.AddChild("Deep")

		fmt.Println(root.String())
		fmt.Println(house.String())
		fmt.Println(dnB.String())
		fmt.Println(deep.String())
		os.Exit(0)
	*/
	r := gin.Default()
	database.ConnectDataBase()
	database.InitDB()
	r.GET("/api/list", routes.GetAdjList)
	r.POST("/api/list", routes.CreateAdjListEntry)
	r.GET("/api/tree", routes.GetTree)
	r.Run()

}
