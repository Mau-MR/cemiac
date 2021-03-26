package main

import (
	"github.com/Mau-MR/cemiac/mail"
	"log"
)

func main() {
	account := mail.Account{
		Mail: "example@gmail.com",
		Password: "exmaple",
	}
	unreadMails ,err := account.ReceiveMails()
	if err !=nil{
		log.Fatal(err)
	}
	for _,m:= range *unreadMails{
		log.Println(m)
	}
}
