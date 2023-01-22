package mapper

import (
	"tiktok/src/common"
	"time"
)

//@author by Hchier
//@Date 2023/1/22 10:29

// InsertFollow 插入关注记录
func InsertFollow(follower, followee int64) {
	_, err := Db.Exec("insert into follow (follower, followee, created_time) values(?, ?, ?)", follower, followee, time.Now())
	if err != nil {
		common.ErrLog("插入关注记录失败：", err.Error())
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
