package mapper

import (
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

// GetPublishedVideoListByUserId
// 返回用户发布的视频列表
func GetPublishedVideoListByUserId(userId int64) (bool, []Video) {
	var videos []Video
	err := Db.Select(&videos, "select * from video where user_id = ?", userId)
	if err != nil {
		common.ErrLog("根据用户id查找视频失败", err.Error())
		return false, videos
	}
	return true, videos
}

// SelectAuthorIdByVideoId
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
