package dbrepo

import (
	"fmt"

	"github.com/Marif226/buysell/internal/models"
)

func (m *myDbRepo) CreateUser(user models.User) error {
	err := m.DB.Write("users", fmt.Sprintf("user%v", user.ID), &user)

	return err
}

func (m *myDbRepo) GetUserById(id uint) (models.User, error) {
	var user models.User

	err := m.DB.Read("users", fmt.Sprintf("user%v", id), &user)

	return user, err
}

func (m *myDbRepo) DeleteUser(id uint) error {
	err := m.DB.Delete("users", fmt.Sprintf("user%v", id))

	return err
}

func (m *myDbRepo) CreateProduct() {

}

func (m *myDbRepo) GetProduct() {

}

func (m *myDbRepo) DeleteProduct() {}