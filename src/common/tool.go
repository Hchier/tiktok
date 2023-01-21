package common

import (
	"encoding/hex"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"os"
)

//@author by Hchier
//@Date 2023/1/20 23:29

const ErrLogDest = "logs/err.log"

func GetFile(dest string) *os.File {
	f, err := os.OpenFile(dest, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	return f
}

var Rdb = GetRdb()

func GetRdb() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetRandStr() string {
	result := make([]byte, 16/2)
	println(len(result))
	rand.Read(result)
	return hex.EncodeToString(result)
}

func Log(dest string, v ...interface{}) {
	file := GetFile(dest)
	defer file.Close()
	hlog.SetOutput(file)
	hlog.Error(v...)
}
