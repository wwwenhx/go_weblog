package main

import (
	"fmt"
	"gin_gomicro/app/user/grpc"
	"gin_gomicro/app/user/mq"
	"gin_gomicro/app/user/repository/db/dao"
	"gin_gomicro/config"
	"gin_gomicro/pkg/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	config.Init()
	err := mq.InitMq()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = mq.ConsumeUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	dao.Init()
	//etcd注册
	etcdClient, err := etcd.RegisterService("userService", config.UserServiceAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(etcdClient *clientv3.Client) {
		err := etcdClient.Close()
		if err != nil {

		}
	}(etcdClient)
	defer func() {
		err := etcd.UnRegisterService("userService")
		if err != nil {

		}
	}()
	grpc.GrpcInit()
}
