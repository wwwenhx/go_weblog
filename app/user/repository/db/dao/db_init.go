package dao

import (
	"context"
	"fmt"
	"gin_gomicro/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 定义数据库模型
var DB *gorm.DB

func Init() {
	host := config.DbHost
	port := config.DbPort
	user := config.DbUser
	password := config.DbPassword
	database := config.DbName
	charset := config.Charset
	fmt.Println(user, password, host, port, database, charset)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", user, password, host, port, database, charset)
	err := Database(dsn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("数据库连接成功")
}

func Database(connString string) error {
	var ormLogger logger.Interface = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.New(mysql.Config{
		//配置mysql
		DSN:                       connString, //dsn data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       //禁用datatime精度，兼容mysql5.6
		DontSupportRenameIndex:    true,       //不支持重命名索引
		DontSupportRenameColumn:   true,       //不支持重命名列
		SkipInitializeWithVersion: false,      // 根据版本自动配置
		//配置gorm
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{ //命名策略
			SingularTable: true, //不使用的话 user-->users
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	DB = db
	return err
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}
