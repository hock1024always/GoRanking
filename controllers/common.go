package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonStruct struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   interface{} `json:"msg"`
	Count int64       `json:"count"`
}

type JsonErrStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	//csde：响应码，msg：泛型信息，data：泛型数据，count：信息条数
	json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(http.StatusOK, json)

}

func ReturnError(c *gin.Context, code int, msg string) {
	json := &JsonErrStruct{Code: code, Msg: msg}
	c.JSON(http.StatusOK, json)
}

func ReturnUserGetListError(c *gin.Context, code int, msg int) {
	json := &JsonErrStruct{Code: code, Msg: msg}
	c.JSON(http.StatusOK, json)
}
