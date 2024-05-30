package main

import (
	"context"
	"fmt"
	pb "gRPC_demo/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// "google.golang.org/grpc/credentials"
)

// type PerRPCCredentials interface {
// 	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
// 	RequireTransportSecurity() bool
// }

// Token 实现
type ClientTokenAuth struct{}

func (cta *ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"app_id":  "test",
		"app_key": "123456",
	}, nil
}

func (cta *ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	// SSL
	// abs_path := "D:/Go_projs/gRPC_demo/keys/"
	// cred, err := credentials.NewClientTLSFromFile(abs_path+"test.pem", "*")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// 安全连接
	// conn, err := grpc.NewClient("127.0.0.1:9090", grpc.WithTransportCredentials(cred))

	// 不安全连接
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth))) // 请求时 带上 token

	conn, err := grpc.NewClient("127.0.0.1:9090", opts...)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		conn.Close()
	}()

	client := pb.NewSayHelloClient(conn)
	helloResp := &pb.HelloRequest{RequestName: "test cag"}
	fmt.Println("client send message: " + fmt.Sprintf("%v", helloResp))

	rep, err := client.SayHello(context.Background(), helloResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rep.GetResponseName())
}
