package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

var users = []User{}

func AddUser(user User) {
	users = append(users, user)
}

func FindUserByEmail(email string) (*User, error) {
	for _, user := range users {
		fmt.Printf(email, user.Email)
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
