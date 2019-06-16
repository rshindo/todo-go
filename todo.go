package main

import "github.com/jinzhu/gorm"

// Todo is todo.
type Todo struct {
	gorm.Model
	// ID      uint
	Title   string
	DueDate string
}
