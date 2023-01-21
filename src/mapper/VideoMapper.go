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
		common.Log(common.ErrLogDest, "视频插入失败", err.Error())
		return
	}
}

func GetVideoListByUserId(userId int64) (bool, []Video) {
	var videos []Video
	err := Db.Select(&videos, "select * from video where user_id = ?", userId)
	if err != nil {
		common.Log("根据用户id查找视频失败", err.Error())
		return false, videos
	}
	return true, videos
}
