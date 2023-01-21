package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
	"tiktok/src/common"
	"tiktok/src/service"
)

//@author by Hchier
//@Date 2023/1/21 14:16

func VideoPublish(ctx context.Context, c *app.RequestContext) {

	token := string(c.FormValue("token"))
	res, _ := common.Rdb.HGet(ctx, "tokens", token).Result()
	userId, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, &common.VideoPublishResp{
			StatusCode: -1,
			StatusMsg:  "身份验证失败",
		})
	}

	c.JSON(http.StatusOK, service.PublishVideo(c, userId))
}
