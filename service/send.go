package service

import (
	"context"
	"github.com/Mau-MR/cemiac/pb"
	"github.com/Mau-MR/cemiac/send"
	"net/smtp"
)
type SendServer struct {
	pb.UnimplementedSendServiceServer
	Auth *smtp.Auth
	SMTPServer *send.SmtpServer
	Account string
}
//TODO: add the firestore connection here
func NewSendServer(sender string, password string, smtpServer * send.SmtpServer)*SendServer{
	auth := smtp.PlainAuth("", sender, password, smtpServer.Host)
	return &SendServer{
		Auth: &auth,
		SMTPServer: smtpServer,
		Account: sender,
	}
}
func (s * SendServer)SendMail( c context.Context, r *pb.SendReq) (*pb.SendResp, error){
	//TODO: MAKE FETCH to firestore to parse the templates and append them to the body and subject
	mail:= send.Mail{
		From: s.Account,
		To: r.GetMail(),
		Subject: "Hey que tal!",
		Body: "This is some important information",
		Server: s.SMTPServer,
		Auth: s.Auth,
	}
	if err := mail.SendEmail(); err !=nil{
		return nil,err
	}

	return &pb.SendResp{},nil
}

func (s*SendServer) SendMails(pb.SendService_SendMailsServer) error{
	return nil
}
