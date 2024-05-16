package etcd

import (
	"context"
	"fmt"
	"gin_gomicro/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func RegisterService(name, addr string) (etcdClient *clientv3.Client, err error) {
	etcdClient, err = clientv3.New(clientv3.Config{Endpoints: []string{config.EtcdHost + ":" + config.EtcdPort}})
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = etcdClient.Put(context.Background(), name, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := etcdClient.Get(context.Background(), "userService")
	fmt.Println(resp)
	fmt.Println("register success")
	return
}

func UnRegisterService(name string) (err error) {
	etcdClient, err := clientv3.New(clientv3.Config{Endpoints: []string{config.EtcdHost + ":" + config.EtcdPort}})
	if err != nil {
		fmt.Println(err)
	}

	_, err = etcdClient.Delete(context.Background(), name, clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("unregister success")
	return
}
