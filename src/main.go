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
		common.StaticResourcePrefix+common.AvatarDest,
		common.StaticResourcePrefix+common.BackgroundImageDest,
		common.StaticResourcePrefix+common.VideoDataDest,
		common.StaticResourcePrefix+common.VideoCoverDest,
	)
}

func main() {
	Preparation()
	wg.Add(1)
	go common.TimeTask(time.Minute*2, service.RemoveExpiredToken)
	controller.Hertz.Spin()
	wg.Wait()
}
