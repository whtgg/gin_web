package model

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/asaskevich/govalidator"
)

type CommonContent struct {
	ParentId		uint		`json:"parent_id" form:"parent_id"`
	AdsPic			string		`json:"ads_pic" form:"ads_pic"`
	AdsTarget		string		`json:"ads_target" form:"ads_target"`
	Subtitle		string		`json:"subtitle" form:"subtitle"`
	Body			string		`json:"body" form:"body"`
	Password 		string		`json:"password"`
	Identity		string		`json:"identity" form:"identity"`
	PublishAt		uint		`json:"publish_at" form:"publish_at"`
	Subscript		uint		`json:"subscript" gorm:"-" form:"subscript"`
	Sort			uint		`json:"sort" form:"sort"`
	IsTop			uint		`json:"is_top" form:"is_top"`
	Pv				uint		`json:"pv" form:"pv"`
	Status			uint		`json:"status" form:"status"`
	CommonDesc
}

type CommonDesc struct {
	Title			string		`json:"title" form:"title"`
	Keywords		string		`json:"keywords" form:"keywords"`
	Description		string		`json:"description" form:"description"`
	Image			string		`json:"image" form:"image"`
	ImageAlt		string		`json:"image_alt" form:"image_alt"`
	Thumbnail		string		`json:"thumbnail" form:"thumbnail"`
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("filter_url", filterUrl)
	}
}

func filterUrl (f validator.FieldLevel) bool {
	return govalidator.IsURL(f.Field().String())
}

