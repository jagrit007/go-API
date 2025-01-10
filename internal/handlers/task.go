package handlers

import (
	"github.com/gin-gonic/gin"
	"go-tasks-app-practice/internal/config"
	"go-tasks-app-practice/internal/models"
	"go-tasks-app-practice/internal/utils"
	"net/http"
	"strconv"
	"time"
)

type CreateTaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Completed   bool   `json:"completed"`
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
		Completed:   input.Completed,
		UserID:      userID,
	}

	dateTime, err := utils.ParseDateTime(input.DueDate)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error in parsing the date."})
		return
	}

	task.DueDate = dateTime.UTC()

	err = task.Create(config.DB)
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

	// get page from query params

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil || limit < 1 {
		limit = 5
	}

	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "asc")

	var completed *bool
	if completedStr := c.Query("completed"); completedStr != "" {
		completedVal, err := strconv.ParseBool(completedStr)
		if err == nil {
			completed = &completedVal
		}
	}

	var dueDate *time.Time
	if dueDateStr := c.Query("dueDate"); dueDateStr != "" {
		parsedDate, err := utils.ParseDateTime(dueDateStr)
		if err == nil {
			dueDate = &parsedDate
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use DD/MM/YYYY."})
			return
		}
	}

	tasks, err := models.FindTaskByUserID(config.DB, userID, page, limit, completed, dueDate, sortBy, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	for i := range tasks {
		tasks[i].DueDate = tasks[i].DueDate.UTC()
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
	task.Completed = input.Completed

	parsedDueDate, err := utils.ParseDateTime(input.DueDate)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "unable to parse date"})
		return
	}

	task.DueDate = parsedDueDate

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
