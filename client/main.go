package main

import (
	"context"
	"log"
	
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/howieyuen/tag-service/internal/middleware"
	pb "github.com/howieyuen/tag-service/proto"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	clientConn, err := GetClientConn(ctx, ":8001", DailOpts())
	/*
		clientConn, err := GetClientConn(ctx, ":8001", []grpc.DialOption{
			grpc.WithUnaryInterceptor(
				grpc_middleware.ChainUnaryClient(
					middleware.UnaryContextTimeout(),
				),
			),
		})
	*/
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer clientConn.Close()
	
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, err := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})
	if err != nil {
		log.Fatalf("tagServiceClient.GetTagList err: %v", err)
	}
	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}

func DailOpts() []grpc.DialOption {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			middleware.UnaryContextTimeout(),
		),
	))
	// opts = append(opts, grpc.WithStreamInterceptor(
	// 	grpc_middleware.ChainStreamClient(
	// 		middleware.StreamContextTimeout(),
	// 	),
	// ))
	return opts
}
