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

// multiple fields unique key
func (u *EntryRanking) TableUnique() [][]string {
	return [][]string{
		[]string{"BeginName", "BeginTime", "Entry"},
	}
}

func (m *EntryRanking) LoadRelated() *EntryRanking {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Entry", 2, DefaultPerEntities)
	return m
}

func (m *EntryRanking) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *EntryRanking) UpdateRank(rank int64) error {
	m.Rank = rank
	if _, err := orm.NewOrm().Update(m, "Rank"); err != nil {
		return err
	}
	return nil
}

func (m *EntryRanking) PreviousRanking() (*EntryRanking, error) {
	var duration = m.BeginTime.UTC()

	switch m.BeginName {
	case "dayly":
		duration = duration.Add(-time.Hour * 24)
	case "weekly":
		duration = duration.AddDate(0, 0, -7)
	case "monthly":
		duration = duration.AddDate(0, -1, 0)
	case "yearly":
		duration = duration.AddDate(-1, 0, 0)
	}

	qs := orm.NewOrm().QueryTable("entry_ranking").
		Filter("entry", m.Entry.Id).
		Filter("begin_time", duration).
		Filter("begin_name", m.BeginName)

	var prev EntryRanking
	err := qs.One(&prev)
	if err != nil {
		return nil, err
	}

	return &prev, nil
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

// multiple fields unique key
func (u *VideoRanking) TableUnique() [][]string {
	return [][]string{
		[]string{"BeginName", "BeginTime", "Video"},
	}
}

func (m *VideoRanking) LoadRelated() *VideoRanking {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Video", 2, DefaultPerEntities)
	return m
}

func (m *VideoRanking) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *VideoRanking) UpdateRank(rank int64) error {
	m.Rank = rank
	if _, err := orm.NewOrm().Update(m, "Rank"); err != nil {
		return err
	}
	return nil
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

func (u *PictureRanking) TableUnique() [][]string {
	return [][]string{
		[]string{"BeginName", "BeginTime", "Picture"},
	}
}

func (m *PictureRanking) LoadRelated() *PictureRanking {
	o := orm.NewOrm()
	_, _ = o.LoadRelated(m, "Picture", 2, DefaultPerEntities)
	return m
}

func (m *PictureRanking) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *PictureRanking) UpdateRank(rank int64) error {
	m.Rank = rank
	if _, err := orm.NewOrm().Update(m, "Rank"); err != nil {
		return err
	}
	return nil
}

func EntryRankings() orm.QuerySeter {
	return orm.NewOrm().QueryTable("entry_ranking").OrderBy("Rank")
}

func VideoRankings() orm.QuerySeter {
	return orm.NewOrm().QueryTable("video_ranking").OrderBy("Rank")
}

func PictureRankings() orm.QuerySeter {
	return orm.NewOrm().QueryTable("picture_ranking").OrderBy("Rank")
}

func init() {
	orm.RegisterModelWithPrefix(beego.AppConfig.String("dbprefix"), new(EntryRanking))
	orm.RegisterModelWithPrefix(beego.AppConfig.String("dbprefix"), new(VideoRanking))
	orm.RegisterModelWithPrefix(beego.AppConfig.String("dbprefix"), new(PictureRanking))
}
