package database

import (
	"genre-tree/backend/src/adjlist"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB() {
	root := adjlist.NewEntry("Electronic", 0)
	house := root.AddChild("House")
	dnB := root.AddChild("DnB")
	deep := dnB.AddChild("Deep")
	wub := deep.AddChild("wub")
	dub := root.AddChild("Dubstep")
	jung := dnB.AddChild("Jungle")

	DB.Create(&root)
	DB.Create(&house)
	DB.Create(&dnB)
	DB.Create(&deep)
	DB.Create(&wub)
	DB.Create(&dub)
	DB.Create(&jung)
}

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&adjlist.AdjList{})

	DB = database

}
