package models

import "time"

// Product is the product model
type Product struct {
	ID          int
	Title       string
	Description string
	Price       int
	Author      User
	Location    string
	CreatedAt   time.Time
	UpdatedAt	time.Time
}

// User is the user model
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}