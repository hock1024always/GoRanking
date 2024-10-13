package controllers

import (
	"Ranking/config"
	"Ranking/models"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (u Controller) Register(c *gin.Context) {
	//接受用户名 密码以及确认密码
	username := c.DefaultPostForm("admin_name", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirm_password", "")
	key := c.DefaultPostForm("key", "")

	//验证 输入是否存在某项为空 密码和确认密码是否一致 是否已经存在该管理员
	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4201, "输入的管理员名、密码、确认密码其中有空项")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4202, "密码和确认密码不一致")
		return
	}
	user1, _ := models.CheckControllerExist(username)
	if user1.ControllerId != 0 {
		ReturnError(c, 4203, "该管理员已存在")
		return
	}
	if key != config.AdminKey {
		ReturnError(c, 4204, "密钥错误")
		return
	}

	//创建用户
	controller, err2 := models.AddController(username, password)
	if err2 != nil {
		ReturnError(c, 4205, "保存用户失败")
		//ReturnError(c, 4205, err2.Error())
		return
	}

	ReturnSuccess(c, 0, "注册成功", controller, 1)
}

func (u Controller) AddActivity(c *gin.Context) {
	//接受用户名 密码
	username := c.DefaultPostForm("admin_name", "")
	password := c.DefaultPostForm("password", "")

	//获得活动名称
	activityName := c.DefaultPostForm("activity_name", "")

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if username == "" || password == "" {
		ReturnError(c, 4211, "管理员名或密码为空")
		return
	}
	user1, err := models.CheckControllerExist(username)
	if err != nil {
		ReturnError(c, 4212, "管理员不存在")
		//ReturnError(c, 4212, err.Error())
		return
	}
	if user1.Password != password {
		ReturnError(c, 4213, "密码错误")
		return
	}

	//验证 活动名称是否已经存在
	activity, _ := models.CheckActivityExist(activityName)
	if activity.Id != 0 {
		ReturnError(c, 4214, "该活动已存在")
		return
	}

	//创建活动
	activity2, err2 := models.AddActivity(activityName)
	if err2 != nil {
		ReturnError(c, 4215, err2.Error())
		return
	}

	ReturnSuccess(c, 0, "注册成功", activity2, 1)

}
