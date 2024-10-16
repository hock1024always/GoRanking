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

// 删除涉及玩家和活动的投票记录 并返回得分
func DeletePlayerByActivityId(playerId, activityId int) (int, error) {
	// 定义一个变量来存储删除的记录数量
	result := dao.Db.Where("player_id = ? AND activity_id = ?", playerId, activityId).Delete(&Vote{})
	var voteNum int = int(result.RowsAffected)
	// 返回删除的条目数量和可能发生的错误
	return voteNum, result.Error
}

// 通过删除记录来扣分
func ReduceScore(playerId int, voteNum int) error {
	var player Player
	err := dao.Db.Model(&player).Where("id =?", playerId).UpdateColumn("score", gorm.Expr("score - ?", voteNum)).Error
	return err
}

func UpdatePlayerAid(id int, aid int) error {
	var player Player
	err := dao.Db.Model(&player).Where("id =?", id).Update("aid", aid).Error
	return err
}

// 查看给这个参赛者投票的投票记录
func GetVoteListForPlayer(playerId int, sort string) ([]Vote, error) {
	var votes []Vote
	err := dao.Db.Where("player_id =?", playerId).Order(sort).Find(&votes).Error
	return votes, err
}

// 查看给这个活动投票给某个用户的投票记录
func GetVoteListForPlayerByActivityId(playerId int, activityId int, sort string) ([]Vote, error) {
	var votes []Vote
	err := dao.Db.Where("player_id =? AND activity_id =?", playerId, activityId).Order(sort).Find(&votes).Error
	return votes, err
}
