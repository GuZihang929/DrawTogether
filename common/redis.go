package common

import (
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func InitRedis() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Config.Server.Redis,
		Password: "",
		DB:       0,
	})
	return
}
func CloseRedis() {
	_ = Rdb.Close()
}
