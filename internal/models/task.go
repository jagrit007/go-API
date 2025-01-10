package models

import (
	"gorm.io/gorm"
)

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Completed   bool   `json:"completed"`
	UserID      uint   `json:"user_id"`
}

func (task *Task) Create(db *gorm.DB) error {
	return db.Create(&task).Error
}

// mock Task DB
//var tasks = []Task{}
//var taskCounter uint = 1
//
//func AddTask(task Task) {
//	task.ID = taskCounter
//	taskCounter++
//	tasks = append(tasks, task)
//}

func FindTaskByUserID(db *gorm.DB, userID uint, page, limit int) ([]Task, error) {
	var userTasks []Task
	offset := (page - 1) * limit
	query := db.Where("user_id = ?", userID)

	err := query.Offset(offset).Limit(limit).Find(&userTasks).Error

	if err != nil {
		return nil, err
	}

	return userTasks, err
}

func FindTaskByID(db *gorm.DB, taskID uint, userID uint) (*Task, error) {
	var task Task
	err := db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (task *Task) Update(db *gorm.DB) error {
	return db.Save(&task).Error
}

func (task *Task) Delete(db *gorm.DB) error {
	return db.Delete(&task).Error
}
