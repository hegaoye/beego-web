package models

import (
	"context"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-redis/redis/v9"
	"time"
)

var redisCache *redis.Client
var ctx = context.Background()

func init() {
	server := beego.AppConfig.String("cache::server")
	password := beego.AppConfig.String("cache::password")
	db, _ := beego.AppConfig.Int("cache::db")
	redisCache = createClient(server, password, db)
}

func createClient(server string, password string, db int) *redis.Client {
	redisOptions := redis.Options{
		Addr:     server,
		Password: password,
		DB:       db,
	}

	client := redis.NewClient(&redisOptions)

	_, error := client.Ping(ctx).Result()
	if error != nil {
		logs.Error("连接redis失败", error)
	}
	return client
}

func SetStr(key, value string, time time.Duration) (err error) {
	err = redisCache.Set(ctx, key, value, time).Err()
	if err != nil {
		logs.Error("set key:", key, ",value:", value, err)
	}
	return
}

func GetKey(key string) (value string) {
	v, _ := redisCache.Get(ctx, key).Result()
	return v
}

func DelKey(key string) (err error) {
	err = redisCache.Del(ctx, key).Err()
	return
}
