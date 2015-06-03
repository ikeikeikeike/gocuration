package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ikeikeikeike/gopkg/convert"
)

type Ranking struct {
	Id int64 `orm:"auto"`

	BeginName string    `orm:"size(255);index"` // dayly, weekly, monthly, yearly
	BeginTime time.Time `orm:"type(datetime);index"`

	Rank     int64 `orm:"default(0);index"` // rank order number
	PageView int64 `orm:"default(0);index"` // 1,2,3,4

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Entry *Entry `orm:"rel(fk);index"`
}

func (m *Ranking) LoadRelated() *Ranking {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Entry", 2, DefaultPerEntities)
	return m
}

func (m *Ranking) RelLoader() {
	m.LoadRelated()
}

func (m *Ranking) IdStr() string {
	return convert.ToStr(m.Id)
}

func (m *Ranking) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Ranking) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Ranking) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Ranking) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Ranking) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Rankings() orm.QuerySeter {
	return orm.NewOrm().QueryTable("ranking").OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Ranking))
}
