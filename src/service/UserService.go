package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"tiktok/src/common"
	"tiktok/src/mapper"
	"time"
)

//@author by Hchier
//@Date 2023/1/20 21:03

// Register
// 用户注册。首先检查用户名是否被占用。
func Register(username, password, avatar, backgroundImage, signature string) *common.UserRegisterOrLoginResp {
	if mapper.ExistUserByUsername(username) {
		return &common.UserRegisterOrLoginResp{
			StatusCode: -1,
			StatusMsg:  "username已被使用",
			UserId:     -1,
			Token:      "",
		}
	}

	id, statusMsg := mapper.InsertUser(username, password, avatar, backgroundImage, signature)
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

func GetUserInfo(targetUserId, currentUserId int64) *common.UserInfoResp {
	userEntity := mapper.SelectUserById(targetUserId)
	if userEntity.Id == -1 {
		return &common.UserInfoResp{
			StatusCode: -1,
			StatusMsg:  "用户不存在",
		}
	}
	return &common.UserInfoResp{
		StatusCode: 0,
		StatusMsg:  "ok",
		User: &struct {
			Id              int64  `json:"id"`
			Name            string `json:"name"`
			FollowCount     int64  `json:"follow_count"`
			FollowerCount   int64  `json:"follower_count"`
			IsFollow        bool   `json:"is_follow"`
			Avatar          string `json:"avatar"`
			BackgroundImage string `json:"background_image"`
			Signature       string `json:"signature"`
			TotalFavorited  int64  `json:"total_favorited"`
			FavoriteCount   int64  `json:"favorite_count"`
			VideoCount      int64  `json:"video_count"`
		}{
			Id:              userEntity.Id,
			Name:            userEntity.Username,
			FollowCount:     userEntity.Follow_count,
			FollowerCount:   userEntity.Follower_count,
			IsFollow:        mapper.ExistFollow(currentUserId, targetUserId),
			Avatar:          common.StaticResources + userEntity.Avatar,
			BackgroundImage: common.StaticResources + userEntity.Background_image,
			Signature:       userEntity.Signature,
			TotalFavorited:  userEntity.Total_favorited,
			FavoriteCount:   userEntity.Favorite_count,
			VideoCount:      userEntity.Video_count,
		},
	}
}

// SetToken 2件事。1，将<token, id>放入tokens(Hash)中。 2，将<token, expireTime>放入expireTime(Zset)中
func SetToken(id int64, username string) string {
	var token string = username + "-" + uuid.NewV4().String()
	ctx := context.Background()

	_, err := common.Rdb.HSet(ctx, "tokens", token, id).Result()
	if err != nil {
		common.ErrLog("将token-id放入redis失败：", err.Error())
	}

	_, err = common.Rdb.ZAdd(ctx, "expireTime", &redis.Z{Member: token, Score: float64(time.Now().Add(time.Minute * 5).Unix())}).Result()
	if err != nil {
		common.ErrLog("将token-expireTime放入redis失败：", err.Error())
	}

	return token
}

// RemoveExpiredToken 移除过期的token。2件事：1，从tokens(Hash)中移除过期的<token, id>。 2，从expireTime(Zset)中移除过期的<token, expireTime>
func RemoveExpiredToken() {
	println("RemoveExpiredToken")
	timeNowStr := strconv.FormatInt(time.Now().Unix(), 10)
	// 得到已经过期的tokens
	tokens, err := common.Rdb.ZRangeByScore(context.Background(), "expireTime", &redis.ZRangeBy{Min: string(rune(0)), Max: timeNowStr}).Result()
	if err != nil {
		common.ErrLog(err.Error())
	}
	//
	count, err := common.Rdb.HDel(context.Background(), "tokens", tokens...).Result()
	if err != nil {
		common.ErrLog(err.Error())
	}
	println("从tokens中移除", count, "个")

	//
	count, err = common.Rdb.ZRemRangeByScore(context.Background(), "expireTime", string(rune(0)), timeNowStr).Result()
	if err != nil {
		println(err.Error())
	}
	println("从expireTime中移除", count, "个")
}
