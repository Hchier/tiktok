package controller

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

//@author by Hchier
//@Date 2023/1/20 22:51

var Hertz *server.Hertz

func init() {
	Hertz = server.Default(server.WithHostPorts("192.168.0.105:8080"))

	Hertz.POST("/douyin/user/register/", UserRegister)
	Hertz.POST("/douyin/user/login/", UserLogin)
	Hertz.GET("/douyin/user/", UserInfo)
}
