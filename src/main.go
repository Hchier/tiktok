package main

import (
	"sync"
	"tiktok/src/common"
	"tiktok/src/controller"
	"tiktok/src/service"
	"time"
)

// @author by Hchier
// @Date 2023/1/20 22:50

var wg sync.WaitGroup

func Preparation() {

	common.DelKeys("tokens", "expireTime")
	//ErrLogDest不用创建，不存在时会自动创建
	common.MakeDirs(
		common.StaticResourcePathPrefix+common.AvatarPathPrefix,
		common.StaticResourcePathPrefix+common.BackgroundImagePathPrefix,
		common.StaticResourcePathPrefix+common.VideoDataPathPrefix,
		common.StaticResourcePathPrefix+common.VideoDataTempPathPrefix,
		common.StaticResourcePathPrefix+common.VideoCoverPathPrefix,
	)

}

func main() {
	Preparation()
	wg.Add(1)
	go common.TimeTask(time.Duration(time.Minute.Nanoseconds()*common.CheckTokenDuration), service.RemoveExpiredToken)
	controller.Hertz.Spin()
	wg.Wait()
}
