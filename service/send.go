package service

import (
	"context"
	"github.com/Mau-MR/cemiac/pb"
)
type SendServer struct {
	pb.UnimplementedSendServiceServer
	User string
	Password string
}
//TODO: add the firestore connection here
func NewSendServer(user string,password string)*SendServer{
	return &SendServer{
		User: user,
		Password: password,
	}
}
func (s * SendServer)SendMail( c context.Context, r *pb.SendReq) (*pb.SendResp, error){
	return &pb.SendResp{},nil
}

func (s*SendServer) SendMails(pb.SendService_SendMailsServer) error{
	return nil
}
