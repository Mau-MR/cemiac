//This file manages all the sending of the mails To another one
package send

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)
type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
	Server  *SmtpServer
	Auth *smtp.Auth
}

type SmtpServer struct {
	Host string
	Port string
}

func (s *SmtpServer) ServerName() string {
	return s.Host + ":" + s.Port
}

func (mail *Mail) BuildMessage() ([]byte,error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message :=""
	message += fmt.Sprintf("From: %s\r\n", mail.From)
	message += fmt.Sprintf("To: %s\r\n", mail.To)
	message += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	handler,err := NewTemplateHandler("rtest.html", &Message{
		Name: "Francisco",
		URL: "http://cemiacac.com",
		Body: mail.Body,
	})
	if err!=nil{
		return nil,err
	}
	return []byte(message+mime+handler.String()),nil
}

func (mail Mail) SendEmail() error {
	messageBody ,err:= mail.BuildMessage()
	if err !=nil{
		return err
	}
	//The auth of the account that is going to be used for sending the emails
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         mail.Server.Host,
	}

	conn, err := tls.Dial("tcp",mail.Server.ServerName(), tlsconfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, mail.Server.Host)
	if err != nil {
		return err
	}

	// step 1: Use Auth
	if err = client.Auth(*mail.Auth); err != nil {
		return err
	}

	// step 2: add all from and to
	if err = client.Mail(mail.From); err != nil {
		return err
	}
	if err = client.Rcpt(mail.To); err != nil {
		return err
	}
	// Data
	w, err := client.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(messageBody)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	client.Quit()

	log.Println("Mail sent successfully")
	return nil
}