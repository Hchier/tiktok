package mapper

import "time"

type Follow struct {
	Id         int64
	Follower   int64
	Followee   int64
	CreateTime time.Time
	Deleted    int8
}

type User struct {
	Id               int64
	Username         string
	Password         string
	Follow_count     int64
	Follower_count   int64
	Avatar           string
	Background_image string
	Signature        string
	Total_favorited  int64
	Favorite_count   int64
	Deleted          int8
}

type Video struct {
	Id            int64
	UserId        int64
	PlayUrl       string `comment:"视频地址"`
	CoverUrl      string `comment:"封面地址"`
	FavoriteCount int64
	CommentCount  int64
	Title         string
	PublishDate   time.Time
	Deleted       int8
}

type VideoComment struct {
	Id         int64
	UserId     int64
	VideoId    int64
	content    string
	CreateDate time.Time
	Deleted    int8
	CommentOf  int64
}

type VideoFavor struct {
	UserId     int64
	VideoId    int64
	UpdateTime time.Time
	Deleted    int8
}
