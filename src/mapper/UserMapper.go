package mapper

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"
	"tiktok/src/common"
)

//@author by Hchier
//@Date 2023/1/20 20:15

// InsertUser
// 成功插入则返回用户id
// 失败则返回-1
func InsertUser(username, password string) (int64, string) {
	res, err := Db.Exec("insert into user (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		file := common.GetFile(common.ErrLogDest)
		defer file.Close()
		hlog.SetOutput(file)
		hlog.Error("插入失败：", err.Error())
		return -1, err.Error()
	}
	id, _ := res.LastInsertId()
	return id, "注册成功"
}

func ExistUser(username string) bool {
	var id int8
	err := Db.Get(&id, "select id from user where username = ?", username)
	if err != nil {
		return false
	}
	return true
}

func GetIdByUsernameAndPassword(username, password string) int64 {
	var id int64
	err := Db.Get(&id, "select id from user where username = ? and password = ?", username, password)
	if err != nil {
		return -1
	}
	return id
}

func SelectUserById(id int64) User {
	var user User
	err := Db.Get(&user, "select * from user where id = ?", id)
	if err != nil {
		common.Log(common.ErrLogDest, "查找用户信息失败：", err.Error())
	}
	return user
}
