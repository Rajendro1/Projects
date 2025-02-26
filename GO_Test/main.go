package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/database"
	"main.go/users"
)

func main() {
	database.InitDB()
	userService := &users.UserService{DB: database.DB}

	r := mux.NewRouter()
	r.HandleFunc("/users", userService.CreateUser).Methods("POST")

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
