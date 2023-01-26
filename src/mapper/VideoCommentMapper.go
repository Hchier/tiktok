package mapper

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"tiktok/src/common"
	"time"
)

//@author by Hchier
//@Date 2023/1/22 13:15

// OperateVideoComment 操控视频评论。opType -> 1：插入		2：删除
// 插入成功返回(true, commentId)
// 删除成功返回(true, -1)
func OperateVideoComment(opType int8, videoId int64, content string, userId int64, commentId int64, tx *sqlx.Tx) (bool, int64) {
	var res sql.Result
	var err error
	if opType == 1 {
		res, err = tx.Exec("insert into video_comment (user_id, video_id, content, create_date) values (?, ?, ?, ?)", userId, videoId, content, time.Now())
		if err != nil {
			common.ErrLog("插入视频评论失败：", err.Error())
			return false, -1
		}
		count, _ := res.RowsAffected()
		if count == 0 {
			common.ErrLog("插入视频评论时RowsAffected为0：", err.Error())
			return false, -1
		}
		commentId, _ := res.LastInsertId()
		return true, commentId
	} else {
		res, err = tx.Exec("update video_comment set deleted = 1 where id = ? and user_id = ? and deleted = 0", commentId, userId)
		if err != nil {
			common.ErrLog("删除视频评论失败：", err.Error())
			return false, -1
		}
		count, _ := res.RowsAffected()
		if count == 0 {
			common.ErrLog("删除视频评论时RowsAffected为0")
			return false, -1
		}
		return true, -1
	}
}

// GetVideoCommentByVideoId 根据videoId查评论
func GetVideoCommentByVideoId(videoId int64) (bool, []VideoComment) {
	var videoCommentList []VideoComment
	err := common.Db.Select(&videoCommentList, "select * from video_comment where video_id = ? and deleted = 0", videoId)
	if err != nil {
		common.ErrLog("根据videoId查评论失败：", videoCommentList)
		return false, nil
	}
	return true, videoCommentList
}
