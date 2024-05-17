package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	Db                 string
	DbHost             string
	DbPort             string
	DbUser             string
	DbPassword         string
	DbName             string
	Charset            string
	EtcdHost           string
	EtcdPort           string
	UserServiceAddress string
	TaskServiceAddress string
	GatewayAddress     string
	MqAddress          string
)

func Init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件错误", err)
	}
	LoadMySqlData(file)
	LoadEtcdData(file)
	LoadServiceAddr(file)
	LoadMqAddr(file)
}

func LoadMySqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
	Charset = file.Section("mysql").Key("Charset").String()
}

func LoadEtcdData(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()
}

func LoadServiceAddr(file *ini.File) {
	UserServiceAddress = file.Section("server").Key("UserServiceAddress").String()
	TaskServiceAddress = file.Section("server").Key("TaskServiceAddress").String()
	GatewayAddress = file.Section("server").Key("GatewayAddress").String()
}

func LoadMqAddr(file *ini.File) {
	MqAddress = file.Section("mq").Key("MqAddress").String()
}
