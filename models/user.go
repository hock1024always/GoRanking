package models

import (
	"Ranking/dao"
	"fmt"
)

type User struct {
	Id       int
	Username string
}

func (User) TableName() string {
	return "user"
}

// GetUsersTest 根据用户ID获取用户信息
func GetUsersTest(id int) (User, error) {
	var user User
	//where在 SQL 中生成一个 WHERE 子句，以便查找满足条件的记录,?是占位符
	//first方法用于查找单个记录，如果找到，则返回该记录，否则返回错误
	err := dao.Db.Where("id =?", id).First(&user).Error
	return user, err
}

// 调用该方法，存储一个新用户 返回主键和错误信息（controllers包中调用）
func AddUser(username string) (int, error) {
	user := User{Username: username}
	err := dao.Db.Create(&user).Error
	if err != nil {
		return 0, fmt.Errorf("添加用户时出错：%w", err) // 返回详细错误
	}
	return user.Id, nil
}
