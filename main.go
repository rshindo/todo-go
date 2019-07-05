package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rshindo/todo-go/common"
	"github.com/rshindo/todo-go/todo"
)

var db *gorm.DB

func main() {
	db := common.Init()
	defer common.Close()

	db.AutoMigrate(&todo.Todo{})

	app := NewApp(db)

	app.Run()
}

// NewApp returns new app.
func NewApp(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Pong)

	v1 := r.Group("/api")
	todo.TodoRegister(v1.Group("todo"))

	return r
}
