package mapper

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"tiktok/src/common"
	"time"
)

//@author by Hchier
//@Date 2023/1/22 10:29

// OperationFollow 操作关注记录。opType -> 1：插入		2：删除
func OperationFollow(opType int8, follower, followee int64, tx *sqlx.Tx) bool {
	var res sql.Result
	var err error
	if opType == 1 {
		res, err = tx.Exec("insert into follow (follower, followee, create_time) values(?, ?, ?)", follower, followee, time.Now())
		if err != nil {
			common.ErrLog("插入关注记录失败：", err.Error())
			return false
		}
		count, _ := res.RowsAffected()
		if count == 0 {
			common.ErrLog("插入关注记录时RowsAffected为0：", err.Error())
			return false
		}
		return true
	} else {
		res, err = tx.Exec("update follow set deleted = 1 where follower = ? and followee = ? and deleted = 0", follower, followee)
		if err != nil {
			common.ErrLog("删除关注记录失败：", err.Error())
			return false
		}
		count, _ := res.RowsAffected()
		if count == 0 {
			common.ErrLog("删除关注记录时RowsAffected为0：", err.Error())
			return false
		}
		return true
	}
}

// ExistFollow 判断是否存在关注记录
func ExistFollow(follower, followee int64) bool {
	var count int32
	err := Db.Get(&count, "select count(*) from follow where follower = ? and followee = ? and deleted = 0", follower, followee)
	if err != nil {
		common.ErrLog("判断是否存在关注记录失败：", err.Error())
		return false
	}
	return count > 0
}

// GetFolloweeIdsByFollowerId 得到用户的偶像们的id
// 成功返回true
func GetFolloweeIdsByFollowerId(followerId int64) (bool, []int64) {
	var ids []int64
	err := Db.Select(&ids, "select followee from follow where follower = ? and deleted = 0", followerId)
	if err != nil {
		common.ErrLog("GetFolloweeListByFollowerId时出错：", ids)
		return false, nil
	}
	return true, ids
}

// GetFollowerIdsByFolloweeId 得到用户的粉丝们的id
// 成功返回true
func GetFollowerIdsByFolloweeId(followeeId int64) (bool, []int64) {
	var ids []int64
	err := Db.Select(&ids, "select follower from follow where followee = ? and deleted = 0", followeeId)
	if err != nil {
		common.ErrLog("GetFollowerIdsByFollowee时出错：", ids)
		return false, nil
	}
	return true, ids
}
