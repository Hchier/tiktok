package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"tiktok/src/common"
	"tiktok/src/mapper"
)

//@author by Hchier
//@Date 2023/1/20 21:03

// Register
// 用户注册。首先检查用户名是否被占用。
func Register(username, password string) *common.UserRegisterOrLoginResp {
	if mapper.ExistUser(username) {
		return &common.UserRegisterOrLoginResp{
			StatusCode: -1,
			StatusMsg:  "username已被使用",
			UserId:     -1,
			Token:      "",
		}
	}

	id, statusMsg := mapper.InsertUser(username, password)
	if id < 0 {
		return &common.UserRegisterOrLoginResp{
			StatusCode: -1,
			StatusMsg:  statusMsg,
			UserId:     -1,
			Token:      "",
		}
	} else {
		return &common.UserRegisterOrLoginResp{
			StatusCode: 0,
			StatusMsg:  statusMsg,
			UserId:     id,
			Token:      SetToken(id, username),
		}
	}
}

// Login
// 用户登录
func Login(username, password string) *common.UserRegisterOrLoginResp {
	if id := mapper.GetIdByUsernameAndPassword(username, password); id < 0 {
		return &common.UserRegisterOrLoginResp{
			StatusCode: -1,
			StatusMsg:  "用户名与密码不匹配",
			UserId:     -1,
			Token:      "",
		}
	} else {

		return &common.UserRegisterOrLoginResp{
			StatusCode: 0,
			StatusMsg:  "登录成功",
			UserId:     id,
			Token:      SetToken(id, username),
		}
	}
}

func GetUserInfo(id int64) *common.UserInfoResp {
	userEntity := mapper.SelectUserById(id)
	return &common.UserInfoResp{
		StatusCode: 0,
		StatusMsg:  "ok",
		User: &struct {
			Id               int64  `json:"id"`
			Name             string `json:"name"`
			FollowCount      int64  `json:"follow_count"`
			FollowerCount    int64  `json:"follower_count"`
			IsFollow         bool   `json:"is_follow"`
			Avatar           string `json:"avatar"`
			Background_image string `json:"background_image"`
			Signature        string `json:"signature"`
			Total_favorited  int64  `json:"total_favorited"`
			Favorite_count   int64  `json:"favorite_count"`
		}{
			Id:               id,
			Name:             userEntity.Username,
			FollowCount:      userEntity.Follow_count,
			FollowerCount:    userEntity.Follower_count,
			IsFollow:         true,
			Avatar:           userEntity.Avatar,
			Background_image: userEntity.Background_image,
			Signature:        userEntity.Signature,
			Total_favorited:  userEntity.Total_favorited,
			Favorite_count:   userEntity.Favorite_count,
		},
	}
}

func SetToken(id int64, username string) string {
	var token string = username + common.GetRandStr()
	ctx := context.Background()
	_, err := common.Rdb.HSet(ctx, "tokens", id, token).Result()
	if err != nil {
		file := common.GetFile(common.ErrLogDest)
		defer file.Close()
		hlog.SetOutput(file)
		hlog.Error("将id-token放入redis失败：", err.Error())
	}
	return token
}
