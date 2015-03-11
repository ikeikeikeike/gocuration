package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ikeikeikeike/gopkg/convert"
)

type Score struct {
	Id    int64  `orm:"auto"`
	Name  string `orm:"size(255);index"` // hatena, twitter, facebook, in, out
	Count int64  `orm:"default(0);index"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Blog    *Blog    `orm:"rel(fk);index;null"`
	Entry   *Entry   `orm:"rel(fk);index;null"`
	Summary *Summary `orm:"rel(fk);on_delete(set_null);index;null"` // 今は使用していないかも
}

// multiple fields unique key
func (u *Score) TableUnique() [][]string {
	return [][]string{
		// []string{"Name", "Blog"},
		[]string{"Name", "Entry"},
	}
}

func (m *Score) LoadRelated() *Score {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Entry", 2, DefaultPerEntities)
	return m
}

func (m *Score) RelLoader() {
	m.LoadRelated()
}

func (m *Score) IdStr() string {
	return convert.ToStr(m.Id)
}

func (m *Score) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Score) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Score) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Score) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Score) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Scores() orm.QuerySeter {
	return orm.NewOrm().QueryTable("blog").OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Score))
}
