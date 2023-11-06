package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	cfg "github.com/zhaokefei/aiplatform/config"
)

var Cfg cfg.Config
var DBClient *gorm.DB

func init() {
	log.Println("storage init")
	cfg, err := cfg.LoadConfig()
	if err != nil {
		panic(err)
	}
	// 赋值
	Cfg = cfg
	NewDBClient()
	// 自动同步Table
	DBClient.AutoMigrate(&User{}, &Role{}, &App{}, &AppCategory{})
	for _, v := range Roles {
		if !IsSet(v) {
			NewRole(v)
		}
	}
	// 创建ADMIN USER
	RegisterAdmin()
	// 创建初始化应用
	RegisterApps()
}

func NewDBClient() {
	if DBClient != nil {
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Cfg.Mysql.Username,
		Cfg.Mysql.Password,
		Cfg.Mysql.Host,
		Cfg.Mysql.Port,
		Cfg.Mysql.Database,
	)

	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("connect mysql success")
	// 赋值给全局变量
	DBClient = client
}
