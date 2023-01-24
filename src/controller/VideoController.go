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
//@Date 2023/1/21 14:16

// VideoPublish 发布视频
func VideoPublish(ctx context.Context, c *app.RequestContext) {
	val, _ := c.Get("id")
	if currentUserId := val.(int64); currentUserId == -1 {
		c.JSON(http.StatusOK, &common.BasicResp{
			StatusCode: -1,
			StatusMsg:  "未登录",
		})
	} else {
		c.JSON(http.StatusOK, service.PublishVideo(c, currentUserId))
	}
}

// ListOfPublishedVideo 发布列表
func ListOfPublishedVideo(ctx context.Context, c *app.RequestContext) {
	val, _ := c.Get("id")
	currentUserId := val.(int64)
	targetUserIdStr := c.Query("user_id")
	targetUserId, _ := strconv.ParseInt(targetUserIdStr, 10, 64)
	c.JSON(http.StatusOK, service.GetListOfPublishedVideo(currentUserId, targetUserId))
}

// ListOfFavoredVideo 点赞列表
func ListOfFavoredVideo(ctx context.Context, c *app.RequestContext) {
	val, _ := c.Get("id")
	currentUserId := val.(int64)
	targetUserIdStr := c.Query("user_id")
	targetUserId, _ := strconv.ParseInt(targetUserIdStr, 10, 64)
	c.JSON(http.StatusOK, service.GetListOfFavoredVideo(currentUserId, targetUserId))
}

// VideoFeed 视频流
func VideoFeed(ctx context.Context, c *app.RequestContext) {
	val, _ := c.Get("id")
	currentUserId := val.(int64)
	c.JSON(http.StatusOK, service.VideoFeed(currentUserId))
}
