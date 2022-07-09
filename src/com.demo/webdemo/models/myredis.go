package models

import (
	beego "github.com/beego/beego/adapter"
	"github.com/beego/beego/core/logs"
	"gopkg.in/redis.v5"
	"time"
)

var redisCache *redis.Client

func init() {
	server := beego.AppConfig.String("cache::server")
	password := beego.AppConfig.String("cache::password")
	db, _ := beego.AppConfig.Int("cache::db")
	redisCache = createClient(server, password, db)
}

func createClient(server string, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{Addr: server, Password: password, DB: db})
	_, error := client.Ping().Result()
	if error != nil {
		logs.Error("连接redis失败", error)
	}
	return client
}

func SetStr(key, value string, time time.Duration) (err error) {
	err = redisCache.Set(key, value, time).Err()
	if err != nil {
		logs.Error("set key:", key, ",value:", value, err)
	}
	return
}

func GetKey(key string) (value string) {
	v, _ := redisCache.Get(key).Result()
	return v
}

func DelKey(key string) (err error) {
	err = redisCache.Del(key).Err()
	return
}
