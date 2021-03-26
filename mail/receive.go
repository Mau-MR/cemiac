package mail

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"io/ioutil"
	"log"
)

//The MailMessage that we are going to open the see the messages
//Important: The IMAP protocol has to be activated on the MailMessage to this properly works
//So first go an activate IMAP on the MailMessage
type ProcessedMail struct {
	From    string
	Subject string
	Body    string
}

func receive(client *client.Client) (*[]ProcessedMail, error) {
	done := make(chan error)
	seqSet, messages, err := getUnprocessedMessages(client, &done)
	if err != nil {
		return nil, err
	}

	processedMailChan, err := getBody(messages, &done)
	if err != nil {
		return nil, err
	}
	var unreadMails []ProcessedMail
	log.Println("New Messages: ")
	for ch := range processedMailChan {
		unreadMails = append(unreadMails,*ch)
	}
	if err := setMessagesAsProcessed(client, seqSet); err != nil {
		return &unreadMails,err
	}
	return &unreadMails, err
}

func getUnprocessedMessages(c *client.Client, done *chan error) (*imap.SeqSet, *chan *imap.Message, error) {
	// Select INBOX
	_, err := c.Select("INBOX", false)
	if err != nil {
		return nil, nil, err
	}
	//performing the search of messages without the processed flag
	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{"processed"}
	ids, err := c.Search(criteria)
	if err != nil {
		//Means thera are no new messages with these flag
		return nil, nil, err
	}
	if len(ids) > 0 {
		log.Println(ids)
		seqset := new(imap.SeqSet)
		seqset.AddNum(ids...)
		messages := make(chan *imap.Message, 10)
		var section imap.BodySectionName
		go func() {
			*done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}, messages)
		}()
		log.Println("Start of the mail processing")
		return seqset, &messages, nil
	}
	return nil, nil, fmt.Errorf("No new unprocessed messages ;)")
}
func setMessagesAsProcessed(c *client.Client, mRange *imap.SeqSet) error {
	item := imap.FormatFlagsOp(imap.AddFlags, true)
	flags := []interface{}{"processed"}
	err := c.Store(mRange, item, flags, nil)
	if err != nil {
		return err
	}
	log.Println("Messages has been marked as processed")
	return nil
}

//The pipeline to process  the  messages
func getBody(ch *chan *imap.Message, done *chan error) (chan *ProcessedMail, error) {

	outLiteral := make(chan imap.Literal)
	prs := func(ch *chan *imap.Message) {
		defer close(outLiteral)
		for c := range *ch {
			outLiteral <- c.GetBody(&imap.BodySectionName{})
		}
	}
	go prs(ch)

	outReaders := make(chan *mail.Reader)
	pts := func(ch *chan imap.Literal) {
		defer close(outReaders)
		for c := range *ch {
			mr, err := mail.CreateReader(c)
			if err != nil {
				*done <- err
				return
			}
			outReaders <- mr
		}
	}
	go pts(&outLiteral)

	processedMails := make(chan *ProcessedMail)
	pms := func(ch *chan *mail.Reader) {
		defer close(processedMails)
		for c := range *ch {
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
				subject, err := c.Header.Subject()
				if err != nil {
					log.Fatal("Unable to retrieve subject of message")
				}
				from, err := c.Header.AddressList("From")
				if err != nil {
					log.Fatal("Error getting the from")
				}
				processedMails <- &ProcessedMail{
					//TODO: Consider if its necessary to loop over the from not to get an error
					From:    from[0].Address,
					Subject: subject,
					Body:    string(b),
				}
			case *mail.AttachmentHeader:
				// This is an attachment
				filename, _ := h.Filename()
				log.Println("Got attachment: ", filename)
			}
		}
	}
	go pms(&outReaders)
	return processedMails, nil
}
