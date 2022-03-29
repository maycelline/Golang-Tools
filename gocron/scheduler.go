package gocron

import (
	"GoTools/gomail"
	"GoTools/model"
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

func Schedule(user model.User) {
	scheduler := gocron.NewScheduler()

	counter := 0
	scheduler.Every(10).Second().Do(func() {
		counter++
		fmt.Print("Testing ke ", counter, ": ")
		gomail.SendEmail(user)
	})
	<-scheduler.Start()
}

func StopSchedule(s *gocron.Scheduler, sc chan bool) {
	//belum diisi, bingung
	time.Sleep(8 * time.Second)
	s.Clear()
	fmt.Println("All task removed")
	close(sc)
	//link https://stackoverflow.com/questions/34453894/cron-job-in-golang
}
