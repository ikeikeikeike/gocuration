package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ikeikeikeike/gopkg/convert"
)

type Video struct {
	Id int64 `orm:"auto"`

	Url      string `orm:"size(255);null" form:"Url" valid:"Required;Match(/^https?/)"`
	Code     string `orm:"type(text);null"` // TODO: Change default value later.
	Duration int    `orm:"default(0);index"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Entry *Entry `orm:"rel(one);unique;null"`

	Site *Site `orm:"rel(fk);index;null"`

	Divas []*Diva `orm:"rel(m2m);index;rel_through(bitbucket.org/ikeikeikeike/antenna/models.VideoDiva)"`
}

func (m *Video) LoadRelated() *Video {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Site")
	_, _ = o.LoadRelated(m, "Entry")
	_, _ = o.LoadRelated(m, "Divas", 2, DefaultPerEntities, 0, "-id")
	return m
}

func (m *Video) RelLoader() {
	m.LoadRelated()
}

func (m *Video) IdStr() string {
	return convert.ToStr(m.Id)
}

func (m *Video) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Video) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Video) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Video) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Video) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Videos() orm.QuerySeter {
	return orm.NewOrm().QueryTable("video").OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Video))
}
