package handlers

import (
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "i get")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "i create")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "i delete")
}