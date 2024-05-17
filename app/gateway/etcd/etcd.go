package etcd

import (
	"context"
	"errors"
	"fmt"
	"gin_gomicro/config"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

func GetServiceAddr(serviceName string) (addr string) {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Error creating zap logger:", err)
		return
	}

	// 将 zap 日志记录器转换为 grpclog.LoggerV2 接口
	fmt.Println([]string{config.EtcdHost + ":" + config.EtcdPort})
	etcdClient, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"},
		Logger: zapLogger})
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error connect to etcd:", err)
		return
	}
	resp, err := etcdClient.Get(context.Background(), serviceName)
	if err != nil {
		fmt.Println("getServiceAddr err")
		fmt.Println(err)
		return
	}
	fmt.Println("getServiceAddr", resp.Kvs)
	if len(resp.Kvs) == 0 {
		addr = ""
		err = errors.New("service not exist")
		fmt.Println(err)
		return
	}
	fmt.Println("success")
	return string(resp.Kvs[0].Value)
}
