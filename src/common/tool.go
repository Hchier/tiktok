package common

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/disintegration/imaging"
	"github.com/go-redis/redis/v8"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

//@author by Hchier
//@Date 2023/1/20 23:29

const ErrLogDest = "logs/err.log"
const VideoDataDest = "static\\video\\data\\"
const VideoCoverDest = "static\\video\\cover\\"
const StaticResources = "http://192.168.0.105:8010/"

var Signatures = [...]string{
	"雄关漫道真如铁，而今迈步从头越。",
	"待到山花烂漫时，她在丛中笑。",
	"天若有情天亦老,人间正道是沧桑。",
	"天高云淡，望断南飞雁。不到长城非好汉，",
	"萧瑟秋风今又是，换了人间。",
	"神女应无恙，当惊世界殊。",
	"不管风吹浪打,胜似闲庭信步,今日得宽馀。",
	"为有牺牲多壮志，敢教日月换新天。",
	"四海翻腾云水怒，五洲震荡风雷激。",
	"世上无难事，只要肯登攀。",
}

func GetFile(dest string) *os.File {
	f, err := os.OpenFile(dest, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	return f
}

var Rdb = GetRdb()

func GetRdb() *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// GetRandStr 得到长度为len的随机字符串
func GetRandStr(len int8) string {
	result := make([]byte, len/2)
	rand.Read(result)
	return hex.EncodeToString(result)
}

func Log(dest string, v ...interface{}) {
	file := GetFile(dest)
	defer file.Close()
	hlog.SetOutput(file)
	hlog.Error(v...)
}

// ErrLog 打印错误日志
func ErrLog(v ...interface{}) {
	Log(ErrLogDest, v)
}

// 雪花算法
const (
	workerBits  uint8 = 10                      //机器码位数
	numberBits  uint8 = 12                      //序列号位数
	workerMax   int64 = -1 ^ (-1 << workerBits) //机器码最大值（即1023）
	numberMax   int64 = -1 ^ (-1 << numberBits) //序列号最大值（即4095）
	timeShift   uint8 = workerBits + numberBits //时间戳偏移量
	workerShift uint8 = numberBits              //机器码偏移量
	epoch       int64 = 1656856144640           //起始常量时间戳（毫秒）,此处选取的时间是2022-07-03 21:49:04
)

type Worker struct {
	mu        sync.Mutex
	timeStamp int64
	workerId  int64
	number    int64
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("WorkerId超过了限制！")
	}
	return &Worker{
		timeStamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) NextId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	//当前时间的毫秒时间戳
	now := time.Now().UnixNano() / 1e6
	//如果时间戳与当前时间相同，则增加序列号
	if w.timeStamp == now {
		w.number++
		//如果序列号超过了最大值，则更新时间戳
		if w.number > numberMax {
			for now <= w.timeStamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else { //如果时间戳与当前时间不同，则直接更新时间戳
		w.number = 0
		w.timeStamp = now
	}
	//ID由时间戳、机器编码、序列号组成
	ID := (now-epoch)<<timeShift | (w.workerId << workerShift) | (w.number)
	return ID
}

func GetSnowId() string {
	worker, err := NewWorker(0)
	if err != nil {
		fmt.Println(err)
	}
	id := worker.NextId()
	return strconv.FormatInt(id, 10)
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

func TimeTask(d time.Duration, task func()) {
	ticker := time.NewTicker(d)
	for range ticker.C {
		task()
	}
}
