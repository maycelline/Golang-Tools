package gomail

import (
	"GoTools/model"
	"bytes"
	"fmt"
	"text/template"

	gm "gopkg.in/gomail.v2"
)

func SendEmail(user model.User) {
	mail := gm.NewMessage()

	template := "gomail/message.html"

	result, _ := parseTemplate(template, user)

	mail.SetHeader("From", "perpushb@gmail.com")
	mail.SetHeader("To", user.Email)
	mail.SetHeader("Subject", "Testing Send Email")
	mail.SetBody("text/html", result)

	sender := gm.NewDialer("smtp.gmail.com", 587, "perpushb@gmail.com", "PerpusHBH1tZ")

	if err := sender.DialAndSend(mail); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email Delivered to ", user.Email)
	}

}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	// mengubah text html ke dalam bentuk byte
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	buff := new(bytes.Buffer)

	//render struct ke dalam file html (td ada var name di dalam htmlnya)
	if err = t.Execute(buff, data); err != nil {
		fmt.Println(err)
		return "", err
	}

	return buff.String(), nil
}
