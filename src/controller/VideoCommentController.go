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
//@Date 2023/1/22 13:11

func VideoCommentAction(ctx context.Context, c *app.RequestContext) {
	val, _ := c.Get("id")
	currentUserId := val.(int64)
	if currentUserId == -1 {
		c.JSON(http.StatusOK, &common.BasicResp{
			StatusCode: -1,
			StatusMsg:  "未登录",
		})
	}

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)

	//发布评论
	if actionType == "1" {
		c.JSON(http.StatusOK, service.DoPublishVideoComment(videoId, commentText, currentUserId, commentId))
	} else { //删除评论
		c.JSON(http.StatusOK, service.DoDeleteVideoComment(videoId, commentText, currentUserId, commentId))
	}
}

func VideoCommentList(ctx context.Context, c *app.RequestContext) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	c.JSON(http.StatusOK, service.GetVideoCommentByVideoId(videoId))
}
