package common

//@author by Hchier
//@Date 2023/1/25 10:10

// 以下5个const在部署时要更改
const RedisAddr = "127.0.0.1:6379" //
const RedisPassword = ""           //
const ErrLogDest = "E:\\hchier\\GoProjects\\tiktok\\logs/err.log"
const StaticResourcePrefix = "E:\\static\\tiktok\\"
const StaticResources = "http://192.168.0.105:8010/tiktok/"

const AvatarDest = "avatar\\"
const BackgroundImageDest = "bg\\"
const VideoDataDest = "video\\data\\"
const VideoCoverDest = "video\\cover\\"

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
