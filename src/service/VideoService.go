package service

import (
	"github.com/cloudwego/hertz/pkg/app"
	"tiktok/src/common"
	"tiktok/src/mapper"
)

//@author by Hchier
//@Date 2023/1/21 14:15

// PublishVideo 发布视频（）
func PublishVideo(c *app.RequestContext, userId int64) *common.VideoPublishResp {
	videoData, _ := c.FormFile("data")
	videoData.Filename = common.GetSnowId() + ".mp4"
	playUrl := common.VideoDataDest + videoData.Filename
	videoAbsPath := common.StaticResourcePrefix + playUrl
	videoTitle := string(c.FormValue("title"))
	coverUrl := common.VideoCoverDest + common.GetSnowId() + ".png"
	coverAbsUrl := common.StaticResourcePrefix + coverUrl

	err := c.SaveUploadedFile(videoData, common.StaticResourcePrefix+playUrl)
	if err != nil {
		common.ErrLog("视频落盘出错：", err.Error())
		return &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "发布失败",
		}
	}

	if !common.CaptureVideoFrameAsPic(videoAbsPath, 1, 480, 270, coverAbsUrl) {
		return &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "发布失败",
		}
	}

	tx, err := common.Db.Beginx()
	if err != nil {
		common.ErrLog("发布视频事务开启失败：", err.Error())
		return &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "发布失败",
		}
	}
	if !mapper.InsertVideo(tx, userId, playUrl, coverUrl, videoTitle) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("发布视频事务回滚失败：", err.Error())
			return &common.VideoPublishResp{
				StatusCode: -1,
				StatusMsg:  "发布失败",
			}
		}
	}
	if !mapper.UpdateUserVideoCount(1, userId, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("发布视频事务回滚失败：", err.Error())
			return &common.VideoPublishResp{
				StatusCode: -1,
				StatusMsg:  "发布失败",
			}
		}
	}

	if tx.Commit() != nil {
		common.ErrLog("发布视频事务提交失败：", err.Error())
		return &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "发布失败",
		}
	}
	return &common.VideoPublishResp{
		StatusCode: 0,
		StatusMsg:  "发布成功",
	}
}

// TransferVideoEntityToVideoVo 将实体Video列表转为VideoVo列表
func TransferVideoEntityToVideoVo(videos []mapper.Video, currentUserId int64) []common.VideoVo {
	var videoVos []common.VideoVo = make([]common.VideoVo, len(videos))
	for i, video := range videos {
		var user mapper.User = mapper.SelectUserById(video.User_id)
		videoVos[i].Author.Id = user.Id
		videoVos[i].Author.Name = user.Username
		videoVos[i].Author.FollowCount = user.Follow_count
		videoVos[i].Author.FollowerCount = user.Follower_count
		videoVos[i].Author.Avatar = common.StaticResources + user.Avatar
		videoVos[i].Author.BackgroundImage = common.StaticResources + user.Background_image
		videoVos[i].Author.Signature = user.Signature
		videoVos[i].Author.TotalFavorited = user.Total_favorited
		videoVos[i].Author.FavoriteCount = user.Favorite_count
		videoVos[i].Author.VideoCount = user.Video_count
		videoVos[i].Author.IsFollow = mapper.ExistFollow(currentUserId, videoVos[i].Author.Id)

		videoVos[i].Id = video.Id
		videoVos[i].PlayUrl = common.StaticResources + video.Play_url
		videoVos[i].CoverUrl = common.StaticResources + video.Cover_url
		videoVos[i].FavoriteCount = video.Favorite_count
		videoVos[i].CommentCount = video.Comment_count
		videoVos[i].IsFavorite = mapper.ExistVideoFavor(currentUserId, video.Id)
		videoVos[i].Title = video.Title
	}
	return videoVos
}

// GetListOfPublishedVideo 拿到用户发布的视频列表
func GetListOfPublishedVideo(currentUserId, targetUserId int64) *common.ListOfPublishedVideoResp {
	if !mapper.ExistUserById(targetUserId) {
		return &common.ListOfPublishedVideoResp{
			StatusCode: -1,
			StatusMsg:  "用户不存在",
		}
	}
	valid, videos := mapper.GetPublishedVideoListByUserId(targetUserId)
	if !valid {
		return &common.ListOfPublishedVideoResp{
			StatusCode: -1,
			StatusMsg:  "查找视频失败",
		}
	}
	var resp common.ListOfPublishedVideoResp = common.ListOfPublishedVideoResp{
		StatusCode: 0,
		StatusMsg:  "查找视频成功",
	}
	resp.VideoList = TransferVideoEntityToVideoVo(videos, currentUserId)

	return &resp
}

func GetAuthorIdByVideoId(videoId int64) int64 {
	return mapper.SelectAuthorIdByVideoId(videoId)
}

// GetListOfFavoredVideo 拿到用户点赞的视频列表
func GetListOfFavoredVideo(currentUserId, targetUserId int64) *common.ListOfPublishedVideoResp {
	if !mapper.ExistUserById(targetUserId) {
		return &common.ListOfPublishedVideoResp{
			StatusCode: -1,
			StatusMsg:  "用户不存在",
		}
	}
	valid, videos := mapper.GetFavoredVideoListByUserId(targetUserId)
	if !valid {
		return &common.ListOfPublishedVideoResp{
			StatusCode: -1,
			StatusMsg:  "查找视频失败",
		}
	}
	var resp common.ListOfPublishedVideoResp = common.ListOfPublishedVideoResp{
		StatusCode: 0,
		StatusMsg:  "查找视频成功",
	}
	resp.VideoList = TransferVideoEntityToVideoVo(videos, currentUserId)

	return &resp
}

// VideoFeed 拿到视频流
func VideoFeed(userId int64) *common.ListOfPublishedVideoResp {
	valid, videos := mapper.GetVideoList()
	if !valid {
		return &common.ListOfPublishedVideoResp{
			StatusCode: -1,
			StatusMsg:  "查找视频失败",
		}
	}
	var resp common.ListOfPublishedVideoResp = common.ListOfPublishedVideoResp{
		StatusCode: 0,
		StatusMsg:  "查找视频成功",
	}
	resp.VideoList = TransferVideoEntityToVideoVo(videos, userId)
	return &resp
}
