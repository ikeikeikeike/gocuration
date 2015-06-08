package ormapper

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type RankingBase struct {
	Model

	BeginName string // dayly, weekly, monthly, yearly
	BeginTime time.Time

	Rank     int64 // rank order number
	PageView int64 // 1,2,3,4
}

type EntryRanking struct {
	RankingBase

	Entry   *Entry
	EntryId sql.NullInt64
}

type VideoRanking struct {
	RankingBase

	Video   *Video
	VideoId sql.NullInt64
}

type PictureRanking struct {
	RankingBase

	Picture   *Picture
	PictureId sql.NullInt64
}

func EntryRankings() *gorm.DB {
	return DB.Table("entry_ranking").
		Preload("Entry").
		Select("entry_ranking.*").
		Joins(`
		INNER JOIN entry ON entry.id = entry_ranking.entry_id 
		INNER JOIN blog ON blog.id = entry.blog_id 
		`)
}
