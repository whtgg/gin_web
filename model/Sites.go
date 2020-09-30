package model

type Site struct {
	ID              uint   `gorm:"primary_key" json:"id" form:"id"`
	UserId          uint   `json:"user_id"`
	Domain          string `json:"domain" form:"domain" binding:"required"`
	Status          uint   `json:"status,omitempty" form:"status"`
	Qq              string `json:"qq,omitempty" form:"qq"`
	Qr              string `json:"qr,omitempty" form:"qr"`
	Phone           string `json:"phone"`
	SiteUrl         string `json:"site_url,omitempty" form:"site_url" binding:"filter_url"`
	Icon            string `json:"icon,omitempty" form:"icon"`
	Logo            string `json:"logo,omitempty" form:"logo"`
	Notice          string `json:"notice,omitempty" form:"notice"`
	SafeTips        string `json:"safe_tips,omitempty" form:"safe_tips"`
	ScriptTop       string `json:"script_top,omitempty" form:"script_top"`
	ScriptAppend    string `json:"script_append,omitempty" form:"script_append"`
	ScriptMarketing string `json:"script_marketing,omitempty" form:"script_marketing"`
	TplId           uint   `json:"tpl_id,omitempty" form:"tpl_id"`
	Sort            uint   `json:"sort,omitempty" form:"sort"`
	WebType         uint   `json:"web_type,omitempty" form:"web_type"`
	State           uint   `json:"state,omitempty" form:"state"`
	Db				string `json:"db"`
	CommonDesc
}
