package models

import "Ranking/dao"

type Activity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (Activity) TableName() string {
	return "activity" // 指定表名为 activity
}

func CheckActivityExist(activityname string) (Activity, error) {
	var activity Activity
	err := dao.Db.Where("name =?", activityname).First(&activity).Error
	return activity, err
}

func AddActivity(activityname string) (Activity, error) {
	activity := Activity{Name: activityname}
	err := dao.Db.Create(&activity).Error
	return activity, err
}

// 通过活动名获取活动id
func GetActivityIdByName(activityname string) (int, error) {
	var activity Activity
	err := dao.Db.Where("name =?", activityname).First(&activity).Error
	return activity.Id, err
}

// 将参赛者活动字段加入到activity表中
func AddPlayerToActivityActivity(playerId int, activityId int) error {
	var player Player
	err := dao.Db.Model(&player).Where("id =?", playerId).UpdateColumn("aid", activityId).Error
	return err
}
