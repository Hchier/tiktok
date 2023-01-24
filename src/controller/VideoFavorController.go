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

// VideoFavor 赞操作。（点赞或取消）
func VideoFavor(ctx context.Context, c *app.RequestContext) {
	val, _ := c.Get("id")
	currentUserId := val.(int64)
	if currentUserId == -1 {
		c.JSON(http.StatusOK, &common.BasicResp{
			StatusCode: -1,
			StatusMsg:  "未登录",
		})
	}

	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		common.ErrLog("string to int 出错", err.Error())
		c.JSON(http.StatusOK, &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "string to int 出错",
		})
		return
	}

	authorId := service.GetAuthorIdByVideoId(videoId)
	if authorId == -1 {
		c.JSON(http.StatusOK, &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "无法根据VideoId找到AuthorId",
		})
		return
	}
	actionType := c.Query("action_type")
	//点赞
	if actionType == "1" {
		c.JSON(http.StatusOK, service.DoFavorVideo(currentUserId, videoId, authorId))
	} else { //取消点赞
		c.JSON(http.StatusOK, service.DoUnFavorVideo(currentUserId, videoId, authorId))
	}
}
