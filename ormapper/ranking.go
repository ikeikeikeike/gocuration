package ormapper

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type RankingBase struct {
	Id int64

	BeginName string // dayly, weekly, monthly, yearly
	BeginTime time.Time

	Rank     int64 // rank order number
	PageView int64 // 1,2,3,4

	Created time.Time
	Updated time.Time
}

type EntryRanking struct {
	RankingBase

	Entry   *Entry
	EntryId sql.NullInt64
}

func (m *EntryRanking) PreviousRanking() *EntryRanking {
	var prev EntryRanking
	db := EntryRankings().
		Where("entry_ranking.id != ?", m.Id).
		Where("entry_ranking.begin_name = ?", m.BeginName).
		Where("entry_ranking.entry_id = ?", m.Entry.Id).
		Order("entry_ranking.begin_time DESC").
		Limit(1).
		First(&prev)

	if db.Error != nil {
		return nil
	}

	return &prev
}

func (m *EntryRanking) RankingsLoader() {
	if m.Entry != nil {
		DB.Model(&m).
			Preload("Picture").Preload("Video").Preload("Blog").
			Related(&m.Entry)
		m.Entry.NewsLoader()
	}
}

type VideoRanking struct {
	RankingBase

	Video   *Video
	VideoId sql.NullInt64
}

func (m *VideoRanking) PreviousRanking() *VideoRanking {
	var prev VideoRanking
	db := VideoRankings().
		Where("video_ranking.id != ?", m.Id).
		Where("video_ranking.begin_name = ?", m.BeginName).
		Where("video_ranking.video_id = ?", m.VideoId.Int64).
		Order("video_ranking.begin_time DESC").
		Limit(1).
		First(&prev)

	if db.Error != nil {
		return nil
	}

	return &prev
}

func (m *VideoRanking) RankingsLoader() {
	if m.Video != nil {
		DB.Model(&m).
			Preload("Entry").Preload("Site").
			Related(&m.Video)
		m.Video.RankingsLoader()
	}
}

type PictureRanking struct {
	RankingBase

	Picture   *Picture
	PictureId sql.NullInt64
}

func (m *PictureRanking) PreviousRanking() *PictureRanking {
	var prev PictureRanking
	db := PictureRankings().
		Where("picture_ranking.id != ?", m.Id).
		Where("picture_ranking.begin_name = ?", m.BeginName).
		Where("picture_ranking.picture_id = ?", m.PictureId.Int64).
		Order("picture_ranking.begin_time DESC").
		Limit(1).
		First(&prev)

	if db.Error != nil {
		return nil
	}

	return &prev
}

func (m *PictureRanking) RankingsLoader() {
	if m.Picture != nil {
		DB.Model(&m).
			Preload("Entry").Preload("Anime").Preload("Images"). // character
			Related(&m.Picture)
		m.Picture.RankingsLoader()
	}
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

func PictureRankings() *gorm.DB {
	return DB.Table("picture_ranking").
		Preload("Picture").
		Select("picture_ranking.*").
		Joins(`
		INNER JOIN picture ON picture.id = picture_ranking.picture_id 
		INNER JOIN entry ON entry.id = picture.entry_id
		INNER JOIN blog ON blog.id = entry.blog_id 
		LEFT OUTER JOIN anime ON anime.id = picture.anime_id
		`)
}

func VideoRankings() *gorm.DB {
	return DB.Table("video_ranking").
		Preload("Video").
		Select("video_ranking.*").
		Joins(`
		INNER JOIN video ON video.id = video_ranking.video_id 
		INNER JOIN entry ON entry.id = video.entry_id
		INNER JOIN blog ON blog.id = entry.blog_id 

		LEFT OUTER JOIN entry_tag et ON et.entry_id = entry.id
		LEFT OUTER JOIN tag ON tag.id = et.tag_id
		`)
}
