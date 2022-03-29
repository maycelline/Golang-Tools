package gocron

import (
	"GoTools/gomail"
	"GoTools/model"
	"fmt"

	"github.com/jasonlvhit/gocron"
)

var scheduler = gocron.NewScheduler()

func Schedule(user model.User) {
	scheduler.Every(10).Second().Do(func() {
		gomail.SendEmail(user)
	})
	scheduler.Start()
}

func StopSchedule() {
	scheduler.Clear()
	fmt.Println("All task removed")
}
