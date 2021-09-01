package adjlist

import (
	"fmt"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type AdjListEntry struct {
	Name       string `json:"genre" binding:"required"`
	ParentName string `json:"parentGenre" binding:"required"`
}

type AdjList struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"genre"`
	Parent uint   `json:"pid"`
}

var globalID uint = 0

func NewEntry(name string, parentID uint) *AdjList {
	globalID += 1
	return &AdjList{
		ID:     globalID,
		Name:   name,
		Parent: parentID,
	}
}

func (parent *AdjList) AddChild(childName string) *AdjList {
	globalID += 1
	return &AdjList{
		ID:     globalID,
		Name:   childName,
		Parent: parent.ID,
	}
}
func (adjList *AdjList) String() string {
	return fmt.Sprint("Name: ", adjList.Name, " ID: ", adjList.ID, " PID: ", adjList.Parent)
}
