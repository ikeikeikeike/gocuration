package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ikeikeikeike/charneoapo"
	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/ikeikeikeike/gopkg/extract/image"
)

type Character struct {
	Id int64 `orm:"auto"`

	Name string `orm:"size(128);index"`

	Kana   string `orm:"size(128);null"`
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

	Outline string `orm:"type(text);null"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Icon *Image `orm:"rel(one);on_delete(set_null);index;null"`

	Product string `orm:"size(128);index"` // XXX: Anime modelsになった
	Anime   *Anime `orm:"rel(fk);null"`    // TODO: Productをなくす

	PicturesCount int        `orm:"default(0);index"`
	Pictures      []*Picture `orm:"reverse(many)"`
}

// multiple fields unique key
func (u *Character) TableUnique() [][]string {
	return [][]string{
		[]string{"Product", "Name"},
	}
}

func (m *Character) LoadRelated() *Character {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Icon")
	_, _ = o.LoadRelated(m, "Anime")
	_, _ = o.LoadRelated(m, "Pictures", 2, DefaultPerEntities)
	return m
}

func (m *Character) RelLoader() {
	m.LoadRelated()
}

func (m *Character) IdStr() string {
	return convert.ToStr(m.Id)
}

func (m *Character) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Character) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Character) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Character) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Character) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Character) UpdateIconByFileInfo(info *image.FileInfo, name string) error {
	m.Icon.Name = name
	m.Icon.Src = info.Url
	m.Icon.Width = info.Width
	m.Icon.Height = info.Height
	m.Icon.Ext = info.Ext
	m.Icon.Mime = info.Mime

	err := m.Icon.Update()
	if err == nil {
		return m.Update()
	}
	return err
}

func (m *Character) UpdateByNeoapo(n *charneoapo.Neoapo) error {
	// m.Product = n.Product()
	// m.Name = n.Name()
	m.Kana = n.Kana()
	m.Birthday = n.Birthday()
	m.Blood = n.Blood()
	m.Height = n.Height()
	m.Weight = n.Weight()
	m.Bust = n.Bust()
	m.Waste = n.Waist()
	m.Hip = n.Hip()
	m.Bracup = n.Bracup()
	m.Outline = n.Comment()

	return m.Update()
}

func Characters() orm.QuerySeter {
	return orm.NewOrm().QueryTable("character").OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Character))
}
