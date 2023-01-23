package service

import (
	"tiktok/src/common"
	"tiktok/src/mapper"
	"time"
)

//@author by Hchier
//@Date 2023/1/22 13:31

// DoPublishVideoComment 发布视频评论
// 1，插入评论		2，更新视频评论数
func DoPublishVideoComment(videoId int64, content string, userId int64, commentId int64) *common.CommentActionResp {
	tx, err := mapper.Db.Beginx()
	if err != nil {
		common.ErrLog("发布视频评论时事务开启失败：", err.Error())
		return &common.CommentActionResp{
			StatusCode: -1,
			StatusMsg:  "发布视频评论时事务开启失败",
		}
	}

	//插入评论
	hasInserted, commentId := mapper.OperateVideoComment(1, videoId, content, userId, commentId, tx)
	if !hasInserted {
		_ = tx.Rollback()
		return &common.CommentActionResp{
			StatusCode: -1,
			StatusMsg:  "评论失败",
		}
	}

	//更新视频评论数
	if !mapper.UpdateVideoCommentCount(1, videoId, tx) {
		_ = tx.Rollback()
		return &common.CommentActionResp{
			StatusCode: -1,
			StatusMsg:  "评论失败",
		}
	}

	err = tx.Commit()
	if err != nil {
		common.ErrLog("插入评论时事务提交失败")
		return &common.CommentActionResp{
			StatusCode: -1,
			StatusMsg:  "评论失败",
		}
	}

	user := mapper.SelectUserById(userId)
	return &common.CommentActionResp{
		StatusCode: 0,
		StatusMsg:  "评论成功",
		Comment: common.CommentVo{
			Id: commentId,
			User: struct {
				Id            int64  `json:"id"`
				Name          string `json:"name"`
				Avatar        string `json:"avatar"`
				FollowCount   int64  `json:"follow_count"`
				FollowerCount int64  `json:"follower_count"`
				IsFollow      bool   `json:"is_follow"`
			}{
				Id:       user.Id,
				Name:     user.Username,
				Avatar:   common.StaticResources + user.Avatar,
				IsFollow: true,
			},
			Content:    content,
			CreateDate: time.Now().Format("01-02"),
		},
	}
}

// DoDeleteVideoComment 删除视频评论
// 1，删除评论		2，更新视频评论数
func DoDeleteVideoComment(videoId int64, content string, userId int64, commentId int64) *common.CommentActionResp {
	tx, err := mapper.Db.Beginx()
	if err != nil {
		common.ErrLog("删除视频评论时事务开启失败：", err.Error())
		return &common.CommentActionResp{
			StatusCode: -1,
			StatusMsg:  "删除视频评论时事务开启失败",
		}
	}

	//删除评论
	hasInserted, commentId := mapper.OperateVideoComment(2, videoId, content, userId, commentId, tx)
	if !hasInserted {
		_ = tx.Rollback()
		return &common.CommentActionResp{
			StatusCode: -1,
			StatusMsg:  "删除评论失败",
		}
	}

	//更新视频评论数
	if !mapper.UpdateVideoCommentCount(2, videoId, tx) {
		_ = tx.Rollback()
		return &common.CommentActionResp{
			StatusCode: -1,
			StatusMsg:  "评论失删除评论失败败",
		}
	}

	err = tx.Commit()
	if err != nil {
		common.ErrLog("删除视频评论时事务提交失败")
		return &common.CommentActionResp{
			StatusCode: -1,
			StatusMsg:  "删除评论失败",
		}
	}

	return &common.CommentActionResp{
		StatusCode: 0,
		StatusMsg:  "删除成功",
	}
}

func GetVideoCommentByVideoId(videoId int64) *common.VideoCommentListResp {
	success, videoCommentList := mapper.GetVideoCommentByVideoId(videoId)
	if !success {
		return &common.VideoCommentListResp{
			StatusCode: -1,
			StatusMsg:  "视频评论查找失败",
		}
	}

	resp := &common.VideoCommentListResp{
		StatusCode: 0,
		StatusMsg:  "ok",
	}
	resp.CommentList = make([]common.CommentVo, len(videoCommentList))
	for i, comment := range videoCommentList {
		resp.CommentList[i].Id = comment.Id

		var user mapper.User = mapper.SelectUserById(comment.User_id)
		resp.CommentList[i].User.Id = user.Id
		resp.CommentList[i].User.Name = user.Username
		resp.CommentList[i].User.Avatar = common.StaticResources + user.Avatar
		resp.CommentList[i].User.FollowCount = user.Follow_count
		resp.CommentList[i].User.FollowerCount = user.Follower_count
		resp.CommentList[i].User.IsFollow = true

		resp.CommentList[i].Content = comment.Content
		resp.CommentList[i].CreateDate = comment.Create_date.Format("01-02")
	}
	return resp
}
