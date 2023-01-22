package mapper

import (
	"database/sql"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
	err := Db.Get(&id, "select id from user where username = ? and deleted = 0", username)
	if err != nil {
		return false
	}
	return true
}

func GetIdByUsernameAndPassword(username, password string) int64 {
	var id int64
	err := Db.Get(&id, "select id from user where username = ? and password = ? and deleted = 0", username, password)
	if err != nil {
		return -1
	}
	return id
}

func SelectUserById(id int64) User {
	var user User
	err := Db.Get(&user, "select * from user where id = ? and deleted = 0", id)
	if err != nil {
		common.ErrLog("查找用户信息失败：", err.Error())
	}
	return user
}

// UpdateUserTotalFavorited 更新作者获赞数。opType -> 1：加1；2：减1
// 成功返回true
func UpdateUserTotalFavorited(opType int8, authorId int64, tx *sqlx.Tx) bool {
	var res sql.Result
	var err error
	if opType == 1 {
		res, err = tx.Exec("update user set total_favorited = total_favorited + 1 where id = ?", authorId)
	} else {
		res, err = tx.Exec("update user set total_favorited = total_favorited - 1 where id = ?", authorId)
	}
	if err != nil {
		common.ErrLog("更新用户点赞数失败", err.Error())
		return false
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		common.ErrLog("更新用户点赞数时RowsAffected为0", err.Error())
		return false
	}
	return true
}

// UpdateUserFavoriteCount 更新用户点赞数。opType -> 1：加1；2：减1
// 成功返回true
func UpdateUserFavoriteCount(opType int8, userId int64, tx *sqlx.Tx) bool {
	var res sql.Result
	var err error
	if opType == 1 {
		res, err = tx.Exec("update user set favorite_count = favorite_count + 1 where id = ?", userId)
	} else {
		res, err = tx.Exec("update user set favorite_count = favorite_count - 1 where id = ?", userId)
	}
	if err != nil {
		common.ErrLog("更新用户点赞数失败", err.Error())
		return false
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		common.ErrLog("更新用户点赞数时RowsAffected为0", err.Error())
		return false
	}
	return true
}
