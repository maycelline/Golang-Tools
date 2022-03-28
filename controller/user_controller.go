package controller

import (
	scheduler "GoTools/gocron"
	mail "GoTools/gomail"
	model "GoTools/model"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT id, fullname, email, password from users"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		responseMessage(w, 150, "Query error")
		return
	}

	var user model.User
	var users []model.User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.Password); err != nil {
			log.Println(err.Error())
			responseMessage(w, 170, "Data error")
			return
		} else {
			users = append(users, user)
		}
	}

	var response model.UsersResponse
	response.Status = 200
	response.Message = "Succes"
	response.Data = users

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertNewUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		responseMessage(w, 100, "Parse error")
		return
	}

	fullname := r.Form.Get("fullname")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	resultQuery, errQuery := db.Exec("INSERT INTO users (fullname, email, password) VALUES (?,?,?)",
		fullname,
		email,
		password,
	)

	if errQuery != nil {
		log.Println(errQuery)
		responseMessage(w, 400, "Query error, Insert failed")
		return
	}

	id, _ := resultQuery.LastInsertId()
	var user model.User = model.User{Id: int(id), FullName: fullname, Email: email, Password: password}

	//asynchronous
	go mail.SendEmail(user)

	var response model.UserResponse
	response.Status = 200
	response.Message = "Succes"
	response.Data = user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		responseMessage(w, 100, "Parse error")
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	var user model.User
	if err := db.QueryRow("SELECT id, fullname, email, password from users where email = ? AND password = ?",
		email, password).Scan(&user.Id, &user.FullName, &user.Email, &user.Password); err != nil {
		log.Println(err.Error())
		responseMessage(w, 170, "Login gagal")
		return
	}

	go scheduler.Schedule(user)
	generateToken(w, user)

	var response model.UserResponse
	response.Status = 200
	response.Message = "Login Success"
	response.Data = user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	go scheduler.StopSchedule()
	resetUserToken(w)
	responseMessage(w, 200, "Success")
}

func responseMessage(w http.ResponseWriter, status int, message string) {
	var response model.Response
	response.Status = status
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
