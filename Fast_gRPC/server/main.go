package main

import (
	"context"
	"errors"
	"fmt"
	pb "gRPC_demo/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	// "google.golang.org/grpc/credentials"
	"net"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("server get message: " + req.GetRequestName())

	// 获取 token，可以用拦截器来实现，而不是每个请求都验证一次
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("未获取到Token")
	}
	var appId, appKey string
	if v, ok := md["app_id"]; ok {
		appId = v[0]
	}
	if v, ok := md["app_key"]; ok {
		appKey = v[0]
	}
	// 验证 token
	fmt.Println(fmt.Sprintf("app_id: %s, app_key: %s", appId, appKey))

	return &pb.HelloResponse{
		ResponseName: "hello " + req.RequestName,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
		return
	}
	// SSL 安全验证
	// abs_path := "./"
	// cred, err := credentials.NewServerTLSFromFile(abs_path+"test.pem", abs_path+"test.key")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// gServer := grpc.NewServer(grpc.Creds(cred))

	// 默认连接
	gServer := grpc.NewServer()

	pb.RegisterSayHelloServer(gServer, &server{})
	err = gServer.Serve(listen)
	if err != nil {
		fmt.Println(err)
		return
	}
}
