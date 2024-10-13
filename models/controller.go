package models

import (
	"Ranking/dao"
)

type Controller struct {
	ControllerId   int    `json:"id"`
	ControllerName string `json:"controller_name"`
	Password       string `json:"password"`
}

// 添加 TableName 方法，指定表名
func (Controller) TableName() string {
	return "controllers" // 指定表名为 controllers
}

// 判断用户名是否已经存在
func CheckControllerExist(controllername string) (Controller, error) {
	var controller Controller
	err := dao.Db.Where("controller_name =?", controllername).First(&controller).Error
	return controller, err
}

// 保存用户
func AddController(controllername, password string) (Controller, error) {
	controller := Controller{ControllerName: controllername, Password: password}
	err := dao.Db.Create(&controller).Error
	return controller, err
}
