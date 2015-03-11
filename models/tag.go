package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id     int64  `orm:"auto"`
	Name   string `orm:"size(255);unique"`
	Kana   string `orm:"size(255);null"`
	Romaji string `orm:"size(255);null"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Image   *Image   `orm:"rel(one);on_delete(set_null);index;null"`
	Entries []*Entry `orm:"reverse(many)"`
}

func (m *Tag) LoadRelated() *Tag {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Image")
	_, _ = o.LoadRelated(m, "Entries", 2, DefaultPerEntities, 0, "-id")
	return m
}

func (m *Tag) RelLoader() {
	m.LoadRelated()
}

func (m *Tag) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Read(fields ...string) error {
	return orm.NewOrm().Read(m, fields...)
}

func (m *Tag) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Tag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Tags() orm.QuerySeter {
	return orm.NewOrm().QueryTable("tag").OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Tag))
}
