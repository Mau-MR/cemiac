package main

import (
	"github.com/Mau-MR/cemiac/receive"
)

func main() {
	imapAccount :=receive.ImapAccount{
		Mail: "example@gmail.com",
		Password: "someapplicationpassword",
	}
	imapAccount.ClassifyMessages()
}
