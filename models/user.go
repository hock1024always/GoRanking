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
func AddUser(username, password string) (int, error) {
	user := User{Username: username, Password: password, AddTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}
