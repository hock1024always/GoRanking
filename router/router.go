package router

import (
	"Ranking/controllers"
	"Ranking/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由 函数的名字要大写，这样才可以被其他包访问！
func Router() *gin.Engine {
	//创建一个路由的实例
	r := gin.Default()

	//日志
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	//实现GET路由 获取
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	//创建这样一个组，简化代码
	user := r.Group("/user")
	{
		//查询单条数据
		user.GET("/info/:id", controllers.UserController{}.GetUserInfo)
		//查询列表数据
		user.POST("/list", controllers.UserController{}.GetList)
		//添加数据
		user.POST("/add", controllers.UserController{}.AddUser)
		//修改数据
		user.POST("/update", controllers.UserController{}.UpdateUser)
		//删除单个用户的数据
		user.POST("/delete", controllers.UserController{}.DeleteUser)
		//获取用户列表
		user.GET("/info/list", controllers.UserController{}.GetAllUsers)
		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user delete")
		})
	}

	order := r.Group("/order")
	{
		order.GET("/list", controllers.OrderController{}.GetList)
	}

	return r
}
