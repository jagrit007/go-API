package handlers

import (
	"github.com/gin-gonic/gin"
	"go-tasks-app-practice/internal/config"
	"go-tasks-app-practice/internal/models"
	"net/http"
	"strconv"
)

type CreateTaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

type UpdateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Completed   bool   `json:"completed"`
}

func CreateTask(c *gin.Context) {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uint) // will be set by the middleware

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		DueDate:     input.DueDate,
		Completed:   false,
		UserID:      userID,
	}

	err := task.Create(config.DB)
	if err != nil {
		return
	}

	//models.AddTask(task)
	c.JSON(http.StatusCreated, gin.H{"message": "task created successfully", "task": task})

}

func GetTaskByID(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	taskIDStr := c.Param("id")
	taskID, _ := strconv.ParseUint(taskIDStr, 10, 32)

	task, _ := models.FindTaskByID(config.DB, uint(taskID), userID)

	c.JSON(http.StatusOK, gin.H{"message": "task retrieved successfully", "task": task})
}

func GetTasks(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	tasks, err := models.FindTaskByUserID(config.DB, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tasks retrieved successfully", "tasks": tasks})
}

func UpdateTask(c *gin.Context) {
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Invalid task ID"})
		return
	}
	userID := c.MustGet("user_id").(uint)

	task, err := models.FindTaskByID(config.DB, uint(taskID), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	task.Title = input.Title
	task.Description = input.Description
	task.DueDate = input.DueDate
	task.Completed = input.Completed

	if err := task.Update(config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully", "task": task})

}

func DeleteTask(c *gin.Context) {
	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.MustGet("user_id").(uint)

	task, err := models.FindTaskByID(config.DB, uint(taskID), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := task.Delete(config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}
