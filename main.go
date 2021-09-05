package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/howieyuen/tag-service/proto"
	"github.com/howieyuen/tag-service/server"
)

func main() {
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTagServiceServer(s, server.NewTagServer())

	listener, err := net.Listen("tcp", ":8001")

	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}
