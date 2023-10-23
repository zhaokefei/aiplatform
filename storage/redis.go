package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"

	cfg "github.com/zhaokefei/aiplatform/config"
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
		Addr:     cfg.Cfg.Redis.Host + ":" + cfg.Cfg.Redis.Port,
		Password: cfg.Cfg.Redis.Password,
		DB:       cfg.Cfg.Redis.DB,
		PoolSize: 100, // 连接池大小
	})
	fmt.Println(Client)
	
	ctx := context.Background()
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	log.Println("connect redis success")
	// 赋值给全局变量
	RedisClient = Client
}
