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

	rs, err := models.GetPlayers(aid)
	if err != nil {
		//ReturnError(c, 4021, err.Error()) //调试阶段打印错误信息到json返回中
		ReturnError(c, 4021, "获取参赛者列表失败")
		return
	}

	ReturnSuccess(c, 0, "获取成功", rs, 1)
}
