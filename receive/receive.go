package receive

import (
	"crypto/tls"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"log"
)

//The account that we are going to open the see the messages
//Important: The IMAP protocol has to be activated on the account to this properly works
//So first go an activate IMAP on the account
type ImapAccount struct{
	Mail string
	Password string
}

func (i ImapAccount) ClassifyMessages() error  {
	client,err:=i.authenticate()
	//Closing the sessión
	defer client.Logout()
	if err !=nil{
		return err
	}
	done :=i.listMailBoxes(client)
	i.getMessages(client,&done)
	return nil
}

func (i ImapAccount) authenticate() (*client.Client,error) {
	log.Println("Connecting to server...")
	// Connect to server
	c, err := client.DialTLS("imap.gmail.com:993", &tls.Config{
		InsecureSkipVerify: true,
		ServerName: "imap.gmail.com" ,
	})
	if err != nil {
		return nil,err
	}
	log.Println("Connected to the IMAP account!")

	// Login
	if err := c.Login(i.Mail,i.Password); err != nil {
		return nil,err
	}
	log.Println("Logged in")
	return c,nil
}
func (i ImapAccount) listMailBoxes(c *client.Client) chan error {
	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func () {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
	return done
}
func (i ImapAccount) getMessages(c* client.Client,done *chan error)  {
	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// We're using unsigned integers here, only subtract if the result is > 0
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	go func() {
		*done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	log.Println("Last 4 messages:")
	for msg := range messages {
		log.Println("* " + msg.Envelope.Subject)
	}

	if err := <-*done; err != nil {
		log.Fatal(err)
	}
	log.Println("Done!")
}