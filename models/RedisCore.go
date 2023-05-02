package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var redisCoretxt = context.Background()
var (
	RedisDb *redis.Client
)

func init() {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisDb.Ping(redisCoretxt).Result() //测试是否redis数据库能被连接
	if err != nil {
		println(err)
	}
}
