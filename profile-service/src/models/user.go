package models

import (
	"time"
)

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Number    string    `json:"number" gorm:"unique"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func NewUser(args map[string]interface{}) *User {
	return &User{
		Name:   args["name"].(string),
		Age:    int(args["age"].(float64)),
		Number: args["number"].(string),
	}
}
