package main

import (
	"fmt"
	"net/http"

	"github.com/Marif226/buysell/pkg/database"
	"github.com/Marif226/buysell/internal/handlers"
	"github.com/Marif226/buysell/internal/repository/dbrepo"
)

const portNumber = ":8080"

func main() {
	dbSetup()

	server := &http.Server {
		Addr: portNumber,
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

func dbSetup() {
	driver, err := database.NewDriver("./db")
	if err != nil {
		fmt.Println(err)
	}

	dbRepo := dbrepo.NewMyDbRepo(driver)

	repo := handlers.NewRepo(dbRepo)
	handlers.SetRepo(repo)
}