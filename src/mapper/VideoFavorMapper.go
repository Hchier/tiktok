package mapper

import (
	"tiktok/src/common"
	"time"
)

//@author by Hchier
//@Date 2023/1/21 21:21

// InsertVideoFavor 插入视频点赞信息
// 3件事
// 1, 插入视频点赞信息  2, 更新视频获赞数  3, 更新作者的获赞数 4, 更新用户的点赞数
// 成功插入返回true
func InsertVideoFavor(userId, videoId, authorId int64) bool {
	tx, err := Db.Begin()
	if err != nil {
		common.ErrLog("插入视频点赞记录时开启事务失败", err.Error())
		return false
	}

	////插入视频点赞信息。在之前，要先确认不存在点赞记录。
	//var count int64
	//_ = Db.Get(&count, "select count(*) from video_favor where user_id = ? and video_id = ? and deleted = 0", userId, videoId)
	//if count > 0 {
	//	//已点赞，不能再点了
	//	return false
	//}
	//res, err := tx.Exec("insert into video_favor (user_id, video_id, update_time) values (?, ?, ?)", userId, videoId, time.Now())
	//if err != nil {
	//	common.ErrLog("插入视频点赞信息失败", err.Error())
	//	return false
	//}
	//count, _ = res.RowsAffected()
	//if count == 0 {
	//	common.ErrLog("插入视频点赞时RowsAffected为0", err.Error())
	//	return false
	//}

	//插入视频点赞信息。在之前，要先确认不存在点赞记录。
	res, err := tx.Exec("update video_favor set deleted = 0 where user_id = ? and video_id = ?", userId, videoId)
	if err != nil {
		common.ErrLog("更新视频点赞信息失败：", err.Error())
		return false
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		res, err = tx.Exec("insert into video_favor (user_id, video_id, update_time) values (?, ?, ?)", userId, videoId, time.Now())
		if err != nil {
			common.ErrLog("插入视频点赞信息失败：", err.Error())
			return false
		}
		count, _ = res.RowsAffected()
		if count == 0 {
			common.ErrLog("插入视频点赞时RowsAffected为0：", err.Error())
			return false
		}
	}

	//更新视频获赞数
	res, err = tx.Exec("update video set favorite_count = favorite_count + 1 where id = ?", videoId)
	if err != nil {
		common.ErrLog("更新视频点赞数失败：", err.Error())
		return false
	}
	count, _ = res.RowsAffected()
	if count == 0 {
		common.ErrLog("更新视频点赞数时RowsAffected为0", err.Error())
		return false
	}

	//更新作者的获赞数
	res, err = tx.Exec("update user set total_favorited = total_favorited + 1 where id = ?", authorId)
	if err != nil {
		common.ErrLog("更新用户点赞数失败", err.Error())
		return false
	}
	count, _ = res.RowsAffected()
	if count == 0 {
		common.ErrLog("更新用户点赞数时RowsAffected为0", err.Error())
		return false
	}

	//更新用户的点赞数
	res, err = tx.Exec("update user set favorite_count = favorite_count + 1 where id = ?", authorId)
	if err != nil {
		common.ErrLog("更新用户点赞数失败", err.Error())
		return false
	}
	count, _ = res.RowsAffected()
	if count == 0 {
		common.ErrLog("更新用户点赞数时RowsAffected为0", err.Error())
		return false
	}

	err = tx.Commit()
	return true
}

// DeleteVideoFavor 删除视频点赞信息
// 3件事
// 1, 删除视频点赞信息  2, 更新视频获赞数  3, 更新作者的获赞数  4, 更新用户的点赞数
// 成功插入返回true
func DeleteVideoFavor(userId, videoId, authorId int64) bool {
	tx, err := Db.Begin()
	if err != nil {
		common.ErrLog("删除视频点赞时开启事务失败", err.Error())
		return false
	}

	//删除视频点赞信息
	res, err := tx.Exec("update video_favor set deleted = 1 where user_id = ? and video_id = ?", userId, videoId)
	if err != nil {
		common.ErrLog("删除视频点赞信息失败：", err.Error())
		return false
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		common.ErrLog("不存在点赞记录，无法取消：", err.Error())
		return false
	}

	//更新视频获赞数
	res, err = tx.Exec("update video set favorite_count = favorite_count - 1 where id = ?", videoId)
	if err != nil {
		common.ErrLog("更新视频点赞数失败", err.Error())
		return false
	}
	count, _ = res.RowsAffected()
	if count == 0 {
		common.ErrLog("更新视频点赞数时RowsAffected为0", err.Error())
		return false
	}

	//更新作者的获赞数
	res, err = tx.Exec("update user set total_favorited = total_favorited - 1 where id = ?", authorId)
	if err != nil {
		common.ErrLog("更新作者获赞数失败", err.Error())
		return false
	}
	count, _ = res.RowsAffected()
	if count == 0 {
		common.ErrLog("更新作者获赞数时RowsAffected为0", err.Error())
		return false
	}

	//更新用户的点赞数
	res, err = tx.Exec("update user set favorite_count = favorite_count - 1 where id = ?", authorId)
	if err != nil {
		common.ErrLog("更新用户点赞数失败", err.Error())
		return false
	}
	count, _ = res.RowsAffected()
	if count == 0 {
		common.ErrLog("更新用户点赞数时RowsAffected为0", err.Error())
		return false
	}

	err = tx.Commit()
	if err != nil {
		common.ErrLog("删除视频点赞信息事务执行失败：", err.Error())
	}
	return true
}

// CheckVideoFavorByUserIdAndVideoId 检查某用户是否点赞了某视频
// 若点了，返回true
func CheckVideoFavorByUserIdAndVideoId(userId, videoId int64) bool {
	var count int64
	_ = Db.Get(&count, "select count(*) from video_favor where user_id = ? and video_id = ? and deleted = 0", userId, videoId)
	return count > 0
}
