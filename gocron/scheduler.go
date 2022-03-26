package gocron

import (
	"GoTools/gomail"
	"GoTools/model"
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func Scheduler(user model.User) {
	scheduler := gocron.NewScheduler()

	counter := 0
	scheduler.Every(10).Second().Do(func() {
		counter++
		fmt.Print("Testing ke ", counter, ": ")
		gomail.SendEmail(user)
	})
	<-scheduler.Start()
}
