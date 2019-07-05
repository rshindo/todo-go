package todo

import "github.com/jinzhu/gorm"

// Todo is todo.
type Todo struct {
	gorm.Model
	Title   string
	DueDate string
}

// TodoJSON .
type TodoJSON struct {
	ID uint `json:"id"`
	Title string `json:"title" binding:"required"`
	DueDate string `json:"due_date"`
}
