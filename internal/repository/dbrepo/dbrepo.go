package dbrepo

import (
	"github.com/Marif226/buysell/internal/repository"
	"github.com/Marif226/buysell/pkg/database"
)

type myDbRepo struct {
	DB *database.Driver
}

func NewMyDbRepo(driver *database.Driver) repository.DatabaseRepo {
	return &myDbRepo{
		DB: driver,
	}
}
