package model

type Tags struct {
	ID			uint		`json:"id" gorm:"primary_key;not null;AUTO_INCREMENT"`
	Batch		uint		`json:"batch"`
	Tag			string		`json:"tag"`
	PicHash		string		`json:"pic_hash"`
	Score		uint		`json:"score"`
}