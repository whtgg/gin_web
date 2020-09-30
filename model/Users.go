package model

import (
	"time"
)

// 基本模型的定义
type Model struct {
	ID        uint `gorm:"primary_key;not null;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID          uint 					`gorm:"primary_key" json:"id"`
	LoginName   string
	LoginPwd    string
	Nickname    string
	Avatar      string
	AccessToken string					`json:"token"`
	SafeCode    string
	Status      uint
}