package main

import (
	"fmt"
	"gin_gomicro/app/user/etcd"
	"gin_gomicro/app/user/grpc"
	"gin_gomicro/app/user/repository/db/dao"
	"gin_gomicro/config"
)

func main() {
	config.Init()
	dao.Init()
	//etcd注册
	etcdClient, err := etcd.RegisterService("userService", config.UserServiceAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer etcdClient.Close()
	defer etcd.UnRegisterService("userService")
	grpc.GrpcInit()
}
