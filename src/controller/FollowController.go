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
	isValid, userId := common.IsValidUser(string(c.FormValue("token")), ctx)
	if !isValid {
		c.JSON(http.StatusOK, &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "身份验证失败",
		})
		return
	}
	actionType := c.Query("action_type")
	followee, err := strconv.ParseInt(c.Query("to_unser_id"), 10, 64)
	if err != nil {
		common.ErrLog("str to int fail：", err.Error())
		c.JSON(http.StatusOK, &common.FollowActionResp{StatusCode: -1, StatusMsg: "fail"})
	}
	if actionType == "1" {
		c.JSON(http.StatusOK, service.DoFollow(userId, followee))
	} else {
		c.JSON(http.StatusOK, service.DoUnFollow(userId, followee))
	}
}
