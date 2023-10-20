package storage

import (
	"log"

	"github.com/redis/go-redis/v9"

	conf "github.com/zhaokefei/aiplatform/config"
)

var RedisClient *redis.Client

func init() {
	NewRedisClient()
}

func NewRedisClient() {
	if RedisClient != nil {
		return
	}
	Client := redis.NewClient(&redis.Options{
		Addr:     conf.Cfg.Redis.Host + ":" + conf.Cfg.Redis.Port,
		Password: conf.Cfg.Redis.Password,
		DB:       conf.Cfg.Redis.DB,
	})
	log.Println("connect redis success")
	// 赋值给全局变量
	RedisClient = Client
}
