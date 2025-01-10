package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

func (user *User) Create(db *gorm.DB) error {
	return db.Create(&user).Error
}

//var users = []User{}
//
//func AddUser(user User) {
//	users = append(users, user)
//}

func FindUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	return &user, err
}
