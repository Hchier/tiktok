package mapper

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"tiktok/src/common"
	"time"
)

//@author by Hchier
//@Date 2023/1/20 20:19

func InsertVideo(userId int64, playUrl string, coverUrl string, title string) {
	_, err := Db.Exec("insert into video(user_id, play_url, cover_url, title, publish_date) values(?,?,?,?,?)", userId, playUrl, coverUrl, title, time.Now())
	if err != nil {
		common.ErrLog("视频插入失败", err.Error())
		return
	}
}

// GetPublishedVideoListByUserId 查询用户发布的视频列表
// 成功：true, 用户发布的视频列表
// 失败：false,
func GetPublishedVideoListByUserId(userId int64) (bool, []Video) {
	var videos []Video
	err := Db.Select(&videos, "select * from video where user_id = ?", userId)
	if err != nil {
		common.ErrLog("根据用户id查找视频失败", err.Error())
		return false, videos
	}
	return true, videos
}

// SelectAuthorIdByVideoId 根据视频id查找作者id
// 成功查到：authorId
// 查不到：-1
func SelectAuthorIdByVideoId(videoId int64) int64 {
	var authorId int64
	err := Db.Get(&authorId, "select user_id from video where id = ?", videoId)
	if err != nil {
		common.ErrLog("根据视频id查找作者id时出错", err.Error())
		return -1
	}
	return authorId
}

// GetFavoredVideoListByUserId 查询用户点赞的视频列表
func GetFavoredVideoListByUserId(userId int64) (bool, []Video) {
	var videos []Video
	err := Db.Select(&videos, "select * from video where id in (select video_id from video_favor where user_id = ? and deleted = 0)", userId)
	if err != nil {
		common.ErrLog("查询用户点赞的视频列表失败", err.Error())
		return false, videos
	}
	return true, videos
}

// UpdateVideoFavorCount 更新视频获赞数。opType -> 1：加1；2：减1
// 成功返回true
func UpdateVideoFavorCount(opType int8, videoId int64, tx *sqlx.Tx) bool {
	var res sql.Result
	var err error
	if opType == 1 {
		res, err = tx.Exec("update video set favorite_count = favorite_count + 1 where id = ?", videoId)
	} else {
		res, err = tx.Exec("update video set favorite_count = favorite_count - 1 where id = ?", videoId)
	}
	if err != nil {
		common.ErrLog("更新视频点赞数失败：", err.Error())
		return false
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		common.ErrLog("更新视频点赞数时RowsAffected为0", err.Error())
		return false
	}
	return true
}

// UpdateVideoCommentCount 更新视频评论数。opType -> 1：加1；2：减1
func UpdateVideoCommentCount(opType int8, videoId int64, tx *sqlx.Tx) bool {
	var res sql.Result
	var err error
	if opType == 1 {
		res, err = tx.Exec("update video set comment_count = comment_count + 1 where id = ?", videoId)
	} else {
		res, err = tx.Exec("update video set comment_count = comment_count - 1 where id = ?", videoId)
	}
	if err != nil {
		common.ErrLog("更新视频评论数失败：", err.Error())
		return false
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		common.ErrLog("更新视频评论数时RowsAffected为0", err.Error())
		return false
	}
	return true
}
