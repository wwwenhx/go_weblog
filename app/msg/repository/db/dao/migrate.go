package dao

import "gin_gomicro/app/msg/repository/db/model"

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.Msg{})
	if err != nil {
		return
	}
}
