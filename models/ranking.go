package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type EntryRanking struct {
	Id int64 `orm:"auto"`

	BeginName string    `orm:"size(255);index" valid:"Required;Match(/^(dayly|weekly|monthly|yearly)$/)"` // dayly, weekly, monthly, yearly
	BeginTime time.Time `orm:"type(datetime);index"`

	Rank     int64 `orm:"default(0);index"` // rank order number
	PageView int64 `orm:"default(0);index"` // 1,2,3,4

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Entry *Entry `orm:"rel(fk);index"`
}

func (m *EntryRanking) LoadRelated() *EntryRanking {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Entry", 2, DefaultPerEntities)
	return m
}

func (m *EntryRanking) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

type VideoRanking struct {
	Id int64 `orm:"auto"`

	BeginName string    `orm:"size(255);index" valid:"Required;Match(/^(dayly|weekly|monthly|yearly)$/)"` // dayly, weekly, monthly, yearly
	BeginTime time.Time `orm:"type(datetime);index"`

	Rank     int64 `orm:"default(0);index"` // rank order number
	PageView int64 `orm:"default(0);index"` // 1,2,3,4

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Video *Video `orm:"rel(fk);index"`
}

func (m *VideoRanking) LoadRelated() *VideoRanking {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Video", 2, DefaultPerEntities)
	return m
}
func (m *VideoRanking) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

type PictureRanking struct {
	Id int64 `orm:"auto"`

	BeginName string    `orm:"size(255);index" valid:"Required;Match(/^(dayly|weekly|monthly|yearly)$/)"` // dayly, weekly, monthly, yearly
	BeginTime time.Time `orm:"type(datetime);index"`

	Rank     int64 `orm:"default(0);index"` // rank order number
	PageView int64 `orm:"default(0);index"` // 1,2,3,4

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Picture *Picture `orm:"rel(fk);index"`
}

func (m *PictureRanking) LoadRelated() *PictureRanking {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Picture", 2, DefaultPerEntities)
	return m
}
func (m *PictureRanking) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func EntryRankings() orm.QuerySeter {
	return orm.NewOrm().QueryTable("entry_ranking").OrderBy("-Id")
}

func VideoRankings() orm.QuerySeter {
	return orm.NewOrm().QueryTable("video_ranking").OrderBy("-Id")
}

func PictureRankings() orm.QuerySeter {
	return orm.NewOrm().QueryTable("picture_ranking").OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(beego.AppConfig.String("dbprefix"), new(EntryRanking))
	orm.RegisterModelWithPrefix(beego.AppConfig.String("dbprefix"), new(VideoRanking))
	orm.RegisterModelWithPrefix(beego.AppConfig.String("dbprefix"), new(PictureRanking))
}
