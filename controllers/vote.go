package controllers

import (
	"Ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type VoteController struct{}

func (v VoteController) AddVote(c *gin.Context) {
	//获取数据
	userIdStr := c.DefaultPostForm("user_id", "0")
	playerIdStr := c.DefaultPostForm("player_id", "0")
	userId, _ := strconv.Atoi(userIdStr)
	playerId, _ := strconv.Atoi(playerIdStr)

	//检查数输入
	if userId == 0 || playerId == 0 {
		ReturnError(c, 4031, "输入参数为空，请重新输入")
	}
	user, _ := models.CheckUserById(userId)
	if user.Id == 0 {
		ReturnError(c, 4032, "用户不存在")
	}
	player, _ := models.GetPlayerById(playerId)
	if player.Id == 0 {
		ReturnError(c, 4033, "参赛者不存在")
	}

	//检查是否已经vote了 返回值重复投票
	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.Id != 0 { //已经投过票了
		ReturnError(c, 4034, "已经投票过了")
	}

	//添加vote
	rs, err := models.AddVote(userId, playerId)
	if err == nil {
		models.UpdateScoreByVote(playerId)
		ReturnSuccess(c, 0, "投票成功", rs, 1)
		return
	}

	ReturnError(c, 4035, err.Error())
}
