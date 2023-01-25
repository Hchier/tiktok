package common

import (
	"context"
	"github.com/go-redis/redis/v8"
)

//@author by Hchier
//@Date 2023/1/25 10:15

var Rdb *redis.Client

func init() {
	Rdb = GetRdb()
}

func GetRdb() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPassword, // no password set
		DB:       0,             // use default DB
	})
}

func DelKeys(keys ...string) {
	Rdb.Del(context.Background(), keys...)
}
