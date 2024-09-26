package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (o OrderController) GetList(c *gin.Context) {
	cid := c.PostForm("cid")
	name := c.DefaultPostForm("name", "未获取输入") //第二个参数是设置为获取输入输出的默认值
	ReturnSuccess(c, 0, cid, name, 1)
}
