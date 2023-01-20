package tool

import "os"

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
