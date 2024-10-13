package controllers

import (
	"Ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PlayerController struct{}

func (p PlayerController) GetPlayerList(c *gin.Context) {
	//获取参赛者列表
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	rs, err := models.GetPlayers(aid, "id asc")
	if err != nil {
		//ReturnError(c, 4021, err.Error()) //调试阶段打印错误信息到json返回中
		ReturnError(c, 4121, "获取参赛者列表失败")
		return
	}

	ReturnSuccess(c, 0, "获取成功", rs, 1)
}

func (p PlayerController) GetRanking(c *gin.Context) {
	//获取活动编号
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	rs, err := models.GetPlayers(aid, "score desc")
	if err != nil {
		ReturnError(c, 4131, "获取排名失败")
		return
	}

	ReturnSuccess(c, 0, "获取成功", rs, 1)
	return
}

func (p PlayerController) PlayerRegister(c *gin.Context) {
	//接受用户名 密码以及确认密码
	nickname := c.DefaultPostForm("nickname", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirm_password", "")

	//验证 输入是否存在某项为空 密码和确认密码是否一致 是否已经存在该用户
	if nickname == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4111, "输入的参赛者名、密码、确认密码其中有空项")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4112, "密码和确认密码不一致")
		return
	}
	user1, _ := models.CheckPlayerExistsByNickname(nickname)
	if user1.Id != 0 {
		ReturnError(c, 4113, "该用户已存在")
		return
	}

	//创建用户
	player, err2 := models.AddPlayer(nickname, password)
	if err2 != nil {
		ReturnError(c, 4114, "保存用户失败")
		return
	}

	ReturnSuccess(c, 0, "注册成功", player, 1)
}

// 参赛者选择活动
func (p PlayerController) PlayerChooseActivity(c *gin.Context) {
	//接受用户名 密码
	nickname := c.DefaultPostForm("nickname", "")
	password := c.DefaultPostForm("password", "")

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if nickname == "" || password == "" {
		ReturnError(c, 4121, "用户名或密码为空")
		return
	}
	user1, err := models.CheckPlayerExistsByNickname(nickname)
	if err != nil {
		ReturnError(c, 41222, "参赛者不存在")
		return
	}
	if user1.Password != password {
		ReturnError(c, 4123, "密码错误")
		return
	}

	//接受活动名称
	activityName := c.DefaultPostForm("activity_name", "0")

	//验证活动是否存在
	activity, err := models.CheckActivityExist(activityName)
	if err != nil {
		ReturnError(c, 4124, "活动不存在")
	}
	//参加活动
	err = models.AddPlayerToActivityActivity(user1.Id, activity.Id)
	if err != nil {
		ReturnError(c, 4125, "参加活动失败")
	}
	ReturnSuccess(c, 0, "参加活动成功", nil, 1)

}

//参赛者设置宣言
