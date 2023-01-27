package common

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"os"
)

// @author by Hchier
// @Date 2023/1/25 10:13

func GetFile(dest string) *os.File {
	f, err := os.OpenFile(dest, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func Log(dest string, v ...interface{}) {
	file := GetFile(dest)
	defer file.Close()
	hlog.SetOutput(file)
	hlog.Error(v...)
}

// ErrLog 打印错误日志
func ErrLog(v ...interface{}) {
	Log(ErrLogPath, v)
}
