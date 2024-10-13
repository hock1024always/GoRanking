package router

import (
	"Ranking/config"
	"Ranking/controllers"
	"Ranking/pkg/logger"
	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// 路由 函数的名字要大写，这样才可以被其他包访问！
func Router() *gin.Engine {
	//创建一个路由的实例
	r := gin.Default()

	//日志中间件
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	//sessions中间件
	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	user := r.Group("/user")
	{
		// 注册用户相关的路由
		user.POST("/register", controllers.UserController{}.Register)
		// 登录用户相关的路由
		user.POST("/login", controllers.UserController{}.Login)
	}

	player := r.Group("/player")
	{
		player.POST("/list", controllers.PlayerController{}.GetPlayerList)
	}

	vote := r.Group("/vote")
	{
		vote.POST("/add", controllers.VoteController{}.AddVote)
	}

	r.POST("/ranking", controllers.PlayerController{}.GetRanking)

	//管理员
	controller := r.Group("/admin")
	{
		controller.POST("/register", controllers.Controller{}.Register)
		controller.POST("/activity", controllers.Controller{}.AddActivity)

		controller.POST("/ranking", controllers.PlayerController{}.GetRanking) //获取排行榜 方便下一步去更改某个player的分数
		controller.POST("/update_score", controllers.Controller{}.UpdatePlayersScore)
	}
	return r
}
