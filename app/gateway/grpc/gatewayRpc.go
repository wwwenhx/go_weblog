package grpc

import (
	"fmt"
	"gin_gomicro/idl/pb"
	"google.golang.org/grpc"
	"log"
)

func GetRpc(addr string) (grpcClient pb.UserClient) {
	conn, err := grpc.NewClient(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("dialer err:", err)
		return
	}
	grpcClient = pb.NewUserClient(conn)
	fmt.Println(conn.Target())
	fmt.Println("grpc client created", grpcClient)
	return
}
