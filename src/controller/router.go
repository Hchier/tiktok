package controller

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"tiktok/src/common"
)

//@author by Hchier
//@Date 2023/1/20 22:51

var Hertz *server.Hertz

func init() {
	Hertz = server.Default(server.WithHostPorts("192.168.0.105:8080"))

	//登录时禁止访问
	Hertz.POST("/douyin/user/register/", UserRegister)
	Hertz.POST("/douyin/user/login/", UserLogin)

	//
	Hertz.Group("").Use(common.IsValidUser).
		//未登录时可以访问，但登陆与否所见不同
		GET("/douyin/feed/", VideoFeed).
		GET("/douyin/user/", UserInfo).
		GET("/douyin/publish/list/", ListOfPublishedVideo).
		GET("/douyin/favorite/list/", ListOfFavoredVideo).
		GET("/douyin/relation/follow/list/", FolloweeList).
		GET("/douyin/relation/follower/list/", FollowerList).
		GET("/douyin/relation/friend/list/", FollowerList).

		//未登录时可以访问，但登陆与否所见相同
		GET("/douyin/comment/list/", VideoCommentList).

		//未登录时禁止访问
		POST("/douyin/publish/action/", VideoPublish).
		POST("/douyin/favorite/action/", VideoFavor).
		POST("/douyin/comment/action/", VideoCommentAction).
		POST("/douyin/relation/action/", FollowOperation)

}
