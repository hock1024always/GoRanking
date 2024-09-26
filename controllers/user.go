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
		ReturnError(c, 400, "用户添加失败"+err.Error()) // 返回具体错误信息
		return
	}
	ReturnSuccess(c, 0, "用户添加成功", id, 1)
}

// 更新用户名
func (u UserController) UpdateUser(c *gin.Context) {
	//获取用户信息
	username := c.DefaultPostForm("username", "")
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)

	//调用方法更新数据库中的用户名
	models.UpdateUser(id, username)
	ReturnSuccess(c, 0, "用户更新成功", true, 1)
}

// 删除用户
func (u UserController) DeleteUser(c *gin.Context) {
	//获取id
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)

	//调用方法删除数据库中的用户
	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 404, "用户删除失败"+err.Error())
	}
	ReturnSuccess(c, 0, "用户删除成功", true, 1)

}

//func (u UserController) GetAllUsers(c *gin.Context) {
//	users, err := models.GetAllUsers()
//	if err != nil {
//		ReturnError(c, 404, "用户列表获取失败"+err.Error())
//	}
//	ReturnSuccess(c, 0, "用户列表获取成功成功", users, 1)
//}

func (u UserController) GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		ReturnError(c, 404, "用户列表获取失败: "+err.Error())
		return // 添加 return 结束函数的执行
	}

	// 处理成功的情况，避免重复的“成功”字样
	ReturnSuccess(c, 0, "用户列表获取成功", users, 1)
}
