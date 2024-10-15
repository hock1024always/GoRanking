package models

import (
	"Ranking/dao"
	"gorm.io/gorm"
)

type Player struct {
	Id          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Aid         int    `json:"aid"`         // 活动序号
	Ref         int    `json:"ref"`         // 注销号码
	Avatar      int    `json:"avatar"`      // 头像序号
	Score       int    `json:"score"`       // 积分
	Declaration string `json:"declaration"` // 将类型改为 string
	Password    string `json:"password"`
	// UpdateTime  int    `json:"updateTime"`
}

func (Player) TableName() string {
	return "player"
}

func AddPlayer(nickname, password string) (Player, error) {
	player := Player{Nickname: nickname, Password: password}
	err := dao.Db.Create(&player).Error
	return player, err
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

// 通过投票来更新得分 加分
func UpdateScoreByVote(id int) error {
	var player Player
	err := dao.Db.Model(&player).Where("id =?", id).UpdateColumn("score", gorm.Expr("score + ?", 1)).Error
	return err
}

// 通过活动来更新得分 减分
func UpdateScoreByActivity(playerId int) error {
	var player Player
	err := dao.Db.Model(&player).Where("id =?", playerId).UpdateColumn("score", gorm.Expr("score - ?", 1)).Error
	return err
}

// 通过管理员修改来更新得分
func UpdateScoreByAdmin(nickname string, score int) error {
	var player Player
	err := dao.Db.Model(&player).Where("nickname =?", nickname).UpdateColumn("score", score).Error
	return err
}

func CheckPlayerExistsByNickname(nickname string) (Player, error) {
	var player Player
	err := dao.Db.Where("nickname =?", nickname).First(&player).Error
	return player, err
}

// 更改宣言
func AddDeclaration(id int, declaration string) (Player, error) {
	var player Player
	err := dao.Db.Model(&player).Where("id =?", id).Update("declaration", declaration).Error
	return player, err
}
