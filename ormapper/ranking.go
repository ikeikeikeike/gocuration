package ormapper

import (
	"database/sql"
	"time"
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
