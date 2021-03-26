//This file manages all the sending of the mails To another one
package mail

import (
	"fmt"
	"net/smtp"
)

type MailMessage struct {
	From    string
	To      string
	Subject string
	Body    string
}

func (mail *MailMessage) buildMessage() ([]byte, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.From)
	message += fmt.Sprintf("To: %s\r\n", mail.To)
	message += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	handler, err := NewTemplateHandler("rtest.html", &Message{
		Name: "Francisco",
		URL:  "http://cemiacac.com",
		Body: mail.Body,
	})
	if err != nil {
		return nil, err
	}
	return []byte(message + mime + handler.String()), nil
}

//TODO ADD ALL THE PRINT  FOR THE ERRORJJ
//Recevies a smtp auth and the host and  port and sends the mail
func (mail MailMessage) send(client *smtp.Client, host, port string) error {
	messageBody, err := mail.buildMessage()
	if err != nil {
		return err
	}

	if err = client.Mail(mail.From); err != nil {
		return err
	}
	if err = client.Rcpt(mail.To); err != nil {
		return err
	}
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

	return nil
}
