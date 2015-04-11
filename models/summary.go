package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Summary struct {
	Id int64 `orm:"auto"`

	Sort int64 `orm:"default(0);index"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Entry  *Entry   `orm:"rel(one);on_delete(cascade);unique;null"`
	Scores []*Score `orm:"reverse(many)"` // 今は使用していないかも
}

func (m *Summary) LoadRelated() *Summary {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Entry")
	_, _ = o.LoadRelated(m, "Scores", 2, DefaultPerEntities)
	return m
}

func (m *Summary) RelLoader() {
	m.LoadRelated()
}

func (m *Summary) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Summary) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Summary) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Summary) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Summary) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Summaries() orm.QuerySeter {
	return orm.NewOrm().QueryTable("summary").OrderBy("-sort")
}

func PictureSummaries() orm.QuerySeter {
	return orm.NewOrm().QueryTable("summary").Filter("entry__picture__isnull", false).OrderBy("-sort")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Summary))
}
