package controllers

import (
	"Ranking/models"
	"github.com/gin-gonic/gin"
)

type ActivityController struct{}

// 获取活动列表
func (a ActivityController) GetActivityList(c *gin.Context) {

	rs, err := models.GetAllActivity("id asc")
	if err != nil {
		//ReturnError(c, 4131, "获取排名失败")
		ReturnError(c, 4131, err.Error())
		return
	}

	ReturnSuccess(c, 0, "获取成功", rs, 1)
	return
}

// 获取可以参加活动的活动列表
func (a ActivityController) GetActivityListForPlayer(c *gin.Context) {
	rs, err := models.GetAllActivityAvailable("id asc")
	if err != nil {
		ReturnError(c, 4131, "获取排名失败")
		return
	}

	ReturnSuccess(c, 0, "获取成功", rs, 1)
	return
}
