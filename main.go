package main

import (
	"GoTools/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", controller.InsertNewUser).Methods("POST")

	router.HandleFunc("/users/login", controller.LoginUser).Methods("POST")
	router.HandleFunc("/users/logout", controller.LogoutUser).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
