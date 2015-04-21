package controllers

import (
	"github.com/astaxie/beego"
	woothee "github.com/woothee/woothee-go"
)

type Meta struct {
	AppName         string
	AntennaURL      string
	BookURL         string
	VideoURL        string
	Copyright       string
	Author          string
	Email           string
	Keywords        string
	Description     string
	ApplicationName string
	Domain          string
	BaseDomain      string
	Host            string
	Url             string
	Type            string
	Title           string
	Image           string
	SiteName        string
	Locale          string
	FBAppId         string
	TWCard          string
	TWDomain        string
	TWSite          string
	TWImage         string
	RunMode         string
	UA              *woothee.Result
}

func NewMeta() *Meta {
	return &Meta{
		AppName:         beego.AppConfig.String("AppName"),
		AntennaURL:      beego.AppConfig.String("antennaurl"),
		BookURL:         beego.AppConfig.String("bookurl"),
		VideoURL:        beego.AppConfig.String("videourl"),
		Copyright:       beego.AppConfig.String("SiteName"),
		Author:          beego.AppConfig.String("Author"),
		Email:           beego.AppConfig.String("Email"),
		Keywords:        beego.AppConfig.String("Keywords"),
		Description:     beego.AppConfig.String("Description"),
		ApplicationName: beego.AppConfig.String("SiteName"),
		Domain:          beego.AppConfig.String("domain"),
		BaseDomain:      beego.AppConfig.String("basedomain"),
		Host:            "",
		Url:             "",
		Type:            "article",
		Title:           beego.AppConfig.String("SiteName"),
		Image:           beego.AppConfig.String("Image"),
		SiteName:        beego.AppConfig.String("SiteName"),
		Locale:          "ja_JP",
		FBAppId:         beego.AppConfig.String("FBAppId"),
		TWCard:          "summary_large_image",
		TWDomain:        "//",
		TWSite:          "//",
		TWImage:         beego.AppConfig.String("Image"),
		RunMode:         beego.RunMode,
	}
}

func (m *Meta) Init(c *BaseController) {
	m.Url = c.BuildRequestUrl("")
	m.Host = c.Ctx.Input.Site()
	m.UA, _ = woothee.Parse(c.Ctx.Input.UserAgent())
}
