# 文档说明

文件格式：```Json```

# 路由实现

首先，使用路由的形式展示整体项目的结构功能

## 用户

```go
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
```

## 参赛者

```go
player := r.Group("/player")
	{
		//注册参赛者相关的路由 nickname password confirm_password
		player.POST("/register", controllers.PlayerController{}.PlayerRegister)
		//实现用户选择自己想参与的活动  activity_id nickname password
		player.POST("/add_activity", controllers.PlayerController{}.PlayerChooseActivity)
		//实现添加自己的参赛宣言功能 declaration nickname password
		player.POST("/add_declaration", controllers.PlayerController{}.UpdateDeclaration)
		//获取目前可以参加的活动列表 GET
		player.GET("/get_activitys", controllers.ActivityController{}.GetActivityListForPlayer)
		//退出正在参与的活动 nickname password
		player.POST("/quit_activity", controllers.PlayerController{}.QuitActivity)
		////注销参赛者 和删除用户逻辑一样，不写了
		//player.POST("/delete", controllers.PlayerController{}.PlayerDelete)
		//获取给自己投票的玩家列表 nickname password
		player.POST("/get_vote_players", controllers.PlayerController{}.GetVoteUsers)
		//获取某个活动中给自己投票的玩家列表 nickname password activity_id
		player.POST("/get_vote_players_in_activity", controllers.PlayerController{}.GetVoteUsersInActivity)
		//获取某项活动的排行榜 activity_id 不合理(获取的是总分)
		player.POST("/get_activity_ranking", controllers.PlayerController{}.GetRanking)
	}
```

## 管理员

```go
controller := r.Group("/admin")
	{
		//注册管理员相关的路由 admin_name password confirm_password key
		controller.POST("/register", controllers.Controller{}.Register)
		//添加 admin_name password
		controller.POST("/activity", controllers.Controller{}.AddActivity)
		//获取参赛者的总分列表 admin_name password activity_name
		controller.POST("/ranking", controllers.Controller{}.GetAllRanking) //获取排行榜 方便下一步去更改某个player的分数
		//更新某个player的分数 player_name admin_name password update_score
		controller.POST("/update_score", controllers.Controller{}.UpdatePlayersScore)
		//将某个活动关闭 admin_name password activity_id
		controller.POST("/close_activity", controllers.Controller{}.CloseActivity)
		//去除某个参赛者在某项活动的得分 admin_name password activity_id player_id
		controller.POST("/delete_player_score", controllers.Controller{}.DeletePlayerScore)
		//去除某个用户对于某个活动的投票 admin_name password activity_id user_id
		controller.POST("/delete_vote", controllers.Controller{}.DeleteVote)
	}
```

## 活动

```go
activity := r.Group("/activity")
	{
		//获取所有活动列表 GET
		activity.GET("/list", controllers.ActivityController{}.GetActivityList)
		//获取参与某项活动的参赛者列表 activity_id
		player.POST("/player_list", controllers.PlayerController{}.GetPlayerList)
	}
```

# 返回值说明

## 用户

### 注册

```json
{
    "code": 0,
    "data": {
        "username": "ht",
        "userid": 38
    },
    "msg": "注册成功",
    "count": 1
}

```

### 登陆

```json
{
    "code": 0,
    "data": {
        "id": 38,
        "username": "ht"
    },
    "msg": "登录成功",
    "count": 1
}

```

### 投票

```json
{
    "code": 0,
    "data": 22,
    "msg": "投票成功",
    "count": 1
}

```

### 删除投票

```json
{
    "code": 0,
    "data": null,
    "msg": "删除投票成功",
    "count": 1
}

```

### 注销用户

```json
{
    "code": 0,
    "data": null,
    "msg": "删除成功",
    "count": 1
}

```

### 获取自己的投票记录

```json
{
    "code": 0,
    "data": [
        {
            "id": 24,
            "user_id": 38,
            "player_id": 1,
            "add_time": 1729058838,
            "activity_id": 1
        },
        {
            "id": 23,
            "user_id": 38,
            "player_id": 2,
            "add_time": 1729058710,
            "activity_id": 1
        }
    ],
    "msg": "获取投票列表成功",
    "count": 1
}

```

### 修改密码

```json
{
    "code": 0,
    "data": "新密码是:111111",
    "msg": "修改密码成功",
    "count": 1
}

```

## 参赛者

### 注册

```json
{
    "code": 0,
    "data": {
        "id": 6,
        "nickname": "MJ",
        "aid": 0,
        "ref": 0,
        "avatar": 0,
        "score": 0,
        "declaration": "",
        "password": "111111"
    },
    "msg": "注册成功",
    "count": 1
}

```

### 参与活动

```json
{
    "code": 0,
    "data": null,
    "msg": "参加活动成功",
    "count": 1
}

```

### 修改参赛宣言

```json
{
    "code": 0,
    "data": "宣言更改为：原来你也玩原神，原神启动",
    "msg": "宣言更改成功",
    "count": 1
}

```

### 查看可参加的活动列表(GET)

```json
{
    "code": 0,
    "data": [
        {
            "id": 1,
            "name": "run",
            "state": 1
        },
        {
            "id": 2,
            "name": "climb",
            "state": 1
        }
    ],
    "msg": "获取成功",
    "count": 1
}

```

### 退出正在参与的活动

```json
{
    "code": 0,
    "data": "被扣除积分为：0",
    "msg": "退出活动成功",
    "count": 1
}

```

### 获得给自己投票的玩家列表

```json
{
    "code": 0,
    "data": [
        {
            "id": 19,
            "user_id": 31,
            "player_id": 6,
            "add_time": 0,
            "activity_id": 1
        },
        {
            "id": 17,
            "user_id": 32,
            "player_id": 6,
            "add_time": 0,
            "activity_id": 1
        }
    ],
    "msg": "获取投票列表成功",
    "count": 1
}

```

### 获得在某个活动中给自己投票的玩家列表

```json
{
    "code": 0,
    "data": [
        {
            "id": 19,
            "user_id": 31,
            "player_id": 6,
            "add_time": 0,
            "activity_id": 1
        },
        {
            "id": 17,
            "user_id": 32,
            "player_id": 6,
            "add_time": 0,
            "activity_id": 1
        }
    ],
    "msg": "获取投票列表成功",
    "count": 1
}

```

### 获得某项活动的排行榜（待重构）

想使用借鉴```BTC```的交易型结构，通过遍历数据库vote的条目来得出某项活动中各个参赛者的得分。

## 管理员

### 注册

```json
{
    "code": 0,
    "data": {
        "id": 0,
        "controller_name": "AweiRu",
        "password": "111111"
    },
    "msg": "注册成功",
    "count": 1
}

```

### 添加活动

```json
{
    "code": 0,
    "data": {
        "id": 3,
        "name": "rush",
        "state": 1
    },
    "msg": "注册成功",
    "count": 1
}

```

### 获取参赛者总分列表

```json
{
    "code": 0,
    "data": [
        {
            "id": 2,
            "nickname": "Yossi",
            "aid": 1,
            "ref": 2,
            "avatar": 1122,
            "score": 124,
            "declaration": "我是吆西，日服球王",
            "password": "111111"
        },
        {
            "id": 1,
            "nickname": "manba",
            "aid": 1,
            "ref": 1,
            "avatar": 1112,
            "score": 100,
            "declaration": "我是牢大，点赞复活我",
            "password": "111111"
        },
        {
            "id": 4,
            "nickname": "Zhyz",
            "aid": 2,
            "ref": 4,
            "avatar": 3423,
            "score": 100,
            "declaration": "球队输了，我詹没输",
            "password": "111111"
        },
        {
            "id": 5,
            "nickname": "iqun",
            "aid": 1,
            "ref": 0,
            "avatar": 0,
            "score": 100,
            "declaration": "我爱玩原神，原神启动",
            "password": "111111"
        },
        {
            "id": 6,
            "nickname": "MJ",
            "aid": 1,
            "ref": 0,
            "avatar": 0,
            "score": 100,
            "declaration": "原来你也玩原神，原神启动",
            "password": "111111"
        },
        {
            "id": 3,
            "nickname": "Cr7",
            "aid": 0,
            "ref": 3,
            "avatar": 7777,
            "score": 99,
            "declaration": "我是罗，爱越位和偷金球",
            "password": "111111"
        }
    ],
    "msg": "获取成功",
    "count": 1
}

```

### 更新某个参赛者总分

```json
{
    "code": 0,
    "data": "Yossi更新之后的得分为:100",
    "msg": "更新成功",
    "count": 1
}

```

### 关闭某项活动

```json
{
    "code": 0,
    "data": {
        "id": 1,
        "name": "run",
        "state": 1
    },
    "msg": "关闭成功",
    "count": 1
}

```

### 去除某个参赛者在某项活动中的得分

```json
{
    "code": 0,
    "data": "删除了MJ在run活动的2票",
    "msg": "删除成功",
    "count": 1
}

```

### 去除某个用户对于某个活动的投票

```json
{
    "code": 0,
    "data": "扣分的参赛者有:1,2",
    "msg": "删除成功",
    "count": 1
}

```

## 活动

### 获取所有活动列表

```json
{
    "code": 0,
    "data": [
        {
            "id": 1,
            "name": "run",
            "state": 0
        },
        {
            "id": 2,
            "name": "climb",
            "state": 1
        },
        {
            "id": 3,
            "name": "rush",
            "state": 1
        }
    ],
    "msg": "获取成功",
    "count": 1
}

```

### 获取参与某项活动的参赛者列表

```json
{
    "code": 0,
    "data": [
        {
            "id": 1,
            "nickname": "manba",
            "aid": 1,
            "ref": 1,
            "avatar": 1112,
            "score": 99,
            "declaration": "我是牢大，点赞复活我",
            "password": "111111"
        },
        {
            "id": 2,
            "nickname": "Yossi",
            "aid": 1,
            "ref": 2,
            "avatar": 1122,
            "score": 99,
            "declaration": "我是吆西，日服球王",
            "password": "111111"
        },
        {
            "id": 5,
            "nickname": "iqun",
            "aid": 1,
            "ref": 0,
            "avatar": 0,
            "score": 100,
            "declaration": "我爱玩原神，原神启动",
            "password": "111111"
        },
        {
            "id": 6,
            "nickname": "MJ",
            "aid": 1,
            "ref": 0,
            "avatar": 0,
            "score": 98,
            "declaration": "原来你也玩原神，原神启动",
            "password": "111111"
        }
    ],
    "msg": "获取成功",
    "count": 1
}

```

