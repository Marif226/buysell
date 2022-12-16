package database

import (
	"errors"

	"github.com/Marif226/buysell/internal/models"
)

type Database map[int]models.User

func CreateDB() Database {
	return make(Database)
}

func (db Database) AddUser(newUser models.User) (error) {
	db[newUser.ID] = newUser
	return nil
}

func (db Database) DeleteUser(ID int) (error) {
	if _, ok := db[ID]; ok {
		delete(db, ID)

		return nil
	}
	
	return errors.New("user id is not fount")
}

func (db Database) GetUser(ID int) (error, models.User) {
	if _, ok := db[ID]; ok {
		return nil, db[ID]
	}

	return errors.New("user id is not found"), db[ID]
}