package main

import (
	scheduler "GoTools/gocron"
	// mail "GoTools/gomail"
	model "GoTools/model"
)

func main() {
	var user model.User
	user.FullName = "Maycelline Selvyanti"
	user.Email = "maycelinesudarsono@gmail.com"

	//Call Gomail
	// mail.SendEmail(user)

	//Call GoCron
	scheduler.Scheduler(user)

}
