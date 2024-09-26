package controllers

import (
	"Ranking/models"
	"Ranking/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 实现关于用户的功能
type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	name := c.Param("name")

	id, _ := strconv.Atoi(idStr)
	user, _ := models.GetUsersTest(id)

	ReturnSuccess(c, 0, name, user, 1)

}

func (u UserController) GetList(c *gin.Context) {

	// 程序员手动设置的日志
	logger.Write("日志信息", "user")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常：", err)
		}
	}()

	num1, num2 := 1, 0
	num3 := num1 / num2
	ReturnUserGetListError(c, 404, num3)
}

func (u UserController) AddUser(c *gin.Context) {
	//logger.Write("日志信息", "user")
	username := c.DefaultPostForm("username", "")

	// 输入检查
	if username == "" {
		ReturnError(c, 400, "用户名不能为空")
		return
	}

	id, err := models.AddUser(username)
	if err != nil {
		ReturnError(c, 400, err.Error()) // 返回具体错误信息
		return
	}
	ReturnSuccess(c, 0, "用户添加成功", id, 1)
}
