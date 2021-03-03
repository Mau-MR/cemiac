package service

import (
	"context"
	"github.com/Mau-MR/cemiac/pb"
	"github.com/Mau-MR/cemiac/receive"
)

type ReceiveServer struct{
	pb.UnimplementedReceiveServiceServer
	Mail string
	Password string
}
func NewReceiveServer(user string, password string) *ReceiveServer  {
	return &ReceiveServer{
		Mail: user,
		Password: password,
	}
}
func(s*ReceiveServer)ReceiveMessages(c context.Context, req *pb.ReceiveMessageReq) (*pb.ReceiveMessageRes, error){
	imap :=receive.ImapAccount{
		Mail: s.Mail,
		Password: s.Password,
	}
	if err := imap.ClassifyMessages(); err !=nil{
		return nil,err
	}
	return &pb.ReceiveMessageRes{},nil
}
