package models

import (
	"Ranking/dao"
	"gorm.io/gorm"
)

type Player struct {
	Id          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Aid         int    `json:"aid"`         // 活动名称
	Ref         int    `json:"ref"`         // 活动序号
	Avatar      int    `json:"avatar"`      // 头像序号
	Score       int    `json:"score"`       // 积分
	Declaration string `json:"declaration"` // 将类型改为 string
	// UpdateTime  int    `json:"updateTime"`
}

func (Player) TableName() string {
	return "player"
}

// 获取某种顺序排列的某一活动的玩家列表 DESC降序 ASC升序
func GetPlayers(aid int, sort string) ([]Player, error) {
	var players []Player
	err := dao.Db.Where("aid =?", aid).Order(sort).Find(&players).Error
	return players, err
}

func GetPlayerById(id int) (Player, error) {
	var player Player
	err := dao.Db.Where("id =?", id).First(&player).Error
	return player, err
}

// 通过投票来更新得分
func UpdateScoreByVote(id int) {
	var player Player
	dao.Db.Model(&player).Where("id =?", id).UpdateColumn("score", gorm.Expr("score + ?", 1))
}
