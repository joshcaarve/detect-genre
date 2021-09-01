package adjlist

import (
	"fmt"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type AdjList struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"genre"`
	Parent uint   `json:"pid"`
}

var global_id uint = 0

func NewEntry(name string, parentID uint) *AdjList {
	global_id += 1
	return &AdjList{
		ID:     global_id,
		Name:   name,
		Parent: parentID,
	}
}

func (parent *AdjList) AddChild(childName string) *AdjList {
	global_id += 1
	return &AdjList{
		ID:     global_id,
		Name:   childName,
		Parent: parent.ID,
	}
}
func (adjList *AdjList) String() string {
	return fmt.Sprint("Name: ", adjList.Name, " ID: ", adjList.ID, " PID: ", adjList.Parent)
}
