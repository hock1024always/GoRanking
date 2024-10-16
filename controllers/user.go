package controllers

import (
	"Ranking/config"
	"Ranking/models"
	"github.com/gin-gonic/gin"
)

// 实现关于用户的功能
type UserController struct{}
type UserLoginApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

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
	user1, _ := models.CheckUserExist(username)
	if user1.Id != 0 {
		ReturnError(c, 4003, "该用户已存在")
		return
	}

	//创建用户
	userapi, err2 := models.AddUser(username, password)
	if err2 != nil {
		ReturnError(c, 4004, "保存用户失败")
		return
	}

	ReturnSuccess(c, 0, "注册成功", userapi, 1)
}

// 登陆
func (u UserController) Login(c *gin.Context) {
	//接受用户名 密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if username == "" || password == "" {
		ReturnError(c, 4011, "用户名或密码为空")
		return
	}
	user1, err := models.CheckUserExist(username)
	if err != nil {
		ReturnError(c, 4012, "用户名不存在")
		return
	}
	if user1.Password != password {
		ReturnError(c, 4013, "密码错误")
		return
	}

	////使用sessions
	//session := sessions.Default(c)
	//session.Set("Login"+strconv.Itoa(user1.Id), user1.Id)
	//session.Save()

	//返回登录信息
	date := UserLoginApi{
		Id:       user1.Id,
		Username: user1.Username,
	}
	ReturnSuccess(c, 0, "登录成功", date, 1)
}

// 注销用户
func (u UserController) UserDelete(c *gin.Context) {
	//接受用户名 密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	sentence := c.DefaultPostForm("confirm_sentence", "")

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if username == "" || password == "" {
		ReturnError(c, 4011, "用户名或密码为空")
		return
	}
	user1, err := models.CheckUserExist(username)
	if err != nil {
		ReturnError(c, 4012, "用户名不存在")
		return
	}
	if user1.Password != password {
		ReturnError(c, 4013, "密码错误")
		return
	}
	if sentence != config.UserDeleteConfirmSentence {
		ReturnError(c, 4015, "确认语句错误")
		return
	}
	//删除用户
	err2 := models.DeleteUserByUsername(username)
	if err2 != nil {
		ReturnError(c, 4014, "删除用户失败")
		return
	}
	//返回删除信息
	ReturnSuccess(c, 0, "删除成功", nil, 1)
}

func (u UserController) GetVoteList(c *gin.Context) {
	//接受用户名 密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if username == "" || password == "" {
		ReturnError(c, 4011, "用户名或密码为空")
		return
	}
	user1, err := models.CheckUserExist(username)
	if err != nil {
		ReturnError(c, 4012, "用户名不存在")
		return
	}
	if user1.Password != password {
		ReturnError(c, 4013, "密码错误")
		return
	}

	//获取投票列表
	voteList, err2 := models.GetVoteList(user1.Id, "id desc")
	if err2 != nil {
		ReturnError(c, 4014, "获取投票列表失败")
		return
	}
	ReturnSuccess(c, 0, "获取投票列表成功", voteList, 1)
}

// 修改用户密码
func (u UserController) ModifyPassword(c *gin.Context) {
	//接受用户名 密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	newPassword := c.DefaultPostForm("new_password", "")
	confirmNewPassword := c.DefaultPostForm("confirm_new_password", "")

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if username == "" || password == "" {
		ReturnError(c, 4011, "用户名或密码为空")
		return
	}
	user1, err := models.CheckUserExist(username)
	if err != nil {
		ReturnError(c, 4012, "用户名不存在")
		return
	}
	if user1.Password != password {
		ReturnError(c, 4013, "密码错误")
		return
	}
	if newPassword != confirmNewPassword {
		ReturnError(c, 4015, "新密码与确认新密码不一致")
		return
	}

	//修改密码
	updatePassword, err2 := models.UpdateUserPassword(username, newPassword)
	if err2 != nil {
		ReturnError(c, 4014, "修改密码失败")
		return
	}
	//返回修改信息
	ReturnSuccess(c, 0, "修改密码成功", "新密码是:"+updatePassword, 1)
}
