package models

import (
	"errors"
	"time"

	// "github.com/astaxie/beego"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id            int64
	Email         string    `orm:"size(64);unique" form:"Email" valid:"Required;Email"`
	Password      string    `orm:"size(32)" form:"Password" valid:"Required;MinSize(6)"`
	Lastlogintime time.Time `orm:"type(datetime);null" form:"-"`
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now;type(datetime)"`

	Blogs []*Blog `orm:"reverse(many)"`
}

// func (u *User) TableName() string {
// return beego.AppConfig.String("anything_table_name")
// }

func (m *User) LoadRelated() *User {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Blogs", 2, DefaultPerEntities, 0, "-id")
	return m
}

func (m *User) RelLoader() {
	m.LoadRelated()
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		if err.Error() == "UNIQUE constraint failed: user.email" {
			msg := "入力されたメールアドレスは既に登録されています。"
			err = errors.New(msg)
		}
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Users() orm.QuerySeter {
	var table User
	return orm.NewOrm().QueryTable(table).OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(User))
}
