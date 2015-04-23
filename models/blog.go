package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ikeikeikeike/gopkg/convert"
)

type Blog struct {
	Id          int64  `orm:"auto"`
	Rss         string `orm:"size(255);null;unique" form:"Rss" valid:"Required;Match(/^https?/)"`
	Url         string `orm:"size(255);null"`
	Name        string `orm:"size(255);null" form:"Name" valid:"MaxSize(250)"`
	Mediatype   string `orm:"size(16)" form:"Mediatype" valid:"Required;Match(/^(movie|image)$/)"`
	Adsensetype string `orm:"size(16)" form:"Adsensetype" valid:"Required;Match(/^(2d|3d)$/)"`

	VerifyParts     int `orm:"default(1);null" form:"VerifyParts" valid:"Range(0,3)"`
	VerifyRss       int `orm:"default(1);null" form:"VerifyRss" valid:"Range(0,3)"`
	VerifyLink      int `orm:"default(1);null" form:"VerifyLink" valid:"Range(0,3)"`
	VerifyBookRss   int `orm:"default(1);null" form:"VerifyBookRss" valid:"Range(0,3)"`
	VerifyBookLink  int `orm:"default(1);null" form:"VerifyBookLink" valid:"Range(0,3)"`
	VerifyVideoRss  int `orm:"default(1);null" form:"VerifyVideoRss" valid:"Range(0,3)"`
	VerifyVideoLink int `orm:"default(1);null" form:"VerifyVideoLink" valid:"Range(0,3)"`

	IsPenalty bool `orm:"default(0)"`

	LastModified time.Time `orm:"type(datetime);null;index"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	User *User  `orm:"rel(fk);null;index"`
	Icon *Image `orm:"rel(one);on_delete(set_null);index;null"`

	Scores  []*Score `orm:"reverse(many)"`
	Entries []*Entry `orm:"reverse(many)"`
}

func (m *Blog) VerifyScore() int {
	var score int = 1

	if m.VerifyParts >= 3 {
		score++
		score++
		score++
	}
	if m.VerifyRss >= 3 {
		score++
	}
	if m.VerifyLink >= 3 {
		score++
	}
	if m.VerifyBookRss >= 3 {
		score++
	}
	if m.VerifyBookLink >= 3 {
		score++
	}
	if m.VerifyVideoRss >= 3 {
		score++
	}
	if m.VerifyVideoLink >= 3 {
		score++
	}
	return score
}

func (m *Blog) LoadRelated() *Blog {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "User")
	_, _ = o.LoadRelated(m, "Icon")
	_, _ = o.LoadRelated(m, "Scores", 2, DefaultPerEntities)
	_, _ = o.LoadRelated(m, "Entries", 2, DefaultPerEntities, 0, "-id")
	return m
}

func (m *Blog) RelLoader() {
	m.LoadRelated()
}

func (m *Blog) IdStr() string {
	return convert.ToStr(m.Id)
}

func (m *Blog) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		if err.Error() == "UNIQUE constraint failed: blog.rss" {
			msg := "入力されたRSSは既に登録されています。"
			err = errors.New(msg)
		}
		return err
	}
	return nil
}

func (m *Blog) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Blog) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *Blog) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Blog) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

// Ligical delete: Unrelate User and Blog, Set isnull to rss attribue.
func (m *Blog) LogicalDelete() error {
	params := orm.Params{"Rss": nil, "User": nil}
	if _, err := Blogs().Filter("id", m).Update(params); err != nil {
		return err
	}
	return nil
}

func Blogs() orm.QuerySeter {
	return orm.NewOrm().QueryTable("blog").OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(Blog))
}
