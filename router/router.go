package router

import (
	"Ranking/controllers"
	"Ranking/pkg/logger"
	"github.com/gin-gonic/gin"
)

// 路由 函数的名字要大写，这样才可以被其他包访问！
func Router() *gin.Engine {
	//创建一个路由的实例
	r := gin.Default()

	//日志
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	user := r.Group("/user")
	{
		// 注册用户相关的路由
		user.POST("/register", controllers.UserController{}.Register)
	}
	return r
}
