package controller

import (
	scheduler "GoTools/gocron"
	mail "GoTools/gomail"
	cache "GoTools/goredis"
	model "GoTools/model"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	users = cache.GetUsers()

	if users == nil {
		db := connect()
		defer db.Close()

		query := "SELECT id, fullname, email, password from users"

		rows, err := db.Query(query)
		if err != nil {
			log.Println(err)
			showMessage(w, 400, "Query Error")
			return
		}

		var user model.User
		for rows.Next() {
			if err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.Password); err != nil {
				log.Println(err.Error())
				showMessage(w, 400, "Get Failed")
				return
			} else {
				users = append(users, user)
			}
		}
		go cache.SetUsers(users)
	}

	showUsersSuccessMessage(w, 200, "Get Success", users)
}

func InsertNewUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		showMessage(w, 400, "Parse Error")
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
		showMessage(w, 400, "Query Error")
		return
	}

	id, _ := resultQuery.LastInsertId()
	var user model.User = model.User{Id: int(id), FullName: fullname, Email: email, Password: password}

	go cache.SetUsers(nil)
	go mail.SendEmail(user)

	showUserSuccessMessage(w, 200, "Insert Success", user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		showMessage(w, 400, "Parse Error")
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	var user model.User
	if err := db.QueryRow("SELECT id, fullname, email, password from users where email = ? AND password = ?",
		email, password).Scan(&user.Id, &user.FullName, &user.Email, &user.Password); err != nil {
		log.Println(err.Error())
		showMessage(w, 400, "Login Failed")
		return
	}

	go scheduler.Schedule(user)
	generateToken(w, user)

	showUserSuccessMessage(w, 200, "Login Success", user)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	go scheduler.StopSchedule()
	resetUserToken(w)
	showMessage(w, 200, "Success")
}

func showUserSuccessMessage(w http.ResponseWriter, status int, message string, data model.User) {
	var response model.UserResponse
	response.Status = status
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func showUsersSuccessMessage(w http.ResponseWriter, status int, message string, data []model.User) {
	var response model.UsersResponse
	response.Status = status
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func showMessage(w http.ResponseWriter, status int, message string) {
	var response model.Response
	response.Status = status
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
