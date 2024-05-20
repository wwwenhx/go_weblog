package main

import (
	"gin_gomicro/app/msg/repository/db/dao"
	"gin_gomicro/config"
)

func main() {
	config.Init()
	dao.Init()
}
