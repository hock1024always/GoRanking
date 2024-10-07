package controllers

import (
	"Ranking/models"
	"github.com/gin-gonic/gin"
)

// 实现关于用户的功能
type UserController struct{}

//type UserApi struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//	Userid   string `json:"userid"`
//}

func (u UserController) Register(c *gin.Context) {
	//接受用户名 密码以及确认密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirm_password", "")

	//验证 输入是否存在某项为空 密码和确认密码是否一致 是否已经存在该用户
	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "输入的用户名、密码、确认密码其中有空项")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4002, "密码和确认密码不一致")
		return
	}
	user, _ := models.CheckUserExist(username)
	if user.Id != 0 {
		ReturnError(c, 4003, "该用户已存在")
		return
	}

	//创建用户
	_, err := models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4004, "保存用户失败")
		return
	}

	ReturnSuccess(c, 0, "注册成功", user, 1)
}
