package common

//@author by Hchier
//@Date 2023/1/21 0:29

type UserRegisterOrLoginResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

type UserInfoResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	User       *struct {
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
	} `json:"user"`
}

type VideoPublishResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type VideoVo struct {
	Id     int64 `json:"id"`
	Author struct {
		Id            int64  `json:"id"`
		Name          string `json:"name"`
		FollowCount   int64  `json:"follow_count"`
		FollowerCount int64  `json:"follower_count"`
		Avatar        string `json:"avatar"`
		IsFollow      bool   `json:"is_follow"`
	} `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type ListOfPublishedVideoResp struct {
	StatusCode int32     `json:"status_code"`
	StatusMsg  string    `json:"status_msg"`
	VideoList  []VideoVo `json:"video_list"`
}
