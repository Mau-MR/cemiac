package main

import "github.com/Mau-MR/cemiac/send"

func main(){
	emailClient := send.EmailAccount{
		Email: "example@gmail.com",
		Password: "pass",
		SmtpHost: "mail.gmail.com",
		SmtpPort: "587",
	}
	conn := emailClient.NewSMTPClient()
}
