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
		ReturnError(c, 4041, "输入的用户名、密码、确认密码其中有空项")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4042, "密码和确认密码不一致")
		return
	}
	user1, _ := models.CheckControllerExist(username)
	if user1.ControllerId != 0 {
		ReturnError(c, 4043, "该用户已存在")
		return
	}
	if key != config.AdminKey {
		ReturnError(c, 4045, "密钥错误")
		return
	}

	//创建用户
	controller, err2 := models.AddController(username, password)
	if err2 != nil {
		//ReturnError(c, 4044, "保存用户失败")
		ReturnError(c, 4044, err2.Error())
		return
	}

	ReturnSuccess(c, 0, "注册成功", controller, 1)
}
