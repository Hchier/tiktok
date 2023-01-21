package service

import (
	"bytes"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
	"tiktok/src/common"
	"tiktok/src/mapper"
)

//@author by Hchier
//@Date 2023/1/21 14:15

func PublishVideo(c *app.RequestContext, userId int64) *common.VideoPublishResp {

	video_data, _ := c.FormFile("data")
	video_data.Filename = common.GetSnowId() + ".mp4"
	video_path := common.VideoDataDest + video_data.Filename
	video_title := string(c.FormValue("title"))
	cover_url := common.VideoCoverDest + common.GetSnowId() + ".jpeg"

	err := c.SaveUploadedFile(video_data, video_path)
	if err != nil {
		common.ErrLog("视频落盘出错：", err.Error())
		return &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "视频落盘出错",
		}
	}

	//buf := GetJpegFrame(video_path, 1)
	//img, err := imaging.Decode(buf)
	//err = imaging.Save(img, cover_url)
	//if err != nil {
	//	common.Log("落盘失败：", err.Error())
	//	return &common.VideoPublishResp{
	//		StatusCode: -1,
	//		StatusMsg:  "落盘失败",
	//	}
	//}

	mapper.InsertVideo(userId, video_path, cover_url, video_title)
	return &common.VideoPublishResp{
		StatusCode: 0,
		StatusMsg:  "发布成功",
	}
}

func GetJpegFrame(videoPath string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		common.ErrLog("截图失败：", err.Error())
	}
	return buf
}

func TransferVideoEntityToVideoVo(videos []mapper.Video, userId int64) []common.VideoVo {
	var videoVos []common.VideoVo = make([]common.VideoVo, len(videos))
	for i, video := range videos {
		videoVos[i].Author.Id = 22
		videoVos[i].Author.Name = "hchier"
		videoVos[i].Author.Avatar = "http://192.168.0.105:8010/static/video/cover/1.png"
		videoVos[i].Author.Signature = "000000000000000"
		videoVos[i].Author.FollowCount = 2
		videoVos[i].Author.FollowerCount = 22
		videoVos[i].Author.IsFollow = true

		videoVos[i].Id = video.Id
		videoVos[i].PlayUrl = common.StaticResources + video.Play_url
		videoVos[i].CoverUrl = common.StaticResources + video.Cover_url
		videoVos[i].FavoriteCount = video.Favorite_count
		videoVos[i].CommentCount = video.Comment_count
		videoVos[i].IsFavorite = mapper.CheckVideoFavorByUserIdAndVideoId(userId, video.Id)
		videoVos[i].Title = video.Title
	}
	return videoVos
}

func GetListOfPublishedVideo(userId int64) *common.ListOfPublishedVideoResp {
	valid, videos := mapper.GetPublishedVideoListByUserId(userId)
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

func GetAuthorIdByVideoId(videoId int64) int64 {
	return mapper.SelectAuthorIdByVideoId(videoId)
}
