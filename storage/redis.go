package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
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
		Addr:     Cfg.Redis.Host + ":" + Cfg.Redis.Port,
		Password: Cfg.Redis.Password,
		DB:       Cfg.Redis.DB,
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
