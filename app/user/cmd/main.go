package main

import (
	"gin_gomicro/app/user/grpc"
	"gin_gomicro/app/user/repository/db/dao"
	"gin_gomicro/config"
)

func main() {
	config.Init()
	dao.Init()
	//etcd注册
	grpc.GrpcInit()

}
