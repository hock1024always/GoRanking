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

func GetPlayers(aid int) ([]Player, error) {
	var players []Player
	err := dao.Db.Where("aid =?", aid).Find(&players).Error
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
