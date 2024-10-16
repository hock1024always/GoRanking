package controllers

import (
	"Ranking/config"
	"Ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
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
		ReturnError(c, 4215, "保存活动失败")
		return
	}

	ReturnSuccess(c, 0, "注册成功", activity2, 1)
}

func (u Controller) UpdatePlayersScore(c *gin.Context) {
	//接受用户名 密码
	username := c.DefaultPostForm("admin_name", "")
	password := c.DefaultPostForm("password", "")

	//获得用户名称
	playerName := c.DefaultPostForm("player_name", "")
	updateScoreStr := c.DefaultPostForm("update_score", "")
	updateScore, _ := strconv.Atoi(updateScoreStr)

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if username == "" || password == "" {
		ReturnError(c, 4311, "管理员名或密码为空")
		return
	}
	user1, err := models.CheckControllerExist(username)
	if err != nil {
		ReturnError(c, 4312, "管理员不存在")
		//ReturnError(c, 4212, err.Error())
		return
	}
	if user1.Password != password {
		ReturnError(c, 4313, "密码错误")
		return
	}

	//验证 参赛者是否存在
	player, _ := models.CheckPlayerExistsByNickname(playerName)
	if player.Id == 0 {
		ReturnError(c, 4314, "该参赛者不存在")
		return
	}

	//更新参赛者的得分
	err3 := models.UpdateScoreByAdmin(playerName, updateScore)
	if err3 != nil {
		ReturnError(c, 4315, "更新失败")
		return
	}
	ReturnSuccess(c, 0, "更新成功", player.Nickname+"更新之后的得分为:"+updateScoreStr, 1)
}

// 获取所有参赛者的得分
func (u Controller) GetAllRanking(c *gin.Context) {
	//接受用户名 密码
	username := c.DefaultPostForm("admin_name", "")
	password := c.DefaultPostForm("password", "")

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if username == "" || password == "" {
		ReturnError(c, 4311, "管理员名或密码为空")
		return
	}
	user1, err := models.CheckControllerExist(username)
	if err != nil {
		ReturnError(c, 4312, "管理员不存在")
		//ReturnError(c, 4212, err.Error())
		return
	}
	if user1.Password != password {
		ReturnError(c, 4313, "密码错误")
		return
	}

	rs, err2 := models.GetAllPlayers("score desc")
	if err2 != nil {
		ReturnError(c, 4131, "获取排名失败")
		return
	}

	ReturnSuccess(c, 0, "获取成功", rs, 1)

}

// 关闭某个活动
func (u Controller) CloseActivity(c *gin.Context) {
	//接受用户名 密码
	username := c.DefaultPostForm("admin_name", "")
	password := c.DefaultPostForm("password", "")
	aidStr := c.DefaultPostForm("activity_id", "0")
	aid, _ := strconv.Atoi(aidStr)

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if username == "" || password == "" {
		ReturnError(c, 4311, "管理员名或密码为空")
		return
	}
	user1, err := models.CheckControllerExist(username)
	if err != nil {
		ReturnError(c, 4312, "管理员不存在")
		//ReturnError(c, 4212, err.Error())
		return
	}
	if user1.Password != password {
		ReturnError(c, 4313, "密码错误")
		return
	}

	//验证 活动是否存在
	activity, _ := models.CheckActivityExistById(aid)
	if activity.Id == 0 {
		ReturnError(c, 4314, "该活动不存在")
		return
	}
	if activity.State == 0 {
		ReturnError(c, 4315, "该活动已关闭")
		return
	}
	//关闭活动
	err3 := models.CloseActivity(aid, 0)
	if err3 != nil {
		ReturnError(c, 4315, "关闭活动失败")
		return
	}
	ReturnSuccess(c, 0, "关闭成功", activity, 1)
}

func (u Controller) DeletePlayerScore(c *gin.Context) {
	//接受用户名 密码
	username := c.DefaultPostForm("admin_name", "")
	password := c.DefaultPostForm("password", "")
	aidStr := c.DefaultPostForm("activity_id", "0")
	playerIdStr := c.DefaultPostForm("player_id", "0")
	aid, _ := strconv.Atoi(aidStr)
	playerId, _ := strconv.Atoi(playerIdStr)

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if username == "" || password == "" {
		ReturnError(c, 4311, "管理员名或密码为空")
		return
	}
	user1, err := models.CheckControllerExist(username)
	if err != nil {
		ReturnError(c, 4312, "管理员不存在")
		//ReturnError(c, 4212, err.Error())
		return
	}
	if user1.Password != password {
		ReturnError(c, 4313, "密码错误")
		return
	}

	//验证 玩家是否存在
	player, _ := models.GetPlayerById(playerId)
	if player.Id == 0 {
		ReturnError(c, 4314, "该玩家不存在")
		return
	}
	//验证 活动是否存在
	activity, _ := models.CheckActivityExistById(aid)
	if activity.Id == 0 {
		ReturnError(c, 4314, "该活动不存在")
		return
	}

	//删除某个参赛者的得分
	voteNum, err2 := models.DeletePlayerByActivityId(player.Id, player.Aid)
	if err2 != nil {
		ReturnError(c, 4315, "删除失败")
		return
	}
	//扣分
	err3 := models.ReduceScore(player.Id, voteNum)
	if err3 != nil {
		ReturnError(c, 4136, "扣分失败")
		return
	}
	ReturnSuccess(c, 0, "删除成功", "删除了"+player.Nickname+"在"+activity.Name+"活动的"+strconv.Itoa(voteNum)+"票", 1)
}

// 删除某个用户对于某个活动的投票
func (u Controller) DeleteVote(c *gin.Context) {
	//接受用户名 密码
	username := c.DefaultPostForm("admin_name", "")
	password := c.DefaultPostForm("password", "")
	aidStr := c.DefaultPostForm("activity_id", "0")
	userIdStr := c.DefaultPostForm("user_id", "0")
	aid, _ := strconv.Atoi(aidStr)
	userId, _ := strconv.Atoi(userIdStr)

	//验证 用户名或者密码为空 用户名不存在 密码错误
	if username == "" || password == "" {
		ReturnError(c, 4311, "管理员名或密码为空")
		return
	}
	user1, err := models.CheckControllerExist(username)
	if err != nil {
		ReturnError(c, 4312, "管理员不存在")
		//ReturnError(c, 4212, err.Error())
		return
	}
	if user1.Password != password {
		ReturnError(c, 4313, "密码错误")
		return
	}

	//验证 玩家是否存在
	user, _ := models.CheckUserById(userId)
	if user.Id == 0 {
		ReturnError(c, 4314, "该用户不存在")
		return
	}
	//验证 活动是否存在
	activity, _ := models.CheckActivityExistById(aid)
	if activity.Id == 0 {
		ReturnError(c, 4314, "该活动不存在")
		return
	}

	playerIds, err2 := models.DeleteVoteByUserIdAndActivityId(userId, aid)
	if err2 != nil {
		ReturnError(c, 4315, "删除失败")
		return
	}
	if len(playerIds) == 0 {
		ReturnError(c, 4315, "该用户没有投票")
		return
	}
	//扣分
	err3 := models.DeleteVoteScore(playerIds)
	if err3 != nil {
		ReturnError(c, 4136, "扣分失败")
		return
	}
	//将整形数组转化成字符串类型数组
	playerIdStrs := models.IntArrayToStringArray(playerIds)
	ReturnSuccess(c, 0, "删除成功", "扣分的参赛者有:"+strings.Join(playerIdStrs, ","), 1)

}
