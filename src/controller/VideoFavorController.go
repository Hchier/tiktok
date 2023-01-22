package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
	"tiktok/src/common"
	"tiktok/src/service"
)

//@author by Hchier
//@Date 2023/1/21 21:54

// 点赞或取消点赞
func VideoFavor(ctx context.Context, c *app.RequestContext) {
	isValid, userId := common.IsValidUser(c.Query("token"), ctx)
	if !isValid {
		c.JSON(http.StatusOK, &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "身份验证失败",
		})
		return
	}

	video_id, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		common.ErrLog("string to int 出错", err.Error())
		c.JSON(http.StatusOK, &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "string to int 出错",
		})
		return
	}

	authorId := service.GetAuthorIdByVideoId(video_id)
	if authorId == -1 {
		c.JSON(http.StatusOK, &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "无法根据VideoId找到AuthorId",
		})
		return
	}
	action_type := c.Query("action_type")
	//点赞
	if action_type == "1" {
		c.JSON(http.StatusOK, service.DoFavorVideo(userId, video_id, authorId))
	} else { //取消点赞
		c.JSON(http.StatusOK, service.DoUnFavorVideo(userId, video_id, authorId))
	}
}
