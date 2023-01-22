package mapper

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"tiktok/src/common"
	"time"
)

//@author by Hchier
//@Date 2023/1/21 21:21

// OperateVideoFavor 操控视频点赞记录。opType -> 1：插入点赞记录；2：删除点赞记录
// 成功返回true
func OperateVideoFavor(opType int8, userId, videoId int64, tx *sqlx.Tx) bool {
	exist := ExistVideoFavor(userId, videoId)
	//插入点赞记录时，但是已经点赞了
	if opType == 1 && exist {
		return false
	}
	//删除点赞记录，却没点赞
	if opType == 2 && !exist {
		return false
	}

	var res sql.Result
	var err error
	if opType == 1 {
		res, err = tx.Exec("insert into video_favor (user_id, video_id, update_time) values (?, ?, ?)", userId, videoId, time.Now())
	} else {
		res, err = tx.Exec("update video_favor set deleted = 1 where user_id = ? and video_id = ?", userId, videoId)
	}
	if err != nil {
		common.ErrLog("操控视频点赞记录失败：", err.Error())
		return false
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		common.ErrLog("操控视频点赞记录时RowsAffected为0：", err.Error())
		return false
	}
	return true
}

// ExistVideoFavor 检查某用户是否点赞了某视频
// 若点了，返回true
func ExistVideoFavor(userId, videoId int64) bool {
	var count int64
	_ = Db.Get(&count, "select count(*) from video_favor where user_id = ? and video_id = ? and deleted = 0", userId, videoId)
	return count > 0
}
