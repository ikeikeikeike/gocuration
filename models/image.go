package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ikeikeikeike/gopkg/convert"
)

var (
	ImageSizeSmall     = 300
	ImageSizeMiddle    = 700
	ImageLinkAlphabets = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type Image struct {
	Id     int64
	Name   string `orm:"size(255);null"`
	Src    string `orm:"size(255);null;index"`
	Ext    string `orm:"size(255);null;index"`
	Mime   string `orm:"size(255);null"`
	Width  int    `orm:"null"`
	Height int    `orm:"null"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Picture *Picture `orm:"rel(fk);index;null"`

	Diva    *Diva    `orm:"reverse(one)"`
	Blog    *Blog    `orm:"reverse(one)"`
	Entries []*Entry `orm:"reverse(many)"`
}

func (m *Image) LoadRelated() *Image {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Diva")
	_, _ = o.LoadRelated(m, "Blog")
	_, _ = o.LoadRelated(m, "Entries", 2, DefaultPerEntities, 0, "-id")
	return m
}

func (m *Image) RelLoader() {
	m.LoadRelated()
}

func (m *Image) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Image) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Image) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Image) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Image) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Image) LinkFull() string {
	return m.LinkSize(0)
}

func (m *Image) LinkSmall() string {
	var width int
	switch {
	case m.Width > ImageSizeSmall:
		width = ImageSizeSmall
	}
	return m.LinkSize(width)
}

func (m *Image) LinkMiddle() string {
	var width int
	switch {
	case m.Width > ImageSizeMiddle:
		width = ImageSizeMiddle
	}
	return m.LinkSize(width)
}

func (m *Image) LinkSize(width int) string {
	if m.Ext == ".gif" {
		width = 0 // if image is gif then return full size
	}
	var size string
	switch width {
	case ImageSizeSmall, ImageSizeMiddle:
		size = convert.ToStr(width)
	default:
		size = "full"
	}
	return "/img/" + m.GetToken() + "." + size + m.Ext
}

func (m *Image) GetToken() string {
	number := beego.Date(m.Created, "ymds") + convert.ToStr(m.Id)
	return convert.NumberEncode(number, ImageLinkAlphabets)
}

func (m *Image) DecodeToken(token string) error {
	number := convert.NumberDecode(token, ImageLinkAlphabets)
	if len(number) < 9 {
		return fmt.Errorf("token `%s` too short <- `%s`", token, number)
	}

	if t, err := beego.DateParse(number[:8], "ymds"); err != nil {
		return fmt.Errorf("token `%s` date parse error <- `%s`", token, number)
	} else {
		m.Created = t
	}

	var err error
	m.Id, err = convert.StrTo(number[8:]).Int64()
	if err != nil {
		return fmt.Errorf("token `%s` id parse error <- `%s`", token, err)
	}

	return nil
}

func Images() orm.QuerySeter {
	return orm.NewOrm().QueryTable("image").OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Image))
}
