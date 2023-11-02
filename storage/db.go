package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	cfg "github.com/zhaokefei/aiplatform/config"
)

var DBClient *gorm.DB

func init() {
	NewMysqlClient()
	// 自动同步Table
	DBClient.AutoMigrate(&User{}, &Role{})
	for _, v := range Roles {
		if !IsSet(v) {
			NewRole(v)
		}
	}
}

func NewMysqlClient() {
	if DBClient != nil {
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Cfg.Mysql.Username,
		cfg.Cfg.Mysql.Password,
		cfg.Cfg.Mysql.Host,
		cfg.Cfg.Mysql.Port,
		cfg.Cfg.Mysql.Database,
	)

	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("connect mysql success")
	// 赋值给全局变量
	DBClient = client
}
