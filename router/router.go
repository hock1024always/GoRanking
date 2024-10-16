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
		// 注册用户相关的路由 username password confirm_password
		user.POST("/register", controllers.UserController{}.Register)
		// 登录用户相关的路由 username password
		user.POST("/login", controllers.UserController{}.Login)
		//实现投票功能的路由 activity_id username password player_id
		user.POST("/vote", controllers.VoteController{}.AddVote)
		//删除投票功能的路由 activity_id username password player_id
		user.POST("/delete_vote", controllers.VoteController{}.DeleteVote)
		//实现删除用户的路由 username password confirm_sentence
		user.POST("/delete", controllers.UserController{}.UserDelete)
		// 实现用户获取自己的投票记录 username password
		user.POST("/get_vote_list", controllers.UserController{}.GetVoteList)
		//实现用户修改密码 username password new_password confirm_new_password
		user.POST("/modify_password", controllers.UserController{}.ModifyPassword)
	}

	player := r.Group("/player")
	{
		//注册参赛者相关的路由 nickname password confirm_password
		player.POST("/register", controllers.PlayerController{}.PlayerRegister)
		//实现用户选择自己想参与的活动  activity_id username password
		player.POST("/add_activity", controllers.PlayerController{}.PlayerChooseActivity)
		//实现添加自己的参赛宣言功能
		player.POST("/add_declaration", controllers.PlayerController{}.UpdateDeclaration)
		//GET 获取目前可以参加的活动列表
		player.GET("/get_activitys", controllers.ActivityController{}.GetActivityListForPlayer)
		////退出正在参与的活动
		//player.POST("/quit_activity", controllers.PlayerController{}.QuitActivity)
		////注销参赛者
		//player.POST("/delete", controllers.PlayerController{}.PlayerDelete)
	}

	//管理员
	controller := r.Group("/admin")
	{
		//注册管理员相关的路由 username password confirm_password key
		controller.POST("/register", controllers.Controller{}.Register)
		//添加
		controller.POST("/activity", controllers.Controller{}.AddActivity)
		//获取参赛者的总分列表
		//controller.POST("/ranking", controllers.PlayerController{}.GetAllRanking) //获取排行榜 方便下一步去更改某个player的分数
		//更新某个player的分数
		controller.POST("/update_score", controllers.Controller{}.UpdatePlayersScore)
		////将某个活动关闭
		//controller.POST("/close_activity", controllers.Controller{}.CloseActivity)
		////去除某个参赛者在某项活动的得分
		//controller.POST("/delete_player_score", controllers.Controller{}.DeletePlayerScore)
		////去除某个用户对于某个活动的投票
		//controller.POST("/delete_vote", controllers.Controller{}.DeleteVote)
	}

	activity := r.Group("/activity")
	{
		//获取所有活动列表
		activity.GET("/list", controllers.ActivityController{}.GetActivityList)
		//获取参与某项活动的参赛者列表 aid
		player.POST("/player_list", controllers.PlayerController{}.GetPlayerList)
		//获取某项活动的排行榜 aid
		r.POST("/ranking", controllers.PlayerController{}.GetRanking)
	}
	return r
}
