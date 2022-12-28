package models

import (
	"time"
)

// Product is the product model
type Product struct {
	ID          uint		`json:"id"`
	Title       string		`json:"title"`
	Description string		`json:"description"`
	Price       uint		`json:"price"`
	Author      User		`json:"author"`
	Location    string		`json:"location"`
	CreatedAt   time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

// User is the user model
type User struct {
	ID        uint		`json:"id"`
	FirstName string	`json:"first_name"`
	LastName  string	`json:"last_name"`
	Email     string	`json:"email"`
	Password  string	`json:"password"`
}