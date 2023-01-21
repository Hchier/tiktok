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
//@Date 2023/1/20 22:54

func UserRegister(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, &common.UserRegisterOrLoginResp{
			StatusCode: -1,
			StatusMsg:  "username或password为空",
			UserId:     -1,
			Token:      "",
		})
		return
	}
	c.JSON(http.StatusOK, service.Register(username, password))
}

func UserLogin(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, &common.UserRegisterOrLoginResp{
			StatusCode: -1,
			StatusMsg:  "username或password为空",
			UserId:     -1,
			Token:      "err",
		})
		return
	}
	c.JSON(http.StatusOK, service.Login(username, password))

}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	userId := c.Query("user_id")
	token := c.Query("token")
	res, err := common.Rdb.HGet(ctx, "tokens", token).Result()
	if err != nil {
		common.ErrLog("查找token失败：", err.Error())
		c.JSON(http.StatusOK, &common.UserInfoResp{
			StatusCode: -1,
			StatusMsg:  "查找token失败",
		})
		return
	}
	if userId != res {
		c.JSON(http.StatusOK, &common.UserInfoResp{
			StatusCode: -1,
			StatusMsg:  "token不匹配",
		})
		return
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		print(err.Error())
	}
	c.JSON(http.StatusOK, service.GetUserInfo(id))
}
