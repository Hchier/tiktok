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
		common.Log(common.ErrLogDest, "视频落盘出错：", err.Error())
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
		common.Log("截图失败：", err.Error())
	}
	return buf
}

func GetListOfPublishedVideo(userId int64) *common.ListOfPublishedVideoResp {
	valid, videos := mapper.GetVideoListByUserId(userId)
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
	resp.VideoList = make([]common.VideoVo, 10)

	for i, video := range videos {
		resp.VideoList[i].Author.Id = 22
		resp.VideoList[i].Author.Name = "hchier"
		resp.VideoList[i].Author.Avatar = "http://192.168.0.105:8010/static/video/cover/1.png"
		resp.VideoList[i].Author.FollowCount = 2
		resp.VideoList[i].Author.FollowerCount = 22
		resp.VideoList[i].Author.IsFollow = true

		resp.VideoList[i].Id = video.Id
		resp.VideoList[i].PlayUrl = common.StaticResources + video.Play_url
		resp.VideoList[i].CoverUrl = common.StaticResources + video.Cover_url
		resp.VideoList[i].FavoriteCount = video.Favorite_count
		resp.VideoList[i].CommentCount = video.Comment_count
		resp.VideoList[i].Title = video.Title
	}
	return &resp
}
