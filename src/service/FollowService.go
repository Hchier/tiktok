package service

import (
	"tiktok/src/common"
	"tiktok/src/mapper"
)

//@author by Hchier
//@Date 2023/1/22 20:44

// DoFollow 关注。
// 3件事。1，插入关注记录。	2，粉丝偶像+1。	3，偶像粉丝+1
func DoFollow(follower, followee int64) *common.FollowActionResp {
	tx, err := mapper.Db.Beginx()
	if err != nil {
		common.ErrLog("关注操作时事务开启失败")
		return &common.FollowActionResp{StatusCode: -1, StatusMsg: "关注失败"}
	}
	if mapper.ExistFollow(follower, followee) {
		return &common.FollowActionResp{StatusCode: -1, StatusMsg: "已关注"}
	}

	//插入关注记录
	if !mapper.OperationFollow(1, follower, followee, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("关注操作时事务回滚失败", err.Error())
		}
		return &common.FollowActionResp{StatusCode: -1, StatusMsg: "关注失败"}
	}

	//粉丝偶像+1
	if !mapper.UpdateUserFollowCount(1, follower, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("关注操作时事务回滚失败", err.Error())
		}
		return &common.FollowActionResp{StatusCode: -1, StatusMsg: "关注失败"}
	}

	//偶像粉丝+1
	if !mapper.UpdateUserFollowerCount(1, followee, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("关注操作时事务回滚失败", err.Error())
		}
		return &common.FollowActionResp{StatusCode: -1, StatusMsg: "关注失败"}
	}

	err = tx.Commit()
	if err != nil {
		common.ErrLog("关注操作时事务提交失败", err.Error())
	}

	return &common.FollowActionResp{StatusCode: -0, StatusMsg: "关注成功"}
}

// DoUnFollow 取消关注。
// 3件事。1，插入关注记录。	2，粉丝偶像-1。	3，偶像粉丝-1
func DoUnFollow(follower, followee int64) *common.FollowActionResp {
	tx, err := mapper.Db.Beginx()
	if err != nil {
		common.ErrLog("取消关注时事务开启失败")
		return &common.FollowActionResp{StatusCode: -1, StatusMsg: "关取关注失败"}
	}
	if !mapper.ExistFollow(follower, followee) {
		return &common.FollowActionResp{StatusCode: -1, StatusMsg: "仍未关注"}
	}

	//删除关注记录
	if !mapper.OperationFollow(2, follower, followee, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("取关时事务回滚失败", err.Error())
		}
		return &common.FollowActionResp{StatusCode: -1, StatusMsg: "取关失败"}
	}

	//粉丝偶像-1
	if !mapper.UpdateUserFollowCount(2, follower, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("取关时事务回滚失败", err.Error())
		}
		return &common.FollowActionResp{StatusCode: -1, StatusMsg: "取关失败"}
	}

	//偶像粉丝-1
	if !mapper.UpdateUserFollowerCount(2, followee, tx) {
		err := tx.Rollback()
		if err != nil {
			common.ErrLog("取关时事务回滚失败", err.Error())
		}
		return &common.FollowActionResp{StatusCode: -1, StatusMsg: "取关失败"}
	}

	err = tx.Commit()
	if err != nil {
		common.ErrLog("取关时事务提交失败", err.Error())
	}

	return &common.FollowActionResp{StatusCode: -0, StatusMsg: "取关成功"}
}

func GetFolloweeInfo(followerId int64) *common.FollowListResp {
	success, ids := mapper.GetFolloweeIdsByFollowerId(followerId)
	if !success {
		return &common.FollowListResp{StatusCode: -1, StatusMsg: "获取信息失败"}
	}
	var resp = &common.FollowListResp{StatusCode: 0, StatusMsg: "成功"}
	resp.UserList = make([]common.UserInFollowVo, len(ids))
	for i, id := range ids {
		user := mapper.SelectUserById(id)
		resp.UserList[i].Id = id
		resp.UserList[i].Name = user.Username
		resp.UserList[i].Avatar = common.StaticResources + user.Avatar
		resp.UserList[i].IsFollow = true
	}
	return resp
}
