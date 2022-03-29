package main

import (
	"GoTools/controller"
	"fmt"
	"log"
	"net/http"

	// cache "GoTools/goredis"
	// model "GoTools/model"
	// scheduler "GoTools/gocron"

	"github.com/gorilla/mux"
)

func main() {
	// var user model.User
	// user.FullName = "Maycelline Selvyanti"
	// user.Email = "maycelinesudarsono@gmail.com"

	// GoMail
	// mail.SendEmail(user)

	// GoRoutine
	// asynchronous.DoAsynchronousTask()

	// GoCron
	// scheduler.Schedule(user)

	// GoRedis
	// cache.GetUser(user)

	//versi Jon
	router := mux.NewRouter()
	router.HandleFunc("/users", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", controller.InsertNewUser).Methods("POST")
	//rencananya cache dipake buat nyimpen data users, setiap ada new user di append/push

	router.HandleFunc("/users/login", controller.LoginUser).Methods("POST")
	router.HandleFunc("/users/logout", controller.LogoutUser).Methods("POST")
	//nanti yang logout cek posisi nya aja eh lagi login ato ga biar ga sembarangan gitu :")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
