package repository

import "github.com/Marif226/buysell/internal/models"

type DatabaseRepo interface {
	CreateUser(models.User) error
	GetUserById(id uint) (models.User, error)
	// UpdateUser()
	DeleteUser(id uint) error

	CreateProduct()
	GetProduct()
	// UpdateProduct()
	DeleteProduct()
}