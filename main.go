package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
	pb "github.com/howieyuen/tag-service/proto"
	"github.com/howieyuen/tag-service/server"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "8001", "启动端口号")
	flag.Parse()
}

func RunHttpServer(port string) *http.Server {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})
	return &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
}

func RunTCPServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func RunGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	
	return s
}

func main() {
	listener, err := RunTCPServer(port)
	if err != nil {
		log.Fatalf("Run TCP Server err: %v", err)
	}
	
	mux := cmux.New(listener)
	grpcL := mux.MatchWithWriters(
		cmux.HTTP2MatchHeaderFieldPrefixSendSettings(
			"content-type",
			"application/grpc"))
	httpL := mux.Match(cmux.HTTP1Fast())
	
	grpcS := RunGrpcServer()
	go grpcS.Serve(grpcL)
	
	httpS := RunHttpServer(port)
	go httpS.Serve(httpL)
	
	err = mux.Serve()
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
}
