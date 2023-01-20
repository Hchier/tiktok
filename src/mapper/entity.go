package mapper

import "time"

type Follow struct {
	Id         int32
	Follower   int32
	Followee   int32
	CreateTime time.Time
	Deleted    int8
}

type User struct {
	Id            int32
	Username      string
	Password      string
	FollowCount   int32
	FollowerCount int32
	Deleted       int8
}

type Video struct {
	Id            int32
	UserId        int32
	PlayUrl       string `comment:"视频地址"`
	CoverUrl      string `comment:"封面地址"`
	FavoriteCount int32
	CommentCount  int32
	Title         string
	PublishDate   time.Time
	Deleted       int8
}

type VideoComment struct {
	Id         int32
	UserId     int32
	VideoId    int32
	content    string
	CreateDate time.Time
	Deleted    int8
	CommentOf  int32
}

type VideoFavor struct {
	UserId     int32
	VideoId    int32
	UpdateTime time.Time
	Deleted    int8
}
