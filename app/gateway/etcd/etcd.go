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
	fmt.Println("enter")
	zapLogger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Error creating zap logger:", err)
		return
	}

	// 将 zap 日志记录器转换为 grpclog.LoggerV2 接口

	etcdClient, err := clientv3.New(clientv3.Config{Endpoints: []string{config.EtcdHost + ":" + config.EtcdPort},
		Logger: zapLogger})
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error connect to etcd:", err)
		return
	}
	fmt.Println("getServiceAddr")
	fmt.Println(serviceName)
	_, err = etcdClient.Put(context.Background(), "asd", "aqwe")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("mmm")
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
