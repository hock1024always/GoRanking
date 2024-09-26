package controllers

import "github.com/gin-gonic/gin"

type GoodUserController struct{}

type Search struct {
	Cid  int    `json:"cid"` //前端传过来的小写
	Name string `json:"name"`
}

func (p GoodUserController) GetList(c *gin.Context) {

	search := Search{}
	err := c.BindJSON(&search)
	if err != nil {
		ReturnError(c, 400, err.Error()) // 返回错误，状态码可以选择 400（Bad Request）
		return
	}
	ReturnSuccess(c, 0, search.Cid, search.Name, 1) // 仅在没有错误时返回成功
}

func (p GoodUserController) GetList2(c *gin.Context) {
	param := make(map[string]interface{})
	err := c.BindJSON(&param)
	if err != nil {
		ReturnError(c, 400, err.Error())
	}
	ReturnSuccess(c, 0, param, param, 1)
}
