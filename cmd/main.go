package main

import (
	"github.com/gin-gonic/gin"
	"go-tasks-app-practice/internal/config"
	"go-tasks-app-practice/internal/handlers"
	"go-tasks-app-practice/internal/middleware"
	"log"
)

func main() {
	router := gin.Default()

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Error connecting to the Database")
	}
	defer func() {
		db, _ := db.DB()
		_ = db.Close()
	}()

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
		protectedTaskRoutes.GET("/:id", handlers.GetTaskByID)
		protectedTaskRoutes.PUT("/:id", handlers.UpdateTask)
		protectedTaskRoutes.DELETE("/:id", handlers.DeleteTask)

	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error in the server: ", err)
	}
}
