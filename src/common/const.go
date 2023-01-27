package common

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

//@author by Hchier
//@Date 2023/1/25 10:10

//const RedisAddr = "127.0.0.1:6379" //
//const RedisPassword = ""           //
//const ErrLogPath = "E:\\hchier\\GoProjects\\tiktok\\logs/err.log"
//const StaticResourcePathPrefix = "E:\\static\\tiktok\\"
//const StaticResourceUrlPrefix = "http://192.168.0.105:8010/tiktok/"
//const DataSourceName = "root:pyh903903@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
//
//const DriverName = "mysql"
//const TokenValidity = 30      //token有效期（分钟）
//const CheckTokenDuration = 10 //定期检查token是否有效（分钟）
//const AvatarPathPrefix = "avatar\\"
//const BackgroundImagePathPrefix = "bg\\"
//const VideoDataPathPrefix = "video\\data\\"
//const VideoCoverPathPrefix = "video\\cover\\"

var (
	HostPorts                 = ""
	RedisAddr                 = ""
	RedisPassword             = ""
	ErrLogPath                = ""
	StaticResourcePathPrefix  = ""
	StaticResourceUrlPrefix   = ""
	DataSourceName            = ""
	DriverName                = ""
	TokenValidity             = int64(0)
	CheckTokenDuration        = int64(0)
	AvatarPathPrefix          = ""
	BackgroundImagePathPrefix = ""
	VideoDataPathPrefix       = ""
	VideoDataTempPathPrefix   = ""
	VideoCoverPathPrefix      = ""
	FrameRate                 = int64(0)
	VideoHeight               = int64(0)
	VideoWidth                = int64(0)
	PicHeight                 = int64(0)
	PicWidth                  = int64(0)
)

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

func LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	var configs map[string]string = make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		splits := strings.SplitN(line, "=", 2)
		configs[splits[0]] = splits[1]
	}

	HostPorts = configs["HostPorts"]
	RedisAddr = configs["RedisAddr"]
	RedisPassword = configs["RedisPassword"]
	ErrLogPath = configs["ErrLogPath"]
	StaticResourcePathPrefix = configs["StaticResourcePathPrefix"]
	StaticResourceUrlPrefix = configs["StaticResourceUrlPrefix"]
	DataSourceName = configs["DataSourceName"]
	DriverName = configs["DriverName"]
	TokenValidity, _ = strconv.ParseInt(configs["TokenValidity"], 10, 64)
	CheckTokenDuration, _ = strconv.ParseInt(configs["CheckTokenDuration"], 10, 64)
	AvatarPathPrefix = configs["AvatarPathPrefix"]
	BackgroundImagePathPrefix = configs["BackgroundImagePathPrefix"]
	VideoDataPathPrefix = configs["VideoDataPathPrefix"]
	VideoDataTempPathPrefix = configs["VideoDataTempPathPrefix"]
	VideoCoverPathPrefix = configs["VideoCoverPathPrefix"]
	FrameRate, _ = strconv.ParseInt(configs["FrameRate"], 10, 64)
	VideoHeight, _ = strconv.ParseInt(configs["VideoHeight"], 10, 64)
	VideoWidth, _ = strconv.ParseInt(configs["VideoWidth"], 10, 64)
	PicHeight, _ = strconv.ParseInt(configs["PicHeight"], 10, 64)
	PicWidth, _ = strconv.ParseInt(configs["PicWidth"], 10, 64)
}
