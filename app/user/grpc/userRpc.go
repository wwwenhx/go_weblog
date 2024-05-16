package grpc

import (
	"fmt"
	"gin_gomicro/app/user/service"
	"gin_gomicro/config"
	"gin_gomicro/idl/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserRpc struct {
}

func GrpcInit() {
	//创建grpc服务器
	grpcServer := grpc.NewServer()
	//注册userService结构体
	pb.RegisterUserServer(grpcServer, service.GetUserSrv())
	//启动grpc服务器
	lis, err := net.Listen("tcp", config.UserServiceAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("grpc server listening on :8088")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("grpc server listening on :8088")

}

func NewGrpcClient() pb.UserClient {
	conn, err := grpc.Dial("localhost:8080")
	if err != nil {
		return nil
	}
	defer conn.Close()

	return pb.NewUserClient(conn)
}
