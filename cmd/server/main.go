package main

import (
	"github.com/Mau-MR/cemiac/send"
	"net/smtp"
)

func main() {

	auth:=smtp.PlainAuth("", "example@gmail.com", "someapplicationpassword", "smtp.gmail.com")
	mail :=send.Mail{
		From: "example@gmail.com",
		To: "example@gmail.com",
		Subject: "Fachelito eh",
		Body: "This is some  amazing information",
		Server: &send.SmtpServer{
			Host: "smtp.gmail.com",
			Port: "465",
		},
		Auth : &auth,
	}
	mail.SendEmail()


}
