package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Marif226/buysell/internal/models"
	"github.com/Marif226/buysell/internal/repository"
)

// Repository is the repository type
type Repository struct {
	DB repository.DatabaseRepo
}

var Repo *Repository

// NewRepo creates a new repository
func NewRepo(db repository.DatabaseRepo) *Repository {
	return &Repository{
		DB: db,
	}
}

// SetRepo sets repository with a database
func SetRepo(r *Repository) {
	Repo = r
}

// GetUser finds user by id
func (repo *Repository) GetUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	id, err := strconv.ParseUint(query.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := repo.DB.GetUserById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, user)
}

// CreateUser creates new user and adds him to the database
func (repo *Repository) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newUser models.User

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = Repo.DB.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(newUser)

	// Repo.DB.GetUserById(newUser.ID)
}


func (repo *Repository) DeleteUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	id, err := strconv.ParseUint(query.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo.DB.DeleteUser(uint(id))
}