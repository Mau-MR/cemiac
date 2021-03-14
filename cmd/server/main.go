package main

import (
	"flag"
	"fmt"
	"github.com/Mau-MR/cemiac/pb"
	"github.com/Mau-MR/cemiac/send"
	"github.com/Mau-MR/cemiac/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	password:= "some pass"
	sender := "example@gmail.com"
	smtpServer := &send.SmtpServer{
		Host: "smtp.gmail.com",
		//secure port
		Port: "465",
	}

	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)


	grpcServer := grpc.NewServer()
	sendServer := service.NewSendServer(sender,password,smtpServer)
	receiveServer := service.NewReceiveServer(sender,password)

	pb.RegisterSendServiceServer(grpcServer,sendServer)
	pb.RegisterReceiveServiceServer(grpcServer,receiveServer)

	//For testing
	reflection.Register(grpcServer)


	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start sever: ", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
