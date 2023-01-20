package service

import "tiktok/src/mapper"

//@author by Hchier
//@Date 2023/1/20 21:03

// Register
// 用户注册。首先检查用户名是否被占用。
func Register(username, password string) (int32, string) {
	if mapper.ExistUser(username) {
		return -2, "username已被使用"
	}
	return mapper.InsertUser(username, password)
}
