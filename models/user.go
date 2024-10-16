package models

import (
	"Ranking/dao"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AddTime    int64  `json:"addtime"`
	UpdateTime int64  `json:"updatetime"`
}

type UserApi struct {
	Username string `json:"username"`
	Userid   int    `json:"userid"`
}

func (User) TableName() string {
	return "user"
}

// 判断用户名是否已经存在
func CheckUserExist(username string) (User, error) {
	var user User
	err := dao.Db.Where("username =?", username).First(&user).Error
	return user, err
}

// 保存用户
func AddUser(username, password string) (UserApi, error) {
	user := User{Username: username, Password: password, AddTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
	err := dao.Db.Create(&user).Error
	userapi := UserApi{Username: username, Userid: user.Id}
	return userapi, err
}

// 通过Id来查找用户
func CheckUserById(id int) (User, error) {
	var user User
	err := dao.Db.Where("id =?", id).First(&user).Error
	return user, err
}

// 删除通过用户名用户
func DeleteUserByUsername(username string) error {
	var user User
	err := dao.Db.Where("username =?", username).Delete(&user).Error
	return err
}

// 获取投票列表
func GetVoteList(userId int, sort string) ([]Vote, error) {
	var votes []Vote
	err := dao.Db.Where("user_id =?", userId).Order(sort).Find(&votes).Error
	return votes, err
}

// 更新密码
func UpdateUserPassword(username string, password string) (string, error) {
	var user User
	err := dao.Db.Model(&user).Where("username =?", username).UpdateColumn("password", password).Error
	return password, err
}
