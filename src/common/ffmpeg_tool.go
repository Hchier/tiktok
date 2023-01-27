package common

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"os/exec"
	"strconv"
)

//@author by Hchier
//@Date 2023/1/27 15:50

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

// CompressVideo 压缩视频
func CompressVideo(videoPath string, frameRate int64, videoHeight int64, videoWidth int64, videoSavePath string) bool {
	cmdArguments := []string{
		"-i", videoPath,
		"-max_muxing_queue_size", "1024",
		"-r", strconv.FormatInt(frameRate, 10),
		"-b:v", "400k",
		"-crf", "25",
		"-s", strconv.FormatInt(videoHeight, 10) + "*" + strconv.FormatInt(videoWidth, 10),
		videoSavePath}

	cmd := exec.Command("ffmpeg", cmdArguments...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		ErrLog("压缩视频失败：", err.Error())
		return false
	}
	return true
}
