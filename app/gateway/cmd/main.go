package main

import (
	"gin_gomicro/app/gateway/routers"
	"gin_gomicro/config"
)

func main() {
	config.Init()
	err := routers.NewRouter().Run(":10010")
	if err != nil {
		return
	}
}
