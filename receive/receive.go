package receive

import (
	"crypto/tls"
	"fmt"
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
	//TODO: change to omit this step of listing the mailboxes
	done :=i.listMailBoxes(client)
	seqSet,err := i.getUnprocessedMessages(client,&done)
	if err !=nil{
		return err
	}
	i.setMessagesAsProcessed(client,seqSet)
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
func (i ImapAccount) getUnprocessedMessages(c *client.Client, done *chan error) (*imap.SeqSet,error){
	// Select INBOX
	_, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	//performing the search of messages without the processed flag
	criteria:= imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{"processed"}
	ids,err := c.Search(criteria)
	if err !=nil{
		//Means thera are no new messages with these flag
		log.Fatal(err)
	}
	if len(ids)>0{
		seqset := new(imap.SeqSet)
		seqset.AddNum(ids...)
		messages := make(chan *imap.Message, 10)
		go func() {
			*done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
		}()

		log.Println("Unprocessed messages:")
		for msg := range messages {
			log.Println("* " + msg.Envelope.Subject)
		}
		if err := <-*done; err != nil {
			return nil,err
		}
		return seqset,nil
	}
	log.Println("Done!")
	return nil, fmt.Errorf("No new unprocessed messages ;)")
}
func (i ImapAccount) setMessagesAsProcessed(c * client.Client, mRange *imap.SeqSet)  {
	item := imap.FormatFlagsOp(imap.AddFlags, true)
	flags := []interface{}{"processed"}
	err := c.Store(mRange, item, flags, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Messages has been marked as processed")
}