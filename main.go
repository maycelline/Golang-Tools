package main

import (
	scheduler "GoTools/gocron"
	mail "GoTools/gomail"
	cache "GoTools/goredis"
	asynchronous "GoTools/goroutine"
	model "GoTools/model"
)

func main() {
	var user model.User
	user.FullName = "Maycelline Selvyanti"
	user.Email = "maycelinesudarsono@gmail.com"

	// GoMail
	mail.SendEmail(user)

	// GoRoutine
	asynchronous.DoAsynchronousTask()

	// GoCron
	scheduler.Schedule(user)

	// GoRedis
	cache.GetUser(user)

}
