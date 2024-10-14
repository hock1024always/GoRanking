package models

import (
	"Ranking/dao"
	"time"
)

type Vote struct {
	Id         int   `json:"id"`
	UserId     int   `json:"user_id"`
	PlayerId   int   `json:"player_id"`
	AddTime    int64 `json:"add_time"`
	ActivityId int   `json:"activity_id"`
}

func (Vote) TableName() string {
	return "vote"
}

// 用来检查是否投过票了
func GetVoteInfo(userId int, playerId int, activityId int) (Vote, error) {
	var vote Vote
	err := dao.Db.Where("user_id =? AND player_id =? AND activity_id =?", userId, playerId, activityId).First(&vote).Error
	return vote, err
}

// 实现投票的记录
func AddVote(userId int, playerId int, activityId int) (int, error) {
	vote := Vote{
		UserId:     userId,
		PlayerId:   playerId,
		AddTime:    time.Now().Unix(),
		ActivityId: activityId,
	}
	err := dao.Db.Create(&vote).Error
	return vote.Id, err
}
