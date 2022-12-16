package handlers

import (
	"fmt"
	"net/http"

	"github.com/Marif226/buysell/internal/database"
)

// Repository is the repository type
type Repository struct {
	DB *database.Database
}

var Repo *Repository

// NewRepo creates a new repository
func NewRepo(db *database.Database) *Repository {
	return &Repository{
		DB: db,
	}
}

func SetRepo(r *Repository) {
	Repo = r
}

func (repo *Repository) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "i get")
}

func (repo *Repository) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "i create")
}

func (repo *Repository) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "i delete")
}