package grpc

import (
	"fmt"
	"gin_gomicro/app/msg/service"
	"gin_gomicro/config"
	"gin_gomicro/idl/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type MsgRpc struct{}

func GrpcInit() {
	//创建grpc服务器
	grpcServer := grpc.NewServer()
	//注册userService结构体
	pb.RegisterMsgServer(grpcServer, service.GetMsgSrv())
	//启动grpc服务器
	lis, err := net.Listen("tcp", config.MsgServiceAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("grpc server listening on :8090")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("grpc server listening on :8090")

}
