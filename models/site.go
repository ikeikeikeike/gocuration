package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ikeikeikeike/gopkg/convert"
)

type Site struct {
	Id int64 `orm:"auto"`

	Name   string `orm:"size(255);null"`
	Domain string `orm:"size(255);unique"`

	Outline string `orm:"type(text);null"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Icon   *Image   `orm:"rel(one);on_delete(set_null);index;null"`
	Videos []*Video `orm:"reverse(many)"`
}

func (m *Site) LoadRelated() *Site {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Icon")
	_, _ = o.LoadRelated(m, "Videos", 2, DefaultPerEntities, 0, "-id")
	return m
}

func (m *Site) RelLoader() {
	m.LoadRelated()
}

func (m *Site) IdStr() string {
	return convert.ToStr(m.Id)
}

func (m *Site) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Site) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Site) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Site) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Site) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Sites() orm.QuerySeter {
	return orm.NewOrm().QueryTable("site").OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Site))
}
