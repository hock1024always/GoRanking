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

// 参赛者选择活动(每个参赛者只能选择一个)
func (p PlayerController) PlayerChooseActivity(c *gin.Context) {
	//接受用户名 密码
	nickname := c.DefaultPostForm("nickname", "")
	password := c.DefaultPostForm("password", "")
	activityIdStr := c.DefaultPostForm("activity_id", "0")
	activityId, _ := strconv.Atoi(activityIdStr)

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if nickname == "" || password == "" {
		ReturnError(c, 4121, "用户名或密码为空")
		return
	}
	user1, err := models.CheckPlayerExistsByNickname(nickname)
	if err != nil {
		ReturnError(c, 4122, "参赛者不存在")
		return
	}
	if user1.Password != password {
		ReturnError(c, 4123, "密码错误")
		return
	}
	if user1.Aid != 0 {
		ReturnError(c, 4124, "该参赛者正在参加活动"+activityIdStr)
		return
	}

	//验证活动是否存在 检验参赛者是否正在参与某项活动
	activity, err := models.GetActivityById(activityId)
	if err != nil {
		ReturnError(c, 4124, "活动不存在")
		return
	}

	//参加活动
	err = models.AddPlayerToActivityActivity(user1.Id, activity.Id)
	if err != nil {
		ReturnError(c, 4125, "参加活动失败")
		return
	}
	ReturnSuccess(c, 0, "参加活动成功", nil, 1)

}

// 参赛者设置宣言
func (p PlayerController) UpdateDeclaration(c *gin.Context) {
	//接受用户名 密码
	nickname := c.DefaultPostForm("nickname", "")
	password := c.DefaultPostForm("password", "")

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if nickname == "" || password == "" {
		ReturnError(c, 4131, "用户名或密码为空")
		return
	}
	user1, err := models.CheckPlayerExistsByNickname(nickname)
	if err != nil {
		ReturnError(c, 4132, "参赛者不存在")
		return
	}
	if user1.Password != password {
		ReturnError(c, 4133, "密码错误")
		return
	}

	//接受活动名称 要验证宣言字符串是否为空
	declaration := c.DefaultPostForm("declaration", "0")
	if declaration == "" {
		ReturnError(c, 4136, "宣言不能为空")
		return
	}
	//设置宣言
	player, err2 := models.AddDeclaration(user1.Id, declaration)
	if err2 != nil {
		ReturnError(c, 4135, "宣言更改失败")
		return
	}
	ReturnSuccess(c, 0, "宣言更改成功", "宣言更改为："+player.Declaration, 1)
}

// 参赛者取消参与某项活动
func (p PlayerController) QuitActivity(c *gin.Context) {
	//接受用户名 密码
	nickname := c.DefaultPostForm("nickname", "")
	password := c.DefaultPostForm("password", "")

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if nickname == "" || password == "" {
		ReturnError(c, 4131, "用户名或密码为空")
		return
	}
	user1, err := models.CheckPlayerExistsByNickname(nickname)
	if err != nil {
		ReturnError(c, 4132, "参赛者不存在")
		return
	}
	if user1.Password != password {
		ReturnError(c, 4133, "密码错误")
		return
	}

	//验证aid是否为0
	if user1.Aid == 0 {
		ReturnError(c, 4134, "该参赛者并未参与活动，无需删除")
		return
	}

	//删除参赛者参与的活动
	voteNum, err2 := models.DeletePlayerByActivityId(user1.Id, user1.Aid)
	if err2 != nil {
		//ReturnError(c, 4135, err2.Error())
		ReturnError(c, 4135, "删除参赛者参与的活动失败")
		return
	}
	//扣分
	err3 := models.ReduceScore(user1.Id, voteNum)
	if err3 != nil {
		ReturnError(c, 4136, "扣分失败")
		return
	}
	//将aid字段置为0
	err4 := models.UpdatePlayerAid(user1.Id, 0)
	if err4 != nil {
		ReturnError(c, 4137, "更新参赛者参与活动失败")
		return
	}
	ReturnSuccess(c, 0, "退出活动成功", "被扣除积分为："+strconv.Itoa(voteNum), 1)
}

// 查看给自己投票的用户
func (p PlayerController) GetVoteUsers(c *gin.Context) {
	//接受用户名 密码
	username := c.DefaultPostForm("nickname", "")
	password := c.DefaultPostForm("password", "")

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if username == "" || password == "" {
		ReturnError(c, 4011, "用户名或密码为空")
		return
	}
	user1, err := models.CheckPlayerExistsByNickname(username)
	if err != nil {
		ReturnError(c, 4012, "用户名不存在")
		return
	}
	if user1.Password != password {
		ReturnError(c, 4013, "密码错误")
		return
	}

	//获取投票列表
	voteList, err2 := models.GetVoteListForPlayer(user1.Id, "id desc")
	if err2 != nil {
		ReturnError(c, 4014, "获取投票列表失败")
		return
	}
	ReturnSuccess(c, 0, "获取投票列表成功", voteList, 1)
}

// 查看某个活动中给自己投票的用户
func (p PlayerController) GetVoteUsersInActivity(c *gin.Context) {
	//接受用户名 密码
	username := c.DefaultPostForm("nickname", "")
	password := c.DefaultPostForm("password", "")
	activityIdStr := c.DefaultPostForm("activity_id", "0")
	activityId, _ := strconv.Atoi(activityIdStr)

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if username == "" || password == "" {
		ReturnError(c, 4011, "用户名或密码为空")
		return
	}
	user1, err := models.CheckPlayerExistsByNickname(username)
	if err != nil {
		ReturnError(c, 4012, "用户名不存在")
		return
	}
	if user1.Password != password {
		ReturnError(c, 4013, "密码错误")
		return
	}
	if activityId == 0 {
		ReturnError(c, 4015, "活动不存在")
		return
	}

	//获取投票列表
	voteList, err2 := models.GetVoteListForPlayerByActivityId(user1.Id, activityId, "id desc")
	if err2 != nil {
		ReturnError(c, 4014, "获取投票列表失败")
		return
	}
	ReturnSuccess(c, 0, "获取投票列表成功", voteList, 1)
}
