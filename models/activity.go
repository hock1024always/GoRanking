package models

import "Ranking/dao"

type Activity struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	State int    `json:"state"` //标识状态 1为开 0为关闭
}

func (Activity) TableName() string {
	return "activity" // 指定表名为 activity
}

func CheckActivityExist(activityname string) (Activity, error) {
	var activity Activity
	err := dao.Db.Where("name =?", activityname).First(&activity).Error
	return activity, err
}

// 通过id查看活动是否存在
func CheckActivityExistById(activityId int) (Activity, error) {
	var activity Activity
	err := dao.Db.Where("id =?", activityId).First(&activity).Error
	return activity, err
}

func AddActivity(activityname string) (Activity, error) {
	activity := Activity{Name: activityname, State: 1}
	err := dao.Db.Create(&activity).Error
	return activity, err
}

// 通过id获取活动
func GetActivityById(id int) (Activity, error) {
	var activity Activity
	err := dao.Db.Where("id=?", id).First(&activity).Error
	return activity, err
}

// 将activityId字段加入到player表中
func AddPlayerToActivityActivity(playerId int, activityId int) error {
	var player Player
	err := dao.Db.Model(&player).Where("id =?", playerId).UpdateColumn("aid", activityId).Error
	return err
}

// 返回所有的活动
func GetAllActivity(sort string) ([]Activity, error) {
	var activity []Activity
	err := dao.Db.Order(sort).Find(&activity).Error
	return activity, err
}

// 返回所有可以参加的活动
func GetAllActivityAvailable(sort string) ([]Activity, error) {
	var activity []Activity
	err := dao.Db.Where("state =?", 1).Order(sort).Find(&activity).Error
	return activity, err
}

// 更新活动状态
func CloseActivity(activityId int, state int) error {
	err := dao.Db.Model(&Activity{}).Where("id =?", activityId).Update("state", state).Error
	return err
}
