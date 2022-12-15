package main

import (
	"net/http"

	// "github.com/Marif226/buysell/internal/database"
	"github.com/Marif226/buysell/internal/handlers"
)

func main() {
	// db := database.CreateDB()

	server := &http.Server{
		Addr: ":8080",
		Handler: routes(),
	}

	server.ListenAndServe()
}

func routes() http.Handler {
	router := newPathResolver()

	router.Add("POST /user/create/", handlers.CreateUser)
	router.Add("GET /user/", handlers.GetUser)
	router.Add("DELETE /user/delete/", handlers.DeleteUser)
	
	return router
}