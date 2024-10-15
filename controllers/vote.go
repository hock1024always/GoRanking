package controllers

import (
	"Ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type VoteController struct{}

func (v VoteController) AddVote(c *gin.Context) {
	//获取数据
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	playerIdStr := c.DefaultPostForm("player_id", "0") //这是对最终要转化成int的字符串的数据进行的防空操作
	activityIdStr := c.DefaultPostForm("activity_id", "0")

	playerId, _ := strconv.Atoi(playerIdStr)
	activityId, _ := strconv.Atoi(activityIdStr)
	//检查数输入
	if username == "" || playerId == 0 || activityId == 0 {
		ReturnError(c, 4031, "输入参数为空，请重新输入")
		return
	}
	//检查用户是否存在以及返回密码是否和数据库中的密码一致
	user, _ := models.CheckUserExist(username)
	if user.Id == 0 {
		ReturnError(c, 4032, "用户不存在")
		return
	}
	if user.Password != password {
		ReturnError(c, 4033, "密码错误")
		return
	}

	//判断作品是否存在 以及作品是否存在于该活动中
	player, _ := models.GetPlayerById(playerId)
	if player.Id == 0 {
		ReturnError(c, 4034, "参赛者不存在")
		return
	}
	if player.Aid != activityId {
		ReturnError(c, 4035, "参赛者选择的不是这个活动")
		return
	}

	//判断活动是否存在 以及是否在进行中
	activity, _ := models.GetActivityById(activityId)
	if activity.Id == 0 {
		ReturnError(c, 4036, "活动不存在")
		return
	}
	if activity.State != 1 {
		ReturnError(c, 4037, "活动不在投票时间")
		return
	}

	//检查是否已经vote了
	vote, _ := models.GetVoteInfo(user.Id, playerId, activityId)
	if vote.Id != 0 { //已经投过票了
		ReturnError(c, 4038, "已经投票过了")
		return
	}

	//添加vote
	rs, err := models.AddVote(user.Id, playerId, activityId)
	if err != nil {
		ReturnError(c, 4039, "投票失败")
		return
	}
	err2 := models.UpdateScoreByVote(playerId)
	if err2 != nil {
		ReturnError(c, 40310, "更新分数失败")
		return
	}
	ReturnSuccess(c, 0, "投票成功", rs, 1)

}

// 删除投票
func (v VoteController) DeleteVote(c *gin.Context) {
	//获取数据
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	playerIdStr := c.DefaultPostForm("player_id", "0") //这是对最终要转化成int的字符串的数据进行的防空操作
	activityIdStr := c.DefaultPostForm("activity_id", "0")

	playerId, _ := strconv.Atoi(playerIdStr)
	activityId, _ := strconv.Atoi(activityIdStr)
	//检查数输入
	if username == "" || playerId == 0 || activityId == 0 {
		ReturnError(c, 4031, "输入参数为空，请重新输入")
		return
	}
	//检查用户是否存在以及返回密码是否和数据库中的密码一致
	user, _ := models.CheckUserExist(username)
	if user.Id == 0 {
		ReturnError(c, 4032, "用户不存在")
		return
	}
	if user.Password != password {
		ReturnError(c, 4033, "密码错误")
		return
	}

	//判断活动是否存在 以及是否在进行中
	activity, _ := models.GetActivityById(activityId)
	if activity.Id == 0 {
		ReturnError(c, 4036, "活动不存在")
		return
	}
	if activity.State != 1 {
		ReturnError(c, 4037, "活动不在撤票时间")
		return
	}

	//检查是否存在vote
	vote, _ := models.GetVoteInfo(user.Id, playerId, activityId)
	if vote.Id == 0 {
		ReturnError(c, 4038, "没有投票记录")
		return
	}

	//删除vote
	err := models.DeleteVote(user.Id, playerId, activityId)
	if err != nil {
		ReturnError(c, 4039, "删除投票失败")
		return
	}
	err2 := models.UpdateScoreByActivity(playerId)
	if err2 != nil {
		ReturnError(c, 40310, "更新分数失败")
		return
	}

	ReturnSuccess(c, 0, "删除投票成功", nil, 1)
}
