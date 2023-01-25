package service

import (
	"tiktok/src/common"
	"tiktok/src/mapper"
)

//@author by Hchier
//@Date 2023/1/21 21:28

// DoFavorVideo 视频点赞
// 4件事
// 1, 插入视频点赞信息  2, 更新视频获赞数  3, 更新作者的获赞数 4, 更新用户的点赞数
func DoFavorVideo(userId, videoId, authorId int64) *common.VideoFavorResp {
	tx, err := mapper.Db.Beginx()
	if err != nil {
		common.ErrLog("视频点赞时事务开启失败：", err.Error())
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "视频点赞时事务开启失败",
		}
	}

	//插入视频点赞信息
	if !mapper.OperateVideoFavor(1, userId, videoId, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("插入视频点赞信息时事务回滚失败：", err.Error())
		}
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "点赞失败",
		}
	}

	//更新视频获赞数
	if !mapper.UpdateVideoFavorCount(1, videoId, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("更新视频获赞数时事务回滚失败：", err.Error())
		}
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "点赞失败",
		}
	}

	//更新作者的获赞数
	if !mapper.UpdateUserTotalFavorited(1, authorId, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("更新作者的获赞数时事务回滚失败：", err.Error())
		}
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "点赞失败",
		}
	}

	//更新用户的点赞数
	if !mapper.UpdateUserFavoriteCount(1, userId, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("更新用户的点赞数时事务回滚失败：", err.Error())
		}
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "点赞失败",
		}
	}

	err = tx.Commit()
	if err != nil {
		common.ErrLog("视频点赞时事务提交失败：", err.Error())
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "视频点赞时事务提交失败",
		}
	}

	return &common.VideoFavorResp{
		StatusCode: 0,
		StatusMsg:  "点赞成功",
	}
}

// DoUnFavorVideo 视频取消点赞
// 4件事
// 1, 删除视频点赞信息  2, 更新视频获赞数  3, 更新作者的获赞数 4, 更新用户的点赞数
func DoUnFavorVideo(userId, videoId, authorId int64) *common.VideoFavorResp {
	tx, err := mapper.Db.Beginx()
	if err != nil {
		common.ErrLog("视频取消点赞时事务开启失败：", err.Error())
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "点赞失败",
		}
	}

	//删除视频点赞信息
	if !mapper.OperateVideoFavor(2, userId, videoId, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("删除视频点赞信息时事务回滚失败：", err.Error())
		}
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "点赞失败",
		}
	}

	//更新视频获赞数
	if !mapper.UpdateVideoFavorCount(2, videoId, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("更新视频获赞数时事务回滚失败：", err.Error())
		}
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "点赞失败",
		}
	}

	//更新作者的获赞数
	if !mapper.UpdateUserTotalFavorited(2, authorId, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("更新作者的获赞数时事务回滚失败：", err.Error())
		}
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "点赞失败",
		}
	}

	//更新用户的点赞数
	if !mapper.UpdateUserFavoriteCount(2, userId, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("更新用户的点赞数时事务回滚失败：", err.Error())
		}
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "点赞失败",
		}
	}

	err = tx.Commit()
	if err != nil {
		common.ErrLog("取消视频点赞时事务提交失败：", err.Error())
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "取消视频点赞时事务提交失败",
		}
	}

	return &common.VideoFavorResp{
		StatusCode: 0,
		StatusMsg:  "取消点赞成功",
	}
}
