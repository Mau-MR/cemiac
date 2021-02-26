package send

//This file manage the account that sends the email
import (
	"crypto/tls"
	"log"
	"net/smtp"
)

type EmailAccount struct {
	//The email which is going to send de mail
	Email string;
	//The password of that mail
	Password string;
	//The port of the smtp server normaly is 587
	SmtpPort string;
	//The dns of the mail server example: smtp.gmail.com
	SmtpHost string;
}

//return a new smtp client with tls
func (e EmailAccount) NewSMTPClient() *smtp.Client {
	auth := smtp.PlainAuth("", e.Email, e.Password, e.SmtpHost)

	servername := e.SmtpHost+":"+e.SmtpPort

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName: e.SmtpHost,
	}
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}
	c, err := smtp.NewClient(conn, e.SmtpHost)
	if err != nil {
		log.Panic(err)
	}
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}
	return c
}
