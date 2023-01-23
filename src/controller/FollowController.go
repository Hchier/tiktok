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
//@Date 2023/1/22 20:58

func FollowOperation(ctx context.Context, c *app.RequestContext) {
	val, _ := c.Get("id")
	currentUserId := val.(int64)
	if currentUserId == -1 {
		c.JSON(http.StatusOK, &common.BasicResp{
			StatusCode: -1,
			StatusMsg:  "未登录",
		})
	}
	actionType := c.Query("action_type")
	followee, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		common.ErrLog("str to int fail(FollowOperation c.Query(\"to_user_id\"))：", err.Error())
		c.JSON(http.StatusOK, &common.FollowActionResp{StatusCode: -1, StatusMsg: "fail"})
	}
	if actionType == "1" {
		c.JSON(http.StatusOK, service.DoFollow(currentUserId, followee))
	} else {
		c.JSON(http.StatusOK, service.DoUnFollow(currentUserId, followee))
	}
}

// FolloweeList 偶像列表
func FolloweeList(ctx context.Context, c *app.RequestContext) {
	val, _ := c.Get("id")
	currentUserId := val.(int64)
	targetUserIdStr := c.Query("user_id")
	targetUserId, _ := strconv.ParseInt(targetUserIdStr, 10, 64)

	c.JSON(http.StatusOK, service.GetFolloweeInfo(currentUserId, targetUserId))
}

// FollowerList 粉丝列表
func FollowerList(ctx context.Context, c *app.RequestContext) {
	val, _ := c.Get("id")
	currentUserId := val.(int64)
	targetUserIdStr := c.Query("user_id")
	targetUserId, _ := strconv.ParseInt(targetUserIdStr, 10, 64)

	c.JSON(http.StatusOK, service.GetFollowerInfo(currentUserId, targetUserId))
}
