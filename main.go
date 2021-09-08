package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
	pb "github.com/howieyuen/tag-service/proto"
	"github.com/howieyuen/tag-service/server"
)

var (
	grpcPort string
	httpPort string
)

func init() {
	flag.StringVar(&grpcPort, "grpc_port", "8001", "grpc port")
	flag.StringVar(&httpPort, "http_port", "9001", "http port")
	flag.Parse()
}

func RunHttpServer(port string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})
	return http.ListenAndServe("127.0.0.1:"+port, serveMux)
}

func RunGrpcServer(port string) error {
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	return s.Serve(listener)
}

func main() {
	errCh := make(chan error)
	go func() {
		err := RunHttpServer(httpPort)
		if err != nil {
			errCh <- err
		}
	}()
	
	go func() {
		err := RunGrpcServer(grpcPort)
		if err != nil {
			errCh <- err
		}
	}()
	
	select {
	case err := <-errCh:
		log.Fatalf("Run server error: %v", err)
	}
}
