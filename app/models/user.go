package models

import (
	"time"
)

type User struct {
	// gorm.Model
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MyAnnounces struct {
	ann []int
}

var myAnnounces []MyAnnounces

type MyFinds struct {
	ann []int
}

var myFinds []MyFinds
