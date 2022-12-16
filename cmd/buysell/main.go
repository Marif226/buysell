package main

import (
	"net/http"

	"github.com/Marif226/buysell/internal/database"
	"github.com/Marif226/buysell/internal/handlers"
)

var db database.Database

func main() {
	db = database.CreateDB()

	repo := handlers.NewRepo(&db)
	handlers.SetRepo(repo)

	server := &http.Server {
		Addr: ":8080",
		Handler: routes(),
	}

	server.ListenAndServe()
}

func routes() http.Handler {
	router := newPathResolver()

	router.Add("POST /user/create/", handlers.Repo.CreateUser)
	router.Add("GET /user/", handlers.Repo.GetUser)
	router.Add("DELETE /user/delete/", handlers.Repo.DeleteUser)
	
	return router
}