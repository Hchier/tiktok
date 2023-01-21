package service

import (
	"tiktok/src/common"
	"tiktok/src/mapper"
)

//@author by Hchier
//@Date 2023/1/21 21:28

func DoFavorVideo(userId, videoId, authorId int64) *common.VideoFavorResp {
	res := mapper.InsertVideoFavor(userId, videoId, authorId)
	//未成功插入
	if !res {
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "点赞失败",
		}
	}
	return &common.VideoFavorResp{
		StatusCode: 0,
		StatusMsg:  "点赞成功",
	}
}

func DoUnFavorVideo(userId, videoId, authorId int64) *common.VideoFavorResp {
	res := mapper.DeleteVideoFavor(userId, videoId, authorId)
	//未成功取消
	if !res {
		return &common.VideoFavorResp{
			StatusCode: -1,
			StatusMsg:  "取消点赞失败",
		}
	}
	return &common.VideoFavorResp{
		StatusCode: 0,
		StatusMsg:  "取消点赞成功",
	}
}
