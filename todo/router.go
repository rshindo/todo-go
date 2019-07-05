package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rshindo/todo-go/common"
)

func TodoRegister(router *gin.RouterGroup) {
	router.GET("/:id", todoRetrieve)
	router.GET("", todosRetrieve)
	router.POST("", todoCreate)
}

func todoRetrieve(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var todoModel Todo
	db := common.GetDB()
	recordNotFound := db.First(&todoModel, id).RecordNotFound()
	if recordNotFound == true {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID not found."})
		return
	}

	var todo TodoJSON
	modelToJSON(todoModel, &todo)
	c.JSON(http.StatusOK, todo)
}

func todoCreate(c *gin.Context) {
	var json TodoJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := common.GetDB()
	var todoModel Todo
	jsonToModel(json, &todoModel)
	db.Create(&todoModel)

	modelToJSON(todoModel, &json)
	c.JSON(http.StatusCreated, json)
}

func todosRetrieve(c *gin.Context) {
	var todoModels []Todo
	db := common.GetDB()
	db.Find(&todoModels)

	var todos []TodoJSON
	for _, todoModel := range todoModels {
		var todo TodoJSON
		modelToJSON(todoModel, &todo)
		todos = append(todos, todo)
	}
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func modelToJSON(model Todo, todo *TodoJSON) {
	todo.ID = model.ID
	todo.Title = model.Title
	todo.DueDate = model.DueDate
}

func jsonToModel(todo TodoJSON, model *Todo) {
	model.ID = todo.ID
	model.Title = todo.Title
	model.DueDate = todo.DueDate
}
