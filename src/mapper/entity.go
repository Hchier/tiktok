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
	Video_count      int64
	Deleted          int8
}

type Video struct {
	Id             int64
	User_id        int64
	Play_url       string `comment:"视频地址"`
	Cover_url      string `comment:"封面地址"`
	Favorite_count int64
	Comment_count  int64
	Title          string
	Publish_date   time.Time
	Deleted        int8
}

type VideoComment struct {
	Id          int64
	User_id     int64
	Video_id    int64
	content     string
	Create_date time.Time
	Deleted     int8
	Comment_of  int64
}

type VideoFavor struct {
	UserId     int64
	VideoId    int64
	UpdateTime time.Time
	Deleted    int8
}
