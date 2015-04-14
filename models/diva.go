package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ikeikeikeike/divaextractor"
	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/ikeikeikeike/gopkg/extract/image"
)

type Diva struct {
	Id int64 `orm:"auto"`

	Name   string `orm:"size(128);unique"` // gin index
	Kana   string `orm:"size(128);null"`   // gin index
	Romaji string `orm:"size(128);null"`
	Gyou   string `orm:"size(6);index;null"`

	Birthday time.Time `orm:"type(date);index;null"`

	Blood string `orm:"size(6);index;null"`

	Height int `orm:"default(0);index"`
	Weight int `orm:"default(0);index"`

	Bust   int    `orm:"default(0);index"`
	Waste  int    `orm:"default(0);index"`
	Hip    int    `orm:"default(0);index"`
	Bracup string `orm:"size(8);index;null"`

	Outline    string    `orm:"type(text);null"`
	Html       string    `orm:"type(text);null"`
	HtmlExpire time.Time `orm:"type(datetime);null"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	// Socials *Social   `orm:"rel(many);index;null"`  // 本人のソーシャルアカウントを取得してくる
	Icon *Image `orm:"rel(one);on_delete(set_null);index;null"`

	VideosCount int      `orm:"default(0);index"`
	Videos      []*Video `orm:"reverse(many)"`
}

func (m *Diva) LoadRelated() *Diva {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Icon")
	_, _ = o.LoadRelated(m, "Videos", 2, DefaultPerEntities, 0, "-id")
	return m
}

func (m *Diva) RelLoader() {
	m.LoadRelated()
}

func (m *Diva) IdStr() string {
	return convert.ToStr(m.Id)
}

func (m *Diva) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Diva) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Diva) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Diva) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Diva) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Diva) UpdateIconByFileInfo(info *image.FileInfo, name string) error {
	m.Icon.Name = name
	m.Icon.Src = info.Url
	m.Icon.Width = info.Width
	m.Icon.Height = info.Height
	m.Icon.Ext = info.Ext
	m.Icon.Mime = info.Mime

	return m.Icon.Update()
}

// May be adding data
func (m *Diva) UpdateByWikipedia(wiki *divaextractor.Wikipedia) error {

	if wiki.Birthday().Unix() > 1000 {
		m.Birthday = wiki.Birthday()
	}
	if wiki.Blood() != "" {
		m.Blood = wiki.Blood()
	}
	if wiki.Height() > 10 {
		m.Height = wiki.Height()
	}
	if wiki.Weight() > 10 {
		m.Weight = wiki.Weight()
	}
	if wiki.Bust() > 10 {
		m.Bust = wiki.Bust()
	}
	if wiki.Waist() > 10 {
		m.Waste = wiki.Waist()
	}
	if wiki.Hip() > 10 {
		m.Hip = wiki.Hip()
	}
	if wiki.Bracup() != "" {
		m.Bracup = wiki.Bracup()
	}

	return m.Update()
}

func Divas() orm.QuerySeter {
	return orm.NewOrm().QueryTable("diva").OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Diva))
}
