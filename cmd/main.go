package main

import (
	"github.com/gin-gonic/gin"
	"go-tasks-app-practice/internal/handlers"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"message": "Hello World",
			"status":  200,
		})
	})

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error in the server: ", err)
	}
}
