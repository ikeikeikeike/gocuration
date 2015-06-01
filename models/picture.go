package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ikeikeikeike/gopkg/convert"
)

type Picture struct {
	Id         int64 `orm:"auto"`
	PageView   int64 `orm:"default(0);index"`
	ImageCount int   `orm:"default(0);index"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Entry *Entry `orm:"rel(one);unique;null"`

	Anime *Anime `orm:"rel(fk);null"`

	Images     []*Image     `orm:"reverse(many)"`
	Characters []*Character `orm:"rel(m2m);index;rel_through(bitbucket.org/ikeikeikeike/antenna/models.PictureCharacter)"`
}

func (m *Picture) LoadRelated() *Picture {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Entry")
	_, _ = o.LoadRelated(m, "Anime")
	_, _ = o.LoadRelated(m, "Images", 2, DefaultPerEntities, 0, "-id")
	_, _ = o.LoadRelated(m, "Characters", 2, DefaultPerEntities, 0, "-id")
	return m
}

func (m *Picture) RelLoader() {
	m.LoadRelated()
}

func (m *Picture) IdStr() string {
	return convert.ToStr(m.Id)
}

func (m *Picture) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Picture) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Picture) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Picture) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Picture) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Pictures() orm.QuerySeter {
	return orm.NewOrm().QueryTable("picture").OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Picture))
}
