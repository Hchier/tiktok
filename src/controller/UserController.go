package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"tiktok/src/service"
)

//@author by Hchier
//@Date 2023/1/20 22:54

type UserRegisterResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int32  `json:"user_id"`
	Token      string `json:"token"`
}

func UserRefresh() {
	Hertz.POST("/douyin/user/register/", func(ctx context.Context, c *app.RequestContext) {
		username := c.Query("username")
		password := c.Query("password")
		if username == "" || password == "" {
			c.JSON(http.StatusOK, &UserRegisterResp{
				StatusCode: -1,
				StatusMsg:  "username或password为空",
				UserId:     -1,
				Token:      "",
			})
			return
		}
		id, statusMsg := service.Register(username, password)
		if id >= 0 {
			c.JSON(http.StatusOK, UserRegisterResp{
				StatusCode: 0,
				StatusMsg:  statusMsg,
				UserId:     id,
				Token:      "",
			})
		} else {
			c.JSON(http.StatusOK, UserRegisterResp{
				StatusCode: id,
				StatusMsg:  statusMsg,
				UserId:     id,
				Token:      "",
			})
		}
	})
}
