package main

import (
	"flag"
	"fmt"
	"github.com/Mau-MR/cemiac/pb"
	"github.com/Mau-MR/cemiac/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	password:= "someappPassword"
	sender := "example@gmail.com"

	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)


	grpcServer := grpc.NewServer()
	sendServer := service.NewSendServer(sender,password)
	receiveServer := service.NewReceiveServer(sender,password)

	pb.RegisterSendServiceServer(grpcServer,sendServer)
	pb.RegisterReceiveServiceServer(grpcServer,receiveServer)


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
