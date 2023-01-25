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

func main() {
	common.DelKeys("tokens", "expireTime")
	wg.Add(1)
	go common.TimeTask(time.Minute*2, service.RemoveExpiredToken)
	controller.Hertz.Spin()
	wg.Wait()
}
