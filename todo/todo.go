package todo

import "github.com/jinzhu/gorm"

// Todo is todo.
type Todo struct {
	gorm.Model
	Title   string `json:"title" binding:"required"`
	DueDate string `json:"due_date"`
}
