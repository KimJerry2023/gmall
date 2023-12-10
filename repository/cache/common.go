package cache

import (
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"gmall/conf"
	"strconv"
)

// RedisClient redis缓存客户端单例
var RedisClient *redis.Client

// InitCache 在中间件中初始化redis链接，防止循环导包
func InitCache() {
	Redis()
}

func Redis() {
	db, _ := strconv.ParseUint(conf.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}
