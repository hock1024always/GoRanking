package controllers

import (
	"crypto/md5"
	"encoding/hex"
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
	//code：响应码，msg：泛型信息，data：泛型数据，count：信息条数
	json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(http.StatusOK, json)
}

func ReturnError(c *gin.Context, code int, msg string) {
	//code：响应码，msg：错误信息
	json := &JsonErrStruct{Code: code, Msg: msg}
	c.JSON(http.StatusOK, json)
}

func EncryMd5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}
