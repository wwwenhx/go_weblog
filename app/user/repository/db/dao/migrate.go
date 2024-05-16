package dao

import "gin_gomicro/app/user/repository/db/model"

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{})
	if err != nil {
		return
	}
}
