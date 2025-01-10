package main

import (
	"github.com/gin-gonic/gin"
	"go-tasks-app-practice/internal/handlers"
	"go-tasks-app-practice/internal/middleware"
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

	// tasks route (protected)
	protectedTaskRoutes := router.Group("/tasks")
	protectedTaskRoutes.Use(middleware.AuthMiddleware())
	{

		protectedTaskRoutes.POST("/", handlers.CreateTask)
		protectedTaskRoutes.GET("/", handlers.GetTasks)
		protectedTaskRoutes.PUT("/:id", handlers.UpdateTask)
		protectedTaskRoutes.DELETE("/:id", handlers.DeleteTask)

	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error in the server: ", err)
	}
}
