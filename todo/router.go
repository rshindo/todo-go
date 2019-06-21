package todo

import "net/http"
import "strconv"
import "github.com/gin-gonic/gin"
import "github.com/rshindo/todo-go/common"

func TodoRegister(router *gin.RouterGroup) {
	router.GET("/:id", todoRetrieve)
	router.GET("", todosRetrieve)
	router.POST("", todoRetrieve)
}

func todoRetrieve(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var todo Todo
	db := common.GetDB()
	recordNotFound := db.First(&todo, id).RecordNotFound()
	if recordNotFound == true {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID not found."})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func todoCreate(c *gin.Context) {
	var json Todo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := common.GetDB()
	db.Create(&json)
	c.JSON(http.StatusCreated, json)
}

func todosRetrieve(c *gin.Context) {
	var todos []Todo
	db := common.GetDB()
	db.Find(&todos)
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}
