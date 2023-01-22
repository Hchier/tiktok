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

	isValid, userId := common.IsValidUser(string(c.FormValue("token")), ctx)
	if !isValid {
		c.JSON(http.StatusOK, &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "身份验证失败",
		})
		return
	}

	video_id, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	action_type := c.Query("action_type")
	comment_text := c.Query("comment_text")
	comment_id, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)

	//发布评论
	if action_type == "1" {
		c.JSON(http.StatusOK, service.DoPublishVideoComment(video_id, comment_text, userId, comment_id))
	} else { //删除评论
		c.JSON(http.StatusOK, service.DoDeleteVideoComment(video_id, comment_text, userId, comment_id))
	}
}

func VideoCommentList(ctx context.Context, c *app.RequestContext) {
	isValid, userId := common.IsValidUser(string(c.FormValue("token")), ctx)
	if !isValid {
		c.JSON(http.StatusOK, &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "身份验证失败",
		})
		return
	}

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	c.JSON(http.StatusOK, service.GetVideoCommentByVideoId(videoId, userId))
}
