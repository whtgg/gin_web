package model


type Articles struct {
	Id				uint		`json:"id" gorm:"primary_key;not null;AUTO_INCREMENT" form:"id"`
	CategoryId		uint		`json:"category_id" gorm:"index" form:"category_id"`
	UserId			uint		`json:"user_id" gorm:"index"`
	Tag				string		`json:"tag"`
	CommonContent
}
