package mapper

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"
	"tiktok/src/tool"
)

//@author by Hchier
//@Date 2023/1/20 20:15

// InsertUser
// 成功插入则返回用户id
// 失败则返回-1
func InsertUser(username, password string) (int32, string) {
	res, err := Db.Exec("insert into user (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		file := tool.GetFile(tool.ErrLogDest)
		defer file.Close()
		hlog.SetOutput(file)
		hlog.Error("插入失败：", err.Error())
		return -1, err.Error()
	}
	id, _ := res.LastInsertId()
	return int32(id), "注册成功"
}

func ExistUser(username string) bool {
	var id int8
	err := Db.Get(&id, "select id from user where username = ?", username)
	if err != nil {
		return false
	}
	return true
}
