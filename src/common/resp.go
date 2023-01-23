package common

//@author by Hchier
//@Date 2023/1/21 0:29

type BasicResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

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
	} `json:"user"`
}

type VideoPublishResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type VideoVo struct {
	Id     int64 `json:"id"`
	Author struct {
		Id              int64  `json:"id"`
		Name            string `json:"name"`
		FollowCount     int64  `json:"follow_count"`
		FollowerCount   int64  `json:"follower_count"`
		Avatar          string `json:"avatar"`
		BackgroundImage string `json:"background_image"`
		Signature       string `json:"signature"`
		TotalFavorited  int64  `json:"total_favorited"`
		FavoriteCount   int64  `json:"favorite_count"`
		VideoCount      int64  `json:"video_count"`
		IsFollow        bool   `json:"is_follow"`
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

type VideoFavorResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type CommentVo struct {
	Id   int64 `json:"id"`
	User struct {
		Id            int64  `json:"id"`
		Name          string `json:"name"`
		Avatar        string `json:"avatar"`
		FollowCount   int64  `json:"follow_count"`
		FollowerCount int64  `json:"follower_count"`
		IsFollow      bool   `json:"is_follow"`
	} `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type CommentActionResp struct {
	StatusCode int32     `json:"status_code"`
	StatusMsg  string    `json:"status_msg"`
	Comment    CommentVo `json:"comment"`
}

type VideoCommentListResp struct {
	StatusCode  int32       `json:"status_code"`
	StatusMsg   string      `json:"status_msg"`
	CommentList []CommentVo `json:"comment_list"`
}

type FollowActionResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type UserInFollowVo struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	IsFollow bool   `json:"is_follow"`
}

// FollowListResp 偶像列表粉丝列表公用该结构体
type FollowListResp struct {
	StatusCode int32            `json:"status_code"`
	StatusMsg  string           `json:"status_msg"`
	UserList   []UserInFollowVo `json:"user_list"`
}
