package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"tiktok/src/common"
	"tiktok/src/service"
)

//@author by Hchier
//@Date 2023/1/21 14:16

// VideoPublish 发布视频
func VideoPublish(ctx context.Context, c *app.RequestContext) {

	isValid, userId := common.IsValidUser(string(c.FormValue("token")), ctx)
	if !isValid {
		c.JSON(http.StatusOK, &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "身份验证失败",
		})
		return
	}

	c.JSON(http.StatusOK, service.PublishVideo(c, userId))
}

// ListOfPublishedVideo 发布列表
func ListOfPublishedVideo(ctx context.Context, c *app.RequestContext) {
	isValid, userId := common.IsValidUser(c.Query("token"), ctx)
	if !isValid {
		c.JSON(http.StatusOK, &common.ListOfPublishedVideoResp{
			StatusCode: -1,
			StatusMsg:  "身份验证失败",
		})
		return
	}
	c.JSON(http.StatusOK, service.GetListOfPublishedVideo(userId))
}

// ListOfFavoredVideo 点赞列表
func ListOfFavoredVideo(ctx context.Context, c *app.RequestContext) {
	isValid, userId := common.IsValidUser(c.Query("token"), ctx)
	if !isValid {
		c.JSON(http.StatusOK, &common.ListOfPublishedVideoResp{
			StatusCode: -1,
			StatusMsg:  "身份验证失败",
		})
		return
	}
	c.JSON(http.StatusOK, service.GetListOfFavoredVideo(userId))
}

func VideoFeed(ctx context.Context, c *app.RequestContext) {
	isValid, userId := common.IsValidUser(c.Query("token"), ctx)
	if !isValid {
		c.JSON(http.StatusOK, &common.ListOfPublishedVideoResp{
			StatusCode: -1,
			StatusMsg:  "身份验证失败",
		})
		return
	}
	c.JSON(http.StatusOK, service.VideoFeed(userId))
}
