package model

type User struct {
	ID         uint   `gorm:"primary_key;AUTO_INCREMENT"`
	OpenId     string `gorm:"size:255"`
	WxNickname string
	UserAvatar string
	WxAppid    string
}
