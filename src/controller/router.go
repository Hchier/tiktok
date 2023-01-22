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
	Hertz.POST("/douyin/publish/action/", VideoPublish)
	Hertz.GET("/douyin/publish/list/", ListOfPublishedVideo)
	Hertz.GET("/douyin/feed/", ListOfPublishedVideo)

	Hertz.POST("/douyin/favorite/action/", VideoFavor)
	Hertz.GET("/douyin/favorite/list/", ListOfFavoredVideo)
	Hertz.POST("//douyin/comment/action/", VideoCommentAction)
	Hertz.GET("/douyin/comment/list/", VideoCommentList)
}
