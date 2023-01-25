package common

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//@author by Hchier
//@Date 2023/1/20 23:29

// GetRandStr 得到长度为len的随机字符串
func GetRandStr(len int8) string {
	result := make([]byte, len/2)
	rand.Read(result)
	return hex.EncodeToString(result)
}

// IsValidUser 验证身份是否有效
// 已登录：返回当前登录的用户的id
// 未登录：返回-1
func IsValidUser(ctx context.Context, c *app.RequestContext) {
	res, _ := Rdb.HGet(ctx, "tokens", string(c.FormValue("token"))).Result()
	userId, err := strconv.ParseInt(res, 10, 64)
	// 未登录
	if err != nil {
		c.Set("id", int64(-1))
		return
	}
	c.Set("id", userId)
}

// CaptureVideoFrameAsPic 截取视频帧作为图片保存
// 成功返回true
func CaptureVideoFrameAsPic(videoPath string, frameNum int16, picHeight int, picWidth int, picPath string) bool {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "png"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		ErrLog("截图失败：", err.Error())
		return false
	}
	img, err := imaging.Decode(buf)

	err = imaging.Save(imaging.Resize(img, picWidth, picHeight, imaging.Lanczos), picPath)
	if err != nil {
		ErrLog("图片落盘失败：", err.Error())
		return false
	}
	return true
}

// TimeTask 定时任务
func TimeTask(d time.Duration, task func()) {
	ticker := time.NewTicker(d)
	for range ticker.C {
		task()
	}
}

// Md5Encode md5加密
func Md5Encode(str string) string {
	srcCode := md5.Sum([]byte(str))
	// md5.Sum函数加密后返回的是字节数组，需要转换成16进制形式
	code := fmt.Sprintf("%x", srcCode)
	return code
}
