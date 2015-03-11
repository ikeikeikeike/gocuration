package models

import (
	"time"

	// "github.com/astaxie/beego"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type Tmpuser struct {
	Id          int64  `orm:"auto"`
	Rss         string `orm:"size(255)" form:"Rss" valid:"Required;Match(/^https?/)"`
	Email       string `orm:"size(64)" form:"Email" valid:"Required;Email"`
	Password    string `orm:"size(32)" form:"Password" valid:"Required;MinSize(6)"`
	Repassword  string `orm:"-" form:"Repassword" valid:"Required"`
	Mediatype   string `orm:"size(16)" form:"Mediatype" valid:"Required;Match(/^(movie|image)$/)"`
	Adsensetype string `orm:"size(16)" form:"Adsensetype" valid:"Required;Match(/^(2d|3d)$/)"`

	Using     string    `orm:"size(32);null" form:"-"`
	Token     string    `orm:"size(32)" form:"-"`
	ExpiredAt time.Time `orm:"type(datetime);null" form:"-"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

// func (u *Tmpuser) TableName() string {
// return beego.AppConfig.String("anything_table_name")
// }

func (u *Tmpuser) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "2つのパスワードが一致していません")
	}
}

func (m *Tmpuser) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Tmpuser) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Tmpuser) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Tmpuser) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Tmpuser) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Tmpusers() orm.QuerySeter {
	var table Tmpuser
	return orm.NewOrm().QueryTable(table).OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Tmpuser))
}
