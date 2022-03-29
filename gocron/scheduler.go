package gocron

import (
	"GoTools/gomail"
	"GoTools/model"
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func Schedule(user model.User, scheduler *gocron.Scheduler) {
	counter := 0
	scheduler.Every(10).Second().Do(func() {
		counter++
		fmt.Print("Testing ke ", counter, ": ")
		gomail.SendEmail(user)
	})
	<-scheduler.Start()
}

func StopSchedule(s *gocron.Scheduler) {
	s.Clear()
	fmt.Println("All task removed")
}
