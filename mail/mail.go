package mail

import (
	"crypto/tls"
	"github.com/emersion/go-imap/client"
	"log"
	"net/smtp"
)

type Account struct {
	Mail     string
	Password string
}
const SMTPHOST string= "smtp.gmail.com"
const SMTPPORT string ="465"
const IMAPHOST string = "imap.gmail.com"
const IMAPPORT string = "993"


func (a *Account) NewAccount(mail, password string) *Account  {
	return &Account{
		Mail:     mail,
		Password: password,
	}
}
//makes the  imap connection and returns the client
func (a * Account) imapConnection() (*client.Client, error){

	log.Println("Connecting to server...")
	// Connect to server
	c, err := client.DialTLS(IMAPHOST+":"+IMAPPORT, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         IMAPHOST,
	})
	if err != nil {
		return nil, err
	}

	// Login
	if err := c.Login(a.Mail, a.Password); err != nil {
		return nil, err
	}
	log.Println("Logged in with account ",a.Mail)
	return c, nil

}
//makes the smtp connection and  returns  the client
//TODO ADD THE ERR TO PRINT AS WELL
func (a * Account) smtpConnection() (*smtp.Client, error) {
	auth := smtp.PlainAuth("",a.Mail,a.Password,SMTPHOST)
	// TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName: SMTPHOST,
	}
	conn, err := tls.Dial("tcp", SMTPHOST+":"+SMTPPORT, tlsConfig)
	if err != nil {
		log.Fatal("Error opening the tls channel")
		return nil, err
	}
	smtpClient, err := smtp.NewClient(conn,SMTPHOST)
	if err != nil {
		log.Fatal("Error making the mail connection")
		return nil, err
	}
	// authenticate the account
	if err = smtpClient.Auth(auth); err != nil {
		log.Fatal("Error with the authentication of the account")
		return nil, err
	}
	return smtpClient, nil
}

func (a*Account) SendMail(mail *MailMessage) error {
	conn,err:= a.smtpConnection()
	//todo Check if its necessary to catch this  err
	defer conn.Quit()
	if err  !=nil {
		return err
	}
	if err:= mail.send(conn,SMTPHOST,SMTPPORT); err!=nil{
		return err
	}
	log.Println("MailMessage sent successfully to: ", mail.To)
	return nil
}
func (a Account) SendMails(mails *[]MailMessage) error  {
	conn,err:= a.smtpConnection()
	defer conn.Quit()
	if err !=nil{
		return err
	}
	//TODO: MAKE A RECOVERY SYSTEM IN CASE OF FAILURE TO SEND ONE MAIL
	for _,mail := range *mails{
		if err := mail.send(conn,SMTPHOST,SMTPPORT); err!=nil{
			return err
		}
	}
	return nil
}

func (a Account) ReceiveMails() (*[]ProcessedMail,error)  {
	conn,err := a.imapConnection()
	defer conn.LoggedOut()
	if err !=nil{
		return nil,err
	}
	unreadMails,err := receive(conn)
	//TODO TAKE IN  CONSIDERATION THE LEN OF THE UNREADMAILS BECAUSE IT CONTAINS REMAINING MESSAGES
	if err !=nil{
		return unreadMails,err
	}

	return unreadMails,err
}