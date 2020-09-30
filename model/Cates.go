package model

type Categorys struct {
	Id				uint		`json:"id" gorm:"primary_key;not null;AUTO_INCREMENT" form:"id"`
	CategoryType	uint		`json:"category_type" form:"category_type"`
	Slug			string		`json:"slug" form:"slug" binging:"filter_url"`
	CommonContent
}
