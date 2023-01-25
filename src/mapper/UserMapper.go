package mapper

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"tiktok/src/common"
)

//@author by Hchier
//@Date 2023/1/20 20:15

// InsertUser
// 成功插入则返回用户id
// 失败则返回-1
func InsertUser(username, password, avatar, backgroundImage, signature string) (int64, string) {
	res, err := Db.Exec("insert into user (username, password, avatar, background_image, signature) VALUES (?, ?, ?, ?, ?)",
		username, password, avatar, backgroundImage, signature)
	if err != nil {
		common.ErrLog("插入失败：", err.Error())
		return -1, err.Error()
	}
	id, _ := res.LastInsertId()
	return id, "注册成功"
}

func ExistUserByUsername(username string) bool {
	var id int64
	err := Db.Get(&id, "select id from user where username = ? and deleted = 0", username)
	if err != nil {
		return false
	}
	return true
}

func ExistUserById(targetUserId int64) bool {
	var id int64
	err := Db.Get(&id, "select id from user where id = ? and deleted = 0", targetUserId)
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

// SelectUserById 根据id查找用户信息
// 若用户不存在，则返回的user的id为-1
func SelectUserById(id int64) User {
	var user User
	err := Db.Get(&user, "select * from user where id = ? and deleted = 0", id)
	if err != nil {
		common.ErrLog("查找用户信息失败：", err.Error())
		user.Id = -1
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

// UpdateUserFollowerCount 更新用户粉丝数。opType -> 1：加1；2：减1
// 成功返回true
func UpdateUserFollowerCount(opType int8, userId int64, tx *sqlx.Tx) bool {
	var res sql.Result
	var err error
	if opType == 1 {
		res, err = tx.Exec("update user set follower_count = follower_count + 1 where id = ?", userId)
	} else {
		res, err = tx.Exec("update user set follower_count = follower_count - 1 where id = ?", userId)
	}
	if err != nil {
		common.ErrLog("更新用户粉丝数失败", err.Error())
		return false
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		common.ErrLog("更新用户粉丝数时RowsAffected为0", err.Error())
		return false
	}
	return true
}

// UpdateUserFollowCount 更新用户偶像数。opType -> 1：加1；2：减1
// 成功返回true
func UpdateUserFollowCount(opType int8, userId int64, tx *sqlx.Tx) bool {
	var res sql.Result
	var err error
	if opType == 1 {
		res, err = tx.Exec("update user set follow_count = follow_count + 1 where id = ?", userId)
	} else {
		res, err = tx.Exec("update user set follow_count = follow_count - 1 where id = ?", userId)
	}
	if err != nil {
		common.ErrLog("更新用户粉丝数失败", err.Error())
		return false
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		common.ErrLog("更新用户粉丝数时RowsAffected为0", err.Error())
		return false
	}
	return true
}
