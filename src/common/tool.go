package common

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//@author by Hchier
//@Date 2023/1/20 23:29

// GetRandStr 得到长度为len的随机字符串
func GetRandStr(len int8) string {
	result := make([]byte, len/2)
	rand.Read(result)
	return hex.EncodeToString(result)
}

// IsValidUser 验证身份是否有效
// 已登录：返回当前登录的用户的id
// 未登录：返回-1
func IsValidUser(ctx context.Context, c *app.RequestContext) {
	res, _ := Rdb.HGet(ctx, "tokens", string(c.FormValue("token"))).Result()
	userId, err := strconv.ParseInt(res, 10, 64)
	// 未登录
	if err != nil {
		c.Set("id", int64(-1))
		return
	}
	c.Set("id", userId)
}

// TimeTask 定时任务
func TimeTask(d time.Duration, task func()) {
	ticker := time.NewTicker(d)
	for range ticker.C {
		task()
	}
}

// Md5Encode md5加密
func Md5Encode(str string) string {
	srcCode := md5.Sum([]byte(str))
	// md5.Sum函数加密后返回的是字节数组，需要转换成16进制形式
	code := fmt.Sprintf("%x", srcCode)
	return code
}

func MakeDirs(path ...string) {
	for _, item := range path {
		err := os.MkdirAll(item, os.ModePerm) //创建多级目录\
		if err != nil {
			print(err.Error())
		}
	}
}
