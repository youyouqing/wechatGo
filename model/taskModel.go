package model

type Task struct {
	Id             int    `json:"id"`
	DeleteFlag     int    `json:"delete_flag"`
	RawCreateTime  string `json:"raw_create_time"`
	UserId         int    `json:"user_id"`
	NotifyTitle    string `json:"notify_title"`
	NotifyContent  string `json:"notify_content"`
	NotifyLevel    int    `json:"notify_level"`
	RepeatFlag     int    `json:"repeat_flag"`
	RepeatCrontStr string `json:"repeat_cront_str"`
}

// 任务添加模型
type TaskAdd struct {
	Id             int    `form:"id" json:"id"`
	DeleteFlag     int    `form:"delete_flag" json:"delete_flag"`
	RawCreateTime  string `form:"raw_create_time" json:"raw_create_time"`
	KnockTime      string `form:"knock_time" json:"knock_time"`
	UserId         int    `form:"user_id" json:"user_id"`
	NotifyTitle    string `form:"notify_title" json:"notify_title"`
	NotifyContent  string `form:"notify_content" json:"notify_content"`
	NotifyLevel    int    `form:"notify_level" json:"notify_level"`
	RepeatFlag     int    `form:"repeat_flag" json:"repeat_flag"`
	RepeatCrontStr string `form:"repeat_cront_str" json:"repeat_cront_str"`
}
