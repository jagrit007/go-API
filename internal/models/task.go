package models

import (
	"errors"
	"strconv"
)

type Task struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Completed   bool   `json:"completed"`
	UserID      uint   `json:"user_id"`
}

// mock Task DB
var tasks = []Task{}
var taskCounter uint = 1

func AddTask(task Task) {
	task.ID = taskCounter
	taskCounter++
	tasks = append(tasks, task)
}

func FindTaskByUserID(userID uint) []Task {
	var userTasks []Task
	for _, task := range tasks {
		if task.UserID == userID {
			userTasks = append(userTasks, task)
		}
	}
	return userTasks
}

func FindTaskByID(taskID uint, userID uint) (*Task, error) {
	for _, task := range tasks {
		if task.ID == taskID && task.UserID == userID {
			return &task, nil
		}
	}
	return nil, errors.New("Task with ID:" + strconv.Itoa(int(taskID)) + " not found")
}

func UpdateTask(updatedTask Task, userID uint) error {
	for i, task := range tasks {
		if task.ID == updatedTask.ID && task.UserID == userID {
			tasks[i] = updatedTask
			return nil
		}
	}
	return errors.New("task not found or not authorized")
}

func DeleteTask(taskID uint, userID uint) error {
	for i, task := range tasks {
		if task.ID == taskID && task.UserID == userID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found or not authorized")
}
