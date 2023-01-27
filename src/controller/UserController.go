package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"math/rand"
	"net/http"
	"strconv"
	"tiktok/src/common"
	"tiktok/src/service"
)

//@author by Hchier
//@Date 2023/1/20 22:54

func UserRegister(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := common.Md5Encode(c.Query("password"))

	if username == "" || password == "" {
		c.JSON(http.StatusOK, &common.UserRegisterOrLoginResp{
			StatusCode: -1,
			StatusMsg:  "username或password为空",
			UserId:     -1,
			Token:      "",
		})
		return
	}
	avatar := common.AvatarPathPrefix + strconv.Itoa(rand.Intn(10)) + ".png"
	backgroundImage := common.BackgroundImagePathPrefix + strconv.Itoa(rand.Intn(10)) + ".png"
	signature := common.Signatures[rune(rand.Intn(10))]
	c.JSON(http.StatusOK, service.Register(username, password, avatar, backgroundImage, signature))
}

func UserLogin(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := common.Md5Encode(c.Query("password"))
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
	val, _ := c.Get("id")
	currentUserId := val.(int64)
	targetUserIdStr := c.Query("user_id")
	targetUserId, err := strconv.ParseInt(targetUserIdStr, 10, 64)
	if err != nil {
		print(err.Error())
	}

	c.JSON(http.StatusOK, service.GetUserInfo(targetUserId, currentUserId))
}
