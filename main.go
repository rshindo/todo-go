package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"strconv"
)

var db *gorm.DB

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect db")
	}
	defer db.Close()

	db.AutoMigrate(&Todo{})

	r := gin.Default()
	r.GET("/ping", Pong)
	r.POST("/todo", func(c *gin.Context) {
		var json Todo
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&json)
		c.JSON(http.StatusCreated, json)
	})
	r.GET("/todo", func(c *gin.Context) {
		var todos []Todo
		db.Find(&todos)
		c.JSON(http.StatusOK, gin.H{"todos": todos})
	})
	r.GET("/todo/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var todo Todo
		recordNotFound := db.First(&todo, id).RecordNotFound()
		if recordNotFound == true {
			c.JSON(http.StatusNotFound, gin.H{"error": "ID not found."})
			return
		}
		c.JSON(http.StatusOK, todo)
	})
	r.Run()
}
