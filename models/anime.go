package models

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ikeikeikeike/charneoapo"
	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/ikeikeikeike/gopkg/extract/image"
	"github.com/ikeikeikeike/gopkg/str"
)

type Anime struct {
	Id int64 `orm:"auto"`

	Name  string `orm:"size(128);unique"` // gin index
	Alias string `orm:"size(128);null"`   // gin index

	Kana   string `orm:"size(128);null"`
	Romaji string `orm:"size(128);null"`
	Gyou   string `orm:"size(6);index;null"`

	Url string `orm:"size(255);null"`

	Author string `orm:"size(128);null"` // 原作者
	Works  string `orm:"size(128);null"` // 制作会社

	ReleaseDate time.Time `orm:"type(date);index;null"`

	Outline string `orm:"type(text);null"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Icon *Image `orm:"rel(one);on_delete(set_null);index;null"`

	Characters []*Character `orm:"reverse(many)"`

	Pictures      []*Picture `orm:"reverse(many)"`
	PicturesCount int        `orm:"default(0);index"`
}

func (m *Anime) LoadRelated() *Anime {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Icon")
	_, _ = o.LoadRelated(m, "Pictures", 2, DefaultPerEntities, 0, "-id")
	_, _ = o.LoadRelated(m, "Characters", 2, DefaultPerEntities, 0, "-id")
	return m
}

func (m *Anime) RelLoader() {
	m.LoadRelated()
}

func (m *Anime) IdStr() string {
	return convert.ToStr(m.Id)
}

func (m *Anime) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Anime) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Anime) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Anime) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Anime) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Anime) UpdateIconByFileInfo(info *image.FileInfo, name string) error {
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

func (m *Anime) UpdateByNeoapo(n *charneoapo.Neoapo) error {
	// m.Kana = n.Kana()
	m.Alias = getAlias(n.AnimeAlias())
	m.Url = n.AnimeUrl()
	m.Author = n.AnimeAuthor()
	m.Works = n.AnimeWorks()
	m.ReleaseDate = n.AnimeRelease()
	m.Outline = n.Comment()

	return m.Update()
}

func Animes() orm.QuerySeter {
	return orm.NewOrm().QueryTable("anime").OrderBy("-Id")
}

func getAlias(alias string) string {
	alias = str.Clean(alias)
	alias = strings.Replace(alias, "、", ",", -1)
	alias = strings.Split(alias, ",")[0]
	return str.Clean(alias)
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Anime))
}
