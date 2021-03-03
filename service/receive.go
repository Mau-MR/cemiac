package service

import (
	"context"
	"github.com/Mau-MR/cemiac/pb"
)

type ReceiveServer struct{
	pb.UnimplementedReceiveServiceServer
	User string
	Password string
}
func NewReceiveServer(user string, password string) *ReceiveServer  {
	return &ReceiveServer{
		User: user,
		Password: password,
	}
}
func(s*ReceiveServer)ReceiveMessages(c context.Context, req *pb.ReceiveMessageReq) (*pb.ReceiveMessageRes, error){

	return &pb.ReceiveMessageRes{},nil
}
