package receive

import (
	"crypto/tls"
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"io/ioutil"
	"log"
)

//The account that we are going to open the see the messages
//Important: The IMAP protocol has to be activated on the account to this properly works
//So first go an activate IMAP on the account
type ImapAccount struct{
	Mail string
	Password string
}
type ProcessedMail struct{
	From string
	Subject string
	Body string
}

func (i ImapAccount) ClassifyMessages() error  {
	c, err:=i.authenticate()
	//Closing the sessión
	if err !=nil{
		return err
	}
	defer c.Logout()
	done := make(chan error)
	seqSet,messages,err := i.getUnprocessedMessages(c,&done)
	if err !=nil{
		return err
	}
	processedMailChan :=i.getBody(messages)
	for ch :=range processedMailChan{
		log.Println("This is the new message")
		log.Println(ch)
	}
	i.setMessagesAsProcessed(c,seqSet)
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
func (i ImapAccount) listMailBoxes(c *client.Client,done *chan error) *chan error {
	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	go func () {
		*done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-*done; err != nil {
		log.Fatal(err)
	}
	return done
}
func (i ImapAccount) getUnprocessedMessages(c *client.Client, done *chan error) (*imap.SeqSet,*chan*imap.Message,error){
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
		var section imap.BodySectionName
		go func() {
			*done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}, messages)
		}()
		if err := <-*done; err != nil {
			return nil,nil,err
		}
		return seqset,&messages,nil
	}
	return nil, nil, fmt.Errorf("No new unprocessed messages ;)")
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

func (i ImapAccount) getBody(ch *chan *imap.Message) chan *ProcessedMail {

	outLiteral := make(chan imap.Literal)
	prs:= func(ch* chan*imap.Message){
		for c := range *ch{
			outLiteral<- c.GetBody(&imap.BodySectionName{})
		}
		close(outLiteral)
	}
	go prs(ch)

	outReaders := make(chan *mail.Reader)
	pts := func(ch*chan imap.Literal) {
		for c := range *ch{
			mr, err := mail.CreateReader(c)
			if err != nil {
				//TODO: ADD the done or a retunr
				log.Fatal(err)
			}
			outReaders <- mr
		}
		close(outReaders)
	}
	go pts(&outLiteral)


	processedMails := make(chan *ProcessedMail)
	pms := func(ch * chan*mail.Reader) {
		for c := range *ch{
			//In case of the need of mapping the whole body
			//For what we want its okay on this way
			//for {
			mr := c
			//Moves trow all the parts of the body
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			switch h := p.Header.(type) {
			case *mail.InlineHeader:
				// This is the message's text (can be plain-text or HTML)
				b, _ := ioutil.ReadAll(p.Body)
				subject ,err:= c.Header.Subject()
				if err !=nil{
					log.Fatal("Error getting the header")
				}
				from ,err:= c.Header.AddressList("From")
				if err !=nil{
					log.Fatal("Error getting the from")
					continue
				}
				processedMails <- &ProcessedMail{
					//TODO: Consider if its necessary to loop over the from not to get an error
					From: from[0].Address,
					Subject: subject,
					Body: string(b),
				}
			case *mail.AttachmentHeader:
				// This is an attachment
				filename, _ := h.Filename()
				log.Println("Got attachment: %v", filename)
			}
			}
		//}
		close(processedMails)
	}
	go pms(&outReaders)
	return processedMails
}