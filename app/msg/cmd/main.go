package main

import (
	"fmt"
	"gin_gomicro/app/msg/grpc"
	"gin_gomicro/app/msg/msgMq"
	"gin_gomicro/app/msg/repository/db/dao"
	"gin_gomicro/config"
	"gin_gomicro/pkg/etcd"
	"gin_gomicro/pkg/mq"
)

func main() {
	config.Init()
	dao.Init()
	err := mq.InitMq("MsgChannel")
	if err != nil {
		fmt.Println(err)
	}
	msgMq.Consume("MsgChannel")
	//注册etcd
	etcdClient, err := etcd.RegisterService("msgService", config.MsgServiceAddress)
	defer etcd.UnRegisterService("msgService")
	defer etcdClient.Close()
	grpc.GrpcInit()
}
