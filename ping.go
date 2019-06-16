package main
import "github.com/gin-gonic/gin"


// Pong return message
func Pong(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}